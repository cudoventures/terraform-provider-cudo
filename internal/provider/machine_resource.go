package provider

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute"
	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/baremetal"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/grpc/codes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MachineResource{}
var _ resource.ResourceWithConfigure = &MachineResource{}
var _ resource.ResourceWithImportState = &MachineResource{}
var _ resource.ResourceWithModifyPlan = &MachineResource{}

func NewMachineResource() resource.Resource {
	return &MachineResource{}
}

func (*MachineResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_machine"
}

func (*MachineResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Machine resource",
		Attributes: map[string]schema.Attribute{
			"commitment_months": schema.Int32Attribute{
				MarkdownDescription: "The minimum length of time to commit to the machine. It cannot be deleted before the commitment end date.",
				Optional:            true,
				Validators:          []validator.Int32{int32validator.OneOf(1, 3, 6, 12, 24, 36)},
			},
			"data_center_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The id of the data center where the machine is located.",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid RFC1034 resource id")},
			},
			"project_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The project the machine is in.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid RFC1034 resource id")},
			},
			"id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "ID for machine within project.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid RFC1034 resource id")},
			},
			"machine_type": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Machine type of machine. See console for valid options",
				Required:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid RFC1034 resource id")},
			},
			"os": schema.StringAttribute{
				MarkdownDescription: "Operating system deployed to machine. See console for valid options",
				Optional:            true,
			},
			"user_data": schema.StringAttribute{
				MarkdownDescription: "cloud-init cloud-config applied on OS deployment.",
				Optional:            true,
			},
			"external_ip_addresses": schema.ListAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "The external IP addresses allocated to the machine.",
				Computed:            true,
			},
			"state": schema.StringAttribute{
				MarkdownDescription: "State of the machine.",
				Computed:            true,
			},
			"power_state": schema.StringAttribute{
				MarkdownDescription: "The power state of the machine.",
				Computed:            true,
			},
		},
	}
}

// MachineResource defines the resource implementation.
type MachineResource struct {
	client *CudoClientData
}

// MachineResourceModel describes the resource data model.
type MachineResourceModel struct {
	CommitmentMonths types.Int32  `tfsdk:"commitment_months"`
	DataCenterID     types.String `tfsdk:"data_center_id"`
	ID               types.String `tfsdk:"id"`
	MachineType      types.String `tfsdk:"machine_type"`
	ProjectID        types.String `tfsdk:"project_id"`
	OS               types.String `tfsdk:"os"`
	UserData         types.String `tfsdk:"user_data"`

	ExternalIPAddresses types.List   `tfsdk:"external_ip_addresses"`
	State               types.String `tfsdk:"state"`
	PowerState          types.String `tfsdk:"power_state"`
}

func (r *MachineResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MachineResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var projectID types.String
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("project_id"), &projectID)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if projectID.IsUnknown() && r.client.DefaultProjectID != "" {
		projectID = types.StringValue(r.client.DefaultProjectID)
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("project_id"), projectID)...)
	}
}

func (r *MachineResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MachineResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	commitmentTerm := compute.CommitmentTerm_COMMITMENT_TERM_NONE
	if !plan.CommitmentMonths.IsNull() {
		switch plan.CommitmentMonths.ValueInt32() {
		case 1:
			commitmentTerm = compute.CommitmentTerm_COMMITMENT_TERM_1_MONTH
		case 3:
			commitmentTerm = compute.CommitmentTerm_COMMITMENT_TERM_3_MONTHS
		case 6:
			commitmentTerm = compute.CommitmentTerm_COMMITMENT_TERM_6_MONTHS
		case 12:
			commitmentTerm = compute.CommitmentTerm_COMMITMENT_TERM_12_MONTHS
		case 24:
			commitmentTerm = compute.CommitmentTerm_COMMITMENT_TERM_24_MONTHS
		case 36:
			commitmentTerm = compute.CommitmentTerm_COMMITMENT_TERM_36_MONTHS
		}
	}

	params := &baremetal.CreateMachineRequest{
		Machine: &baremetal.Machine{
			CommitmentTerm: commitmentTerm,
			DataCenterId:   plan.DataCenterID.ValueString(),
			Id:             plan.ID.ValueString(),
			MachineTypeId:  plan.MachineType.ValueString(),
			ProjectId:      plan.ProjectID.ValueString(),
			Os:             plan.OS.ValueString(),
			UserData:       plan.UserData.ValueString(),
		},
	}

	_, err := r.client.BareMetalClient.CreateMachine(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create machine",
			err.Error(),
		)
		return
	}

	machine, err := waitForMachineDeploying(ctx, r.client.BareMetalClient, plan.ProjectID.ValueString(), plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for machine to deploy",
			err.Error(),
		)
		return
	}

	// if the machine is created and returned update the state.
	if machine != nil {
		resp.Diagnostics.Append(appendMachineState(machine, &plan)...)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func waitForMachineDeploying(ctx context.Context, c baremetal.BareMetalServiceClient, projectID string, ID string) (*baremetal.Machine, error) {
	refreshFunc := func() (interface{}, string, error) {
		params := &baremetal.GetMachineRequest{
			Id:        ID,
			ProjectId: projectID,
		}
		machine, err := c.GetMachine(ctx, params)
		if err != nil {
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				return machine, "NOT_FOUND", nil
			}
			return nil, baremetal.Machine_STATE_UNSPECIFIED.String(), err
		}

		return machine, machine.State.String(), nil
	}

	stateConf := &helper.StateChangeConf{
		Pending: []string{
			baremetal.Machine_ACTIVE.String(),
			baremetal.Machine_CREATING.String(),
		},
		Target: []string{
			baremetal.Machine_STATE_UNSPECIFIED.String(),
			baremetal.Machine_DELETING.String(),
			baremetal.Machine_DEPLOYING.String(),
			baremetal.Machine_FAILED.String(),
			baremetal.Machine_UPDATING.String(),
			"NOT_FOUND",
		},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for machine %s in project %s: %w", ID, projectID, err)
	} else if machine, ok := res.(*baremetal.Machine); ok {
		return machine, nil
	} else {
		return nil, fmt.Errorf("error waiting for machine: %v", res)
	}
}

func (r *MachineResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MachineResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := &baremetal.GetMachineRequest{
		ProjectId: state.ProjectID.ValueString(),
		Id:        state.ID.ValueString(),
	}

	machine, err := r.client.BareMetalClient.GetMachine(ctx, params)
	if err != nil {
		if ok := helper.IsErrCode(err, codes.NotFound); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to read machine",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(appendMachineState(machine, &state)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *MachineResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Unable to update machine",
		"Updating a machine is not supported",
	)
}

func (r *MachineResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MachineResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	projectID := state.ProjectID.ValueString()
	if state.ProjectID.IsNull() {
		projectID = r.client.DefaultProjectID
	}

	ID := state.ID.ValueString()

	_, err := r.client.BareMetalClient.DeleteMachine(ctx, &baremetal.DeleteMachineRequest{
		ProjectId: projectID,
		Id:        ID,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete machine",
			err.Error(),
		)
		return
	}

	err = waitForMachineDeleting(ctx, r.client.BareMetalClient, projectID, ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for machine to be deleted",
			err.Error(),
		)
		return
	}
}

func waitForMachineDeleting(ctx context.Context, c baremetal.BareMetalServiceClient, projectID string, ID string) error {
	refreshFunc := func() (interface{}, string, error) {
		params := &baremetal.GetMachineRequest{
			Id:        ID,
			ProjectId: projectID,
		}
		machine, err := c.GetMachine(ctx, params)
		if err != nil {
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				return machine, "NOT_FOUND", nil
			}
			return nil, baremetal.Machine_STATE_UNSPECIFIED.String(), err
		}

		return machine, machine.State.String(), nil
	}

	stateConf := &helper.StateChangeConf{
		Pending: []string{
			baremetal.Machine_DELETING.String(),
		},
		Target: []string{
			baremetal.Machine_CREATING.String(),
			baremetal.Machine_DEPLOYING.String(),
			baremetal.Machine_STATE_UNSPECIFIED.String(),
			baremetal.Machine_ACTIVE.String(),
			baremetal.Machine_FAILED.String(),
			"NOT_FOUND",
		},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return fmt.Errorf("error waiting for machine %s in project %s: %w", ID, projectID, err)
	} else if _, ok := res.(*baremetal.Machine); ok {
		return nil
	} else {
		return fmt.Errorf("error waiting for machine: %v", res)
	}
}

func appendMachineState(machine *baremetal.Machine, state *MachineResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	switch machine.CommitmentTerm {
	case compute.CommitmentTerm_COMMITMENT_TERM_NONE:
		state.CommitmentMonths = types.Int32Null()
	case compute.CommitmentTerm_COMMITMENT_TERM_1_MONTH:
		state.CommitmentMonths = types.Int32Value(1)
	case compute.CommitmentTerm_COMMITMENT_TERM_3_MONTHS:
		state.CommitmentMonths = types.Int32Value(3)
	case compute.CommitmentTerm_COMMITMENT_TERM_6_MONTHS:
		state.CommitmentMonths = types.Int32Value(6)
	case compute.CommitmentTerm_COMMITMENT_TERM_12_MONTHS:
		state.CommitmentMonths = types.Int32Value(12)
	case compute.CommitmentTerm_COMMITMENT_TERM_24_MONTHS:
		state.CommitmentMonths = types.Int32Value(24)
	case compute.CommitmentTerm_COMMITMENT_TERM_36_MONTHS:
		state.CommitmentMonths = types.Int32Value(36)
	}

	state.DataCenterID = types.StringValue(machine.DataCenterId)
	state.MachineType = types.StringValue(machine.MachineTypeId)
	state.ProjectID = types.StringValue(machine.ProjectId)
	state.ID = types.StringValue(machine.Id)
	ipAddresses := make([]attr.Value, len(machine.ExternalIpAddresses))
	for i, ipAddress := range machine.ExternalIpAddresses {
		ipAddresses[i] = types.StringValue(ipAddress)
	}
	var d diag.Diagnostics
	state.ExternalIPAddresses, d = types.ListValue(types.StringType, ipAddresses)
	if d != nil {
		diags.Append(d...)
	}
	state.OS = types.StringValue(machine.Os)
	state.PowerState = types.StringValue(machine.PowerState.String())
	state.State = types.StringValue(machine.State.String())

	return diags
}

var machineImportIDRegExp = regexp.MustCompile("projects/(.+)/machines/(.+)")

func (r *MachineResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var projectID, ID string
	if parts := machineImportIDRegExp.FindStringSubmatch(req.ID); parts != nil {
		projectID = parts[1]
		ID = parts[2]
	}

	if projectID == "" || ID == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: \"projects/<project_id>/machines/<id>\". Got: %q", req.ID),
		)
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("project_id"), projectID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), ID)...)
}
