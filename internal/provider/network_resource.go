package provider

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client/networks"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"
	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

func waitForNetworkAvailable(ctx context.Context, projectID string, networkID string, c networks.ClientService) (*networks.GetNetworkOK, error) {
	refreshFunc := func() (interface{}, string, error) {
		params := networks.NewGetNetworkParamsWithContext(ctx)
		params.ID = networkID
		params.ProjectID = projectID
		res, err := c.GetNetwork(params)
		if err != nil {
			if apiErr, ok := err.(*networks.GetNetworkDefault); ok && apiErr.IsCode(404) {
				tflog.Debug(ctx, fmt.Sprintf("Network %s in project %s not found: ", networkID, projectID))
				return res, "done", nil
			}
			return nil, "", err
		}

		tflog.Trace(ctx, fmt.Sprintf("pending network %s in project %s state: %s", networkID, projectID, *res.Payload.Network.ShortState))
		return res, *res.Payload.Network.ShortState, nil
	}

	tflog.Debug(ctx, fmt.Sprintf("waiting for network %s in project %s ", networkID, projectID))

	stateConf := &helper.StateChangeConf{
		Pending:    []string{"clea", "clon", "dsrz", "epil", "hold", "hotp", "init", "migr", "pend", "prol", "save", "shut", "snap", "unkn"},
		Target:     []string{"boot", "done", "fail", "poff", "runn", "stop", "susp", "unde"},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for network %s in project %s to become available: %w", networkID, projectID, err)
	} else if vm, ok := res.(*networks.GetNetworkOK); ok {
		tflog.Trace(ctx, fmt.Sprintf("completed waiting for network %s in project %s (%s)", networkID, projectID, *vm.Payload.Network.ShortState))
		return vm, nil
	}

	return nil, nil
}

func waitForNetworkStop(ctx context.Context, projectID string, networkID string, c networks.ClientService) (*networks.GetNetworkOK, error) {
	refreshFunc := func() (interface{}, string, error) {
		params := networks.NewGetNetworkParamsWithContext(ctx)
		params.ID = networkID
		params.ProjectID = projectID
		res, err := c.GetNetwork(params)
		if err != nil {
			if apiErr, ok := err.(*networks.GetNetworkDefault); ok && apiErr.IsCode(404) {
				tflog.Debug(ctx, fmt.Sprintf("Network %s in project %s is done: ", networkID, projectID))
				return res, "done", nil
			}
			tflog.Error(ctx, fmt.Sprintf("error getting network %s in project %s: %v", networkID, projectID, err))
			return nil, "", err
		}
		if *res.Payload.Network.ShortState == "" {
			tflog.Debug(ctx, fmt.Sprintf("Network %s in project %s is stopped: ", networkID, projectID))
			return res, "done", nil
		}

		tflog.Trace(ctx, fmt.Sprintf("pending network %s in project %s state: %s", networkID, projectID, *res.Payload.Network.ShortState))
		return res, *res.Payload.Network.ShortState, nil
	}

	tflog.Debug(ctx, fmt.Sprintf("waiting for network %s in project %s ", networkID, projectID))

	stateConf := &helper.StateChangeConf{
		Pending:    []string{"fail", "poff", "runn", "stop", "susp", "unde", "boot", "clea", "clon", "dsrz", "epil", "hold", "hotp", "init", "migr", "pend", "prol", "save", "shut", "snap", "unkn"},
		Target:     []string{"done"},
		Refresh:    refreshFunc,
		Timeout:    20 * time.Minute,
		MinTimeout: 3 * time.Second,
	}

	if _, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for network %s in project %s to be stopped: %w", networkID, projectID, err)
	}

	return nil, nil
}

func waitForNetworkDelete(ctx context.Context, projectID string, networkID string, c networks.ClientService) (*networks.GetNetworkOK, error) {
	refreshFunc := func() (interface{}, string, error) {
		params := networks.NewGetNetworkParamsWithContext(ctx)
		params.ID = networkID
		params.ProjectID = projectID
		res, err := c.GetNetwork(params)
		if err != nil {
			if apiErr, ok := err.(*networks.GetNetworkDefault); ok && apiErr.IsCode(404) {
				tflog.Debug(ctx, fmt.Sprintf("Network %s in project %s is done: ", networkID, projectID))
				return res, "done", nil
			}
			tflog.Error(ctx, fmt.Sprintf("error getting network %s in project %s: %v", networkID, projectID, err))
			return nil, "", err
		}
		if *res.Payload.Network.ShortState == "" {
			tflog.Debug(ctx, fmt.Sprintf("Network %s in project %s is stopped: ", networkID, projectID))
			return res, "stop", nil
		}

		tflog.Trace(ctx, fmt.Sprintf("pending network %s in project %s state: %s", networkID, projectID, *res.Payload.Network.ShortState))
		return res, *res.Payload.Network.ShortState, nil
	}

	tflog.Debug(ctx, fmt.Sprintf("waiting for network %s in project %s ", networkID, projectID))

	stateConf := &helper.StateChangeConf{
		Pending:    []string{"fail", "poff", "runn", "stop", "susp", "unde", "boot", "clea", "clon", "dsrz", "epil", "hold", "hotp", "init", "migr", "pend", "prol", "save", "shut", "snap", "unkn"},
		Target:     []string{"done"},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		MinTimeout: 3 * time.Second,
	}

	if _, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for network %s in project %s to be deleted: %w", networkID, projectID, err)
	}

	return nil, nil
}

func (r *NetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state *NetworkResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := networks.NewCreateNetworkParamsWithContext(ctx)
	params.ProjectID = r.client.DefaultProjectID
	params.Body = &models.CreateNetworkBody{
		CidrPrefix:   state.IPRange.ValueStringPointer(),
		DataCenterID: state.DataCenterId.ValueStringPointer(),
		ID:           state.ID.ValueStringPointer(),
		VrouterSize:  models.V1VRouterSizeVROUTERINSTANCESMALL.Pointer(),
	}
	_, err := r.client.Client.Networks.CreateNetwork(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create network resource",
			err.Error(),
		)
		return
	}

	res, err := waitForNetworkAvailable(ctx, params.ProjectID, state.ID.ValueString(), r.client.Client.Networks)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create network resource",
			err.Error(),
		)
		return
	}

	state.Gateway = types.StringValue(*res.Payload.Network.Gateway)
	state.ExternalIPAddress = types.StringValue(*res.Payload.Network.ExternalIPAddress)
	state.InternalIPAddress = types.StringValue(*res.Payload.Network.InternalIPAddress)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *NetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *NetworkResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	getParams := networks.NewGetNetworkParamsWithContext(ctx)
	getParams.ID = state.ID.ValueString()
	getParams.ProjectID = r.client.DefaultProjectID

	resget, err := r.client.Client.Networks.GetNetwork(getParams)
	if err != nil {
		if apiErr, ok := err.(*networks.GetNetworkDefault); ok && apiErr.IsCode(404) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to read network resource",
			err.Error(),
		)
		return
	}

	state.ID = types.StringValue(*resget.Payload.Network.ID)
	state.DataCenterId = types.StringValue(*resget.Payload.Network.DataCenterID)
	state.ExternalIPAddress = types.StringValue(*resget.Payload.Network.ExternalIPAddress)
	state.InternalIPAddress = types.StringValue(*resget.Payload.Network.InternalIPAddress)
	state.IPRange = types.StringValue(*resget.Payload.Network.IPRange)
	state.Gateway = types.StringValue(*resget.Payload.Network.Gateway)

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

	stopParams := networks.NewStopNetworkParamsWithContext(ctx)
	stopParams.ProjectID = r.client.DefaultProjectID
	stopParams.ID = state.ID.ValueString()

	_, err := r.client.Client.Networks.StopNetwork(stopParams)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to stop network resource",
			err.Error(),
		)
		return
	}

	_, err = waitForNetworkStop(ctx, stopParams.ProjectID, stopParams.ID, r.client.Client.Networks)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for network resource to be stopped",
			err.Error(),
		)
		return
	}

	deleteParams := networks.NewDeleteNetworkParamsWithContext(ctx)
	deleteParams.ProjectID = r.client.DefaultProjectID
	deleteParams.ID = state.ID.ValueString()

	_, err = r.client.Client.Networks.DeleteNetwork(deleteParams)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete network resource",
			err.Error(),
		)
		return
	}

	_, err = waitForNetworkDelete(ctx, deleteParams.ProjectID, deleteParams.ID, r.client.Client.Networks)
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
