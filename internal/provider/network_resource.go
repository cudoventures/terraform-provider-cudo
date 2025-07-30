package provider

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/network"
	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/vm"
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
var _ resource.ResourceWithImportState = &NetworkResource{}

func NewNetworkResource() resource.Resource {
	return &NetworkResource{}
}

// NetworkResource defines the resource implementation.
type NetworkResource struct {
	client *CudoClientData
}

// NetworkResourceModel describes the resource data model.
type NetworkResourceModel struct {
	ID                types.String `tfsdk:"id"`
	DataCenterId      types.String `tfsdk:"data_center_id"`
	IPRange           types.String `tfsdk:"ip_range"`
	Gateway           types.String `tfsdk:"gateway"`
	ExternalIPAddress types.String `tfsdk:"external_ip_address"`
	InternalIPAddress types.String `tfsdk:"internal_ip_address"`
}

func (r *NetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "cudo_network"
}

func (r *NetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Network resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Network ID",
				Required:            true,
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"data_center_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The unique identifier of the datacenter where the network is located.",
				Required:            true,
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"ip_range": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "IP range of network in CIDR format e.g 192.168.0.0/24",
				Required:            true,
			},
			"gateway": schema.StringAttribute{
				MarkdownDescription: "Internal IP of the network gateway",
				Computed:            true,
			},
			"external_ip_address": schema.StringAttribute{
				MarkdownDescription: "External IP of the network router",
				Computed:            true,
			},
			"internal_ip_address": schema.StringAttribute{
				MarkdownDescription: "Internal IP of the network router",
				Computed:            true,
			},
		},
	}
}

func (r *NetworkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
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

func (r *NetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state NetworkResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.NetworkClient.CreateNetwork(ctx, &network.CreateNetworkRequest{
		ProjectId:    r.client.DefaultProjectID,
		CidrPrefix:   state.IPRange.ValueString(),
		DataCenterId: state.DataCenterId.ValueString(),
		Id:           state.ID.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create network resource",
			err.Error(),
		)
		return
	}

	network, err := waitForNetworkAvailable(ctx, r.client.DefaultProjectID, state.ID.ValueString(), r.client.NetworkClient)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create network resource",
			err.Error(),
		)
		return
	}

	if network != nil {
		state.Gateway = types.StringValue(network.Gateway)
		state.ExternalIPAddress = types.StringValue(network.ExternalIpAddress)
		state.InternalIPAddress = types.StringValue(network.InternalIpAddress)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *NetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.NetworkClient.GetNetwork(ctx, &network.GetNetworkRequest{
		Id:        state.ID.ValueString(),
		ProjectId: r.client.DefaultProjectID,
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

	state.ID = types.StringValue(res.Id)
	state.DataCenterId = types.StringValue(res.DataCenterId)
	state.ExternalIPAddress = types.StringValue(res.ExternalIpAddress)
	state.InternalIPAddress = types.StringValue(res.InternalIpAddress)
	state.IPRange = types.StringValue(res.IpRange)
	state.Gateway = types.StringValue(res.Gateway)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *NetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworkResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Error getting network plan",
			"Error getting network plan",
		)
		return
	}

	// Read Terraform state data into the model
	var state NetworkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *NetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *NetworkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	projectId := r.client.DefaultProjectID
	networkId := state.ID.ValueString()
	_, err := r.client.NetworkClient.StopNetwork(ctx, &network.StopNetworkRequest{
		ProjectId: projectId,
		Id:        networkId,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to stop network resource",
			err.Error(),
		)
		return
	}

	_, err = waitForNetworkStop(ctx, projectId, networkId, r.client.NetworkClient)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for network resource to be stopped",
			err.Error(),
		)
		return
	}

	_, err = r.client.NetworkClient.DeleteNetwork(ctx, &network.DeleteNetworkRequest{
		ProjectId: projectId,
		Id:        networkId,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete network resource",
			err.Error(),
		)
		return
	}

	_, err = waitForNetworkDelete(ctx, projectId, networkId, r.client.NetworkClient)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for network resource to be deleted",
			err.Error(),
		)
		return
	}
}

func (r *NetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func waitForNetworkAvailable(ctx context.Context, projectID string, networkID string, c network.NetworkServiceClient) (*network.Network, error) {
	refreshFunc := func() (interface{}, string, error) {
		res, err := c.GetNetwork(ctx, &network.GetNetworkRequest{
			Id:        networkID,
			ProjectId: projectID,
		})
		if err != nil {
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				tflog.Debug(ctx, fmt.Sprintf("Network %s in project %s not found: ", networkID, projectID))
				return res, vm.VM_UNKNOWN.String(), nil
			}
			return nil, "", err
		}

		tflog.Trace(ctx, fmt.Sprintf("pending network %s in project %s state: %s", networkID, projectID, res.State.String()))
		return res, res.State.String(), nil
	}

	tflog.Debug(ctx, fmt.Sprintf("waiting for network %s in project %s", networkID, projectID))

	stateConf := &helper.StateChangeConf{
		Pending: []string{
			vm.VM_CLONING.String(),
			vm.VM_CREATING_SNAPSHOT.String(),
			vm.VM_DELETING_SNAPSHOT.String(),
			vm.VM_DELETING.String(),
			vm.VM_HOTPLUGGING.String(),
			vm.VM_MIGRATING.String(),
			vm.VM_RECREATING.String(),
			vm.VM_RESIZING_DISK.String(),
			vm.VM_RESIZING.String(),
			vm.VM_REVERTING_SNAPSHOT.String(),
			vm.VM_STARTING.String(),
			vm.VM_STOPPING.String(),
			vm.VM_SUSPENDING.String(),
			vm.VM_UNKNOWN.String(),
			vm.VM_PENDING.String(),
		},
		Target: []string{
			vm.VM_ACTIVE.String(),
		},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for network %s in project %s to become available: %w", networkID, projectID, err)
	} else if res, ok := res.(*network.Network); ok {
		tflog.Trace(ctx, fmt.Sprintf("completed waiting for network %s in project %s (%s)", networkID, projectID, res.State.String()))
		return res, nil
	}

	return nil, nil
}

func waitForNetworkStop(ctx context.Context, projectID string, networkID string, c network.NetworkServiceClient) (*network.Network, error) {
	refreshFunc := func() (interface{}, string, error) {
		res, err := c.GetNetwork(ctx, &network.GetNetworkRequest{
			Id:        networkID,
			ProjectId: projectID,
		})
		if err != nil {
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				tflog.Debug(ctx, fmt.Sprintf("Network %s in project %s is done: ", networkID, projectID))
				return res, vm.VM_DELETED.String(), nil
			}
			tflog.Error(ctx, fmt.Sprintf("error getting network %s in project %s: %v", networkID, projectID, err))
			return nil, "", err
		}

		tflog.Trace(ctx, fmt.Sprintf("pending network %s in project %s state: %s", networkID, projectID, res.State.String()))
		return res, res.State.String(), nil
	}

	tflog.Debug(ctx, fmt.Sprintf("waiting for network %s in project %s ", networkID, projectID))

	stateConf := &helper.StateChangeConf{
		Pending: []string{
			vm.VM_CLONING.String(),
			vm.VM_CREATING_SNAPSHOT.String(),
			vm.VM_DELETING_SNAPSHOT.String(),
			vm.VM_HOTPLUGGING.String(),
			vm.VM_MIGRATING.String(),
			vm.VM_RECREATING.String(),
			vm.VM_RESIZING_DISK.String(),
			vm.VM_RESIZING.String(),
			vm.VM_REVERTING_SNAPSHOT.String(),
			vm.VM_STARTING.String(),
			vm.VM_STOPPING.String(),
			vm.VM_SUSPENDING.String(),
			vm.VM_PENDING.String(),
			// network could be stopped whilst in these states, so network could be in the state when the delete is requested
			vm.VM_UNKNOWN.String(),
			vm.VM_FAILED.String(),
			vm.VM_ACTIVE.String(),
			vm.VM_SUSPENDED.String(),
		},
		Target: []string{
			vm.VM_STOPPED.String(),
			vm.VM_DELETING.String(),
			vm.VM_DELETED.String(),
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

func waitForNetworkDelete(ctx context.Context, projectID string, networkID string, c network.NetworkServiceClient) (*network.Network, error) {
	refreshFunc := func() (interface{}, string, error) {
		res, err := c.GetNetwork(ctx, &network.GetNetworkRequest{
			Id:        networkID,
			ProjectId: projectID,
		})
		if err != nil {
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				tflog.Debug(ctx, fmt.Sprintf("Network %s in project %s is done: ", networkID, projectID))
				return res, "deleted", nil
			}

			tflog.Error(ctx, fmt.Sprintf("error getting network %s in project %s: %v", networkID, projectID, err))
			return nil, "", err
		}

		tflog.Trace(ctx, fmt.Sprintf("pending network %s in project %s state: %s", networkID, projectID, res.State.String()))
		return res, "pending", nil
	}

	tflog.Debug(ctx, fmt.Sprintf("waiting for network %s in project %s ", networkID, projectID))

	stateConf := &helper.StateChangeConf{
		Pending:    []string{"pending"},
		Target:     []string{"deleted"},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		MinTimeout: 3 * time.Second,
	}

	if _, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for network %s in project %s to be deleted: %w", networkID, projectID, err)
	}

	return nil, nil
}
