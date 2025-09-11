package provider

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/network"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"google.golang.org/grpc/codes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NetworkResource{}
var _ resource.ResourceWithConfigure = &NetworkResource{}
var _ resource.ResourceWithImportState = &NetworkResource{}
var _ resource.ResourceWithModifyPlan = &NetworkResource{}

func NewNetworkResource() resource.Resource {
	return &NetworkResource{}
}

// NetworkResource defines the resource implementation.
type NetworkResource struct {
	client *CudoClientData
}

// NetworkResourceModel describes the resource data model.
type NetworkResourceModel struct {
	DataCenterID      types.String `tfsdk:"data_center_id"`
	ExternalIPAddress types.String `tfsdk:"external_ip_address"`
	Gateway           types.String `tfsdk:"gateway"`
	ID                types.String `tfsdk:"id"`
	InternalIPAddress types.String `tfsdk:"internal_ip_address"`
	IPRange           types.String `tfsdk:"ip_range"`
	ProjectID         types.String `tfsdk:"project_id"`
}

func (r *NetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network"
}

func (r *NetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Network resource",
		Attributes: map[string]schema.Attribute{
			"data_center_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The unique identifier of the datacenter where the network is located.",
				Required:            true,
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"external_ip_address": schema.StringAttribute{
				MarkdownDescription: "External IP of the network router",
				Computed:            true,
			},
			"gateway": schema.StringAttribute{
				MarkdownDescription: "Internal IP of the network gateway",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Network ID",
				Required:            true,
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"internal_ip_address": schema.StringAttribute{
				MarkdownDescription: "Internal IP of the network router",
				Computed:            true,
			},
			"ip_range": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "IP range of network in CIDR format e.g 192.168.0.0/24",
				Required:            true,
			},
			"project_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The project the network is in.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid RFC1034 resource id")},
			},
		},
	}
}

func (r *NetworkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*CudoClientData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *CudoClientData got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *NetworkResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var projectID types.String

	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("project_id"), &projectID)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if projectID.IsUnknown() && r.client.DefaultProjectID != "" {
		projectID = types.StringValue(r.client.DefaultProjectID)
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("project_id"), projectID)...)
	}
}

func (r *NetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.NetworkClient.CreateNetwork(ctx, &network.CreateNetworkRequest{
		Network: &network.Network{
			IpRange:      plan.IPRange.ValueString(),
			DataCenterId: plan.DataCenterID.ValueString(),
			Id:           plan.ID.ValueString(),
			ProjectId:    plan.ProjectID.ValueString(),
		},
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create network resource",
			err.Error(),
		)
		return
	}

	network, err := waitForNetworkAvailable(ctx, r.client.NetworkClient, r.client.DefaultProjectID, plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create network resource",
			err.Error(),
		)
		return
	}

	if network != nil {
		plan.Gateway = types.StringValue(network.Gateway)
		plan.ExternalIPAddress = types.StringValue(network.ExternalIpAddress)
		plan.InternalIPAddress = types.StringValue(network.InternalIpAddress)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.ProjectID.IsNull() {
		state.ProjectID = types.StringValue(r.client.DefaultProjectID)
	}

	res, err := r.client.NetworkClient.GetNetwork(ctx, &network.GetNetworkRequest{
		Id:        state.ID.ValueString(),
		ProjectId: state.ProjectID.ValueString(),
	})
	if err != nil {
		if ok := helper.IsErrCode(err, codes.NotFound); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to read network resource",
			err.Error(),
		)
		return
	}

	state.DataCenterID = types.StringValue(res.DataCenterId)
	state.ExternalIPAddress = types.StringValue(res.ExternalIpAddress)
	state.Gateway = types.StringValue(res.Gateway)
	state.ID = types.StringValue(res.Id)
	state.InternalIPAddress = types.StringValue(res.InternalIpAddress)
	state.IPRange = types.StringValue(res.IpRange)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *NetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Unable to update network",
		"Updating a network is not supported",
	)
}

func (r *NetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *NetworkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	projectID := state.ProjectID.ValueString()
	if state.ProjectID.IsNull() {
		projectID = r.client.DefaultProjectID
	}

	ID := state.ID.ValueString()

	_, err := r.client.NetworkClient.DeleteNetwork(ctx, &network.DeleteNetworkRequest{
		ProjectId: projectID,
		Id:        ID,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete network resource",
			err.Error(),
		)
		return
	}

	_, err = waitForNetworkDelete(ctx, r.client.NetworkClient, projectID, ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for network resource to be deleted",
			err.Error(),
		)
		return
	}
}

func waitForNetworkAvailable(ctx context.Context, c network.NetworkServiceClient, projectID string, networkID string) (*network.Network, error) {
	refreshFunc := func() (interface{}, string, error) {
		res, err := c.GetNetwork(ctx, &network.GetNetworkRequest{
			Id:        networkID,
			ProjectId: projectID,
		})
		if err != nil {
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				return res, network.Network_DELETED.String(), nil
			}
			return nil, network.Network_STATE_UNKNOWN.String(), err
		}

		return res, res.State.String(), nil
	}

	stateConf := &helper.StateChangeConf{
		Pending: []string{
			network.Network_CREATING.String(),
			network.Network_DELETING.String(),
			network.Network_STARTING.String(),
			network.Network_STOPPING.String(),
			network.Network_UPDATING.String(),
		},
		Target: []string{
			network.Network_STATE_UNKNOWN.String(),
			network.Network_ACTIVE.String(),
			network.Network_DELETED.String(),
			network.Network_FAILED.String(),
			network.Network_STOPPED.String(),
		},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for network %s in project %s to become available: %w", networkID, projectID, err)
	} else if res, ok := res.(*network.Network); ok {
		var state string
		if res != nil {
			state = res.State.String()
		}
		tflog.Trace(ctx, fmt.Sprintf("completed waiting for network %s in project %s (%s)", networkID, projectID, state))
		return res, nil
	}

	return nil, nil
}

func waitForNetworkDelete(ctx context.Context, c network.NetworkServiceClient, projectID string, networkID string) (*network.Network, error) {
	refreshFunc := func() (interface{}, string, error) {
		res, err := c.GetNetwork(ctx, &network.GetNetworkRequest{
			Id:        networkID,
			ProjectId: projectID,
		})
		if err != nil {
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				return res, network.Network_DELETED.String(), nil
			}
			return nil, network.Network_STATE_UNKNOWN.String(), err
		}

		return res, res.State.String(), nil
	}

	stateConf := &helper.StateChangeConf{
		Pending: []string{
			network.Network_ACTIVE.String(),
			network.Network_CREATING.String(),
			network.Network_DELETING.String(),
			network.Network_STARTING.String(),
			network.Network_STOPPED.String(),
			network.Network_STOPPING.String(),
			network.Network_UPDATING.String(),
		},
		Target: []string{
			network.Network_STATE_UNKNOWN.String(),
			network.Network_DELETED.String(),
			network.Network_FAILED.String(),
		},
		Refresh:    refreshFunc,
		Timeout:    20 * time.Minute,
		MinTimeout: 3 * time.Second,
	}

	if _, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for network %s in project %s to be stopped: %w", networkID, projectID, err)
	}

	return nil, nil
}

var networkImportIDRegExp = regexp.MustCompile("projects/(.+)/networks/(.+)")

func (r *NetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var projectID, ID string
	if parts := networkImportIDRegExp.FindStringSubmatch(req.ID); parts != nil {
		projectID = parts[1]
		ID = parts[2]
	}

	if projectID == "" || ID == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: \"projects/<project_id>/networks/<id>\". Got: %q", req.ID),
		)
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("project_id"), projectID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), ID)...)
}
