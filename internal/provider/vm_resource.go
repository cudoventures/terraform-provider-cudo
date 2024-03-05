package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client/compute/vm"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"google.golang.org/genproto/googleapis/type/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VMResource{}
var _ resource.ResourceWithImportState = &VMResource{}

func NewVMResource() resource.Resource {
	return &VMResource{}
}

// VMResource defines the resource implementation.
type VMResource struct {
	client *CudoClientData
}

// VMResourceModel describes the resource data model.
type VMResourceModel struct {
	BootDisk          *VMBootDiskResourceModel `tfsdk:"boot_disk"`
	DataCenterID      types.String             `tfsdk:"data_center_id"`
	CPUModel          types.String             `tfsdk:"cpu_model"`
	GPUs              types.Int64              `tfsdk:"gpus"`
	GPUModel          types.String             `tfsdk:"gpu_model"`
	ID                types.String             `tfsdk:"id"`
	MachineType       types.String             `tfsdk:"machine_type"`
	MaxPriceHr        types.String             `tfsdk:"max_price_hr"`
	MemoryGib         types.Int64              `tfsdk:"memory_gib"`
	Password          types.String             `tfsdk:"password"`
	PriceHr           types.String             `tfsdk:"price_hr"`
	ProjectID         types.String             `tfsdk:"project_id"`
	SSHKeys           []types.String           `tfsdk:"ssh_keys"`
	SSHKeySource      types.String             `tfsdk:"ssh_key_source"`
	StartScript       types.String             `tfsdk:"start_script"`
	VCPUs             types.Int64              `tfsdk:"vcpus"`
	Networks          []*VMNICResourceModel    `tfsdk:"networks"`
	InternalIPAddress types.String             `tfsdk:"internal_ip_address"`
	ExternalIPAddress types.String             `tfsdk:"external_ip_address"`
	RenewableEnergy   types.Bool               `tfsdk:"renewable_energy"`
	SecurityGroupIDs  types.Set                `tfsdk:"security_group_ids"`
}

type VMBootDiskResourceModel struct {
	ImageID types.String `tfsdk:"image_id"`
	SizeGib types.Int64  `tfsdk:"size_gib"`
}

type VMNICResourceModel struct {
	NetworkID         types.String `tfsdk:"network_id"`
	AssignPublicIP    types.Bool   `tfsdk:"assign_public_ip"`
	InternalIPAddress types.String `tfsdk:"internal_ip_address"`
	ExternalIPAddress types.String `tfsdk:"external_ip_address"`
	SecurityGroupIDs  types.Set    `tfsdk:"security_group_ids"`
}

func (r *VMResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *VMResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state *VMResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	sshKeySource := vm.CreateVMRequest_SSH_KEY_SOURCE_PROJECT
	switch state.SSHKeySource.ValueString() {
	case "user":
		sshKeySource = vm.CreateVMRequest_SSH_KEY_SOURCE_USER
	case "custom":
		sshKeySource = vm.CreateVMRequest_SSH_KEY_SOURCE_NONE
	}

	var customKeys []string
	if sshKeySource == vm.CreateVMRequest_SSH_KEY_SOURCE_NONE {
		for _, key := range state.SSHKeys {
			customKeys = append(customKeys, key.ValueString())
		}
	}

	// retryClient := retryablehttp.NewClient()
	// retryClient.RetryMax = 10
	projectId := r.client.DefaultProjectID
	if !state.ProjectID.IsNull() {
		projectId = state.ProjectID.ValueString()
	}

	var bootDisk vm.Disk
	if !state.BootDisk.SizeGib.IsNull() {
		sizeGib := int32(state.BootDisk.SizeGib.ValueInt64())
		bootDisk.SizeGib = sizeGib
	}
	nics := make([]*vm.CreateVMRequest_NIC, len(state.Networks))

	for i, nic := range state.Networks {
		var securityGroupIDS []string
		if !nic.SecurityGroupIDs.IsNull() {
			resp.Diagnostics.Append(nic.SecurityGroupIDs.ElementsAs(ctx, &securityGroupIDS, false)...)
			if resp.Diagnostics.HasError() {
				return
			}
		}
		nics[i] = &vm.CreateVMRequest_NIC{
			AssignPublicIp:   nic.AssignPublicIP.ValueBool(),
			NetworkId:        nic.NetworkID.ValueString(),
			SecurityGroupIds: securityGroupIDS,
		}
	}

	var securityGroupIDs []string
	resp.Diagnostics.Append(state.SecurityGroupIDs.ElementsAs(ctx, &securityGroupIDs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var maxPriceHr *decimal.Decimal
	if !state.MaxPriceHr.IsNull() {
		maxPriceHr = &decimal.Decimal{Value: state.MaxPriceHr.ValueString()}
	}

	params := &vm.CreateVMRequest{
		ProjectId:        projectId,
		BootDisk:         &bootDisk,
		DataCenterId:     state.DataCenterID.ValueString(),
		Gpus:             int32(state.GPUs.ValueInt64()),
		MachineType:      state.MachineType.ValueString(),
		MaxPriceHr:       maxPriceHr,
		MemoryGib:        int32(state.MemoryGib.ValueInt64()),
		Nics:             nics,
		BootDiskImageId:  state.BootDisk.ImageID.ValueString(),
		Password:         state.Password.ValueString(),
		Vcpus:            int32(state.VCPUs.ValueInt64()),
		VmId:             state.ID.ValueString(),
		SecurityGroupIds: securityGroupIDs,
		SshKeySource:     sshKeySource,
		CustomSshKeys:    customKeys,
		StartScript:      state.StartScript.ValueString(),
	}

	_, err := r.client.VMClient.CreateVM(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating VM resource",
			"Could not create VM, unexpected error: "+err.Error(),
		)
		return
		// TODO: sort this out
		// if apiErr, ok := err.(*vm.CreateVMDefault); ok {
		// 	if apiErr.Code() != 409 {
		// 		resp.Diagnostics.AddError(
		// 			"Error creating VM resource",
		// 			"Could not create VM, unexpected error: "+err.Error(),
		// 		)
		// 		return
		// 	}
		// }
	}

	vm, err := waitForVmAvailable(ctx, params.ProjectId, state.ID.ValueString(), r.client.VMClient)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating VM resource",
			"Could not wait for VM resource to become available: "+err.Error(),
		)
		return
	}

	fillState(state, vm.VM)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *VMResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *VMResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := &vm.GetVMRequest{
		ProjectId: r.client.DefaultProjectID,
		Id:        state.ID.ValueString(),
	}

	vm, err := r.client.VMClient.GetVM(ctx, params)
	if err != nil {
		if ok := isErrCode(err, codes.NotFound); ok {
			resp.State.RemoveResource(ctx)
			return

		}
		resp.Diagnostics.AddError(
			"Unable to read VM resource",
			err.Error(),
		)
		return
	}

	fillState(state, vm.VM)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *VMResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan VMResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Error getting vm plan",
			"Error getting vm plan",
		)
		return
	}

	// Read Terraform state data into the model
	var state VMResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *VMResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *VMResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	params := &vm.TerminateVMRequest{}
	params.ProjectId = r.client.DefaultProjectID
	params.Id = state.ID.ValueString()

	if _, err := waitForVmAvailable(ctx, params.ProjectId, params.Id, r.client.VMClient); err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for VM resource to be available",
			err.Error(),
		)
		return
	}

	_, err := r.client.VMClient.TerminateVM(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete VM resource",
			err.Error(),
		)
		return
	}

	_, err = waitForVmDelete(ctx, params.ProjectId, params.Id, r.client.VMClient)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for VM resource to be deleted",
			err.Error(),
		)
		return
	}
}

func (r *VMResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func waitForVmAvailable(ctx context.Context, projectId string, vmID string, c vm.VMServiceClient) (*vm.GetVMResponse, error) {
	refreshFunc := func() (interface{}, string, error) {
		params := &vm.GetVMRequest{
			Id:        vmID,
			ProjectId: projectId,
		}
		res, err := c.GetVM(ctx, params)
		if err != nil {
			if ok := isErrCode(err, codes.NotFound); ok {
				tflog.Debug(ctx, fmt.Sprintf("VM %s in project %s not found: ", vmID, projectId))
				return res, "done", nil
			}
			return nil, "", err
		}

		tflog.Trace(ctx, fmt.Sprintf("pending VM %s in project %s state: %s", vmID, projectId, res.VM.ShortState))
		return res, res.VM.ShortState, nil
	}

	tflog.Debug(ctx, fmt.Sprintf("waiting for VM %s in project %s ", vmID, projectId))

	stateConf := &helper.StateChangeConf{
		Pending:    []string{"boot", "clea", "clon", "dsrz", "epil", "hold", "hotp", "init", "migr", "pend", "prol", "save", "shut", "snap", "unkn"},
		Target:     []string{"done", "fail", "poff", "runn", "stop", "susp", "unde"},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for VM %s in project %s to become available: %w", vmID, projectId, err)
	} else if vm, ok := res.(*vm.GetVMResponse); ok {
		var shortState string
		if vm != nil && vm.VM != nil {
			shortState = vm.VM.ShortState
		}
		tflog.Trace(ctx, fmt.Sprintf("completed waiting for VM %s in project %s (%s)", vmID, projectId, shortState))
		return vm, nil
	} else {
		return nil, fmt.Errorf("error waiting for VM: %v", res)
	}
}

func waitForVmDelete(ctx context.Context, projectId string, vmID string, c vm.VMServiceClient) (*vm.GetVMResponse, error) {
	refreshFunc := func() (interface{}, string, error) {
		params := &vm.GetVMRequest{
			Id:        vmID,
			ProjectId: projectId,
		}
		res, err := c.GetVM(ctx, params)
		if err != nil {
			if ok := isErrCode(err, codes.NotFound); ok {
				tflog.Debug(ctx, fmt.Sprintf("VM %s in project %s is done: ", vmID, projectId))
				return res, "done", nil

			}
			tflog.Error(ctx, fmt.Sprintf("error getting VM %s in project %s: %v", vmID, projectId, err))
			return nil, "", err
		}

		tflog.Trace(ctx, fmt.Sprintf("pending VM %s in project %s state: %s", vmID, projectId, res.VM.ShortState))
		return res, res.VM.ShortState, nil
	}

	tflog.Debug(ctx, fmt.Sprintf("waiting for VM %s in project %s ", vmID, projectId))

	stateConf := &helper.StateChangeConf{
		Pending:    []string{"fail", "poff", "runn", "stop", "susp", "unde", "boot", "clea", "clon", "dsrz", "epil", "hold", "hotp", "init", "migr", "pend", "prol", "save", "shut", "snap", "unkn"},
		Target:     []string{"done"},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for VM %s in project %s to become done: %w", vmID, projectId, err)
	} else if vm, ok := res.(*vm.GetVMResponse); ok {
		var shortState string
		if vm != nil && vm.VM != nil {
			shortState = vm.VM.ShortState
		}
		tflog.Trace(ctx, fmt.Sprintf("completed waiting for VM %s in project %s (%s)", vmID, projectId, shortState))
		return vm, nil
	} else {
		return nil, fmt.Errorf("error waiting for VM: %v", res)
	}
}

// isErrCode checks to see if the provided err is a grpc status with the correct code
func isErrCode(err error, wantCode codes.Code) bool {
	if e, ok := status.FromError(err); ok {
		if e.Code() == wantCode {
			return true
		}
	}
	return false
}

func fillState(state *VMResourceModel, vm *vm.VM) {
	state.DataCenterID = types.StringValue(vm.DatacenterId)
	state.CPUModel = types.StringValue(vm.CpuModel)
	state.GPUs = types.Int64Value(int64(vm.GpuQuantity))
	state.BootDisk.SizeGib = types.Int64Value(int64(vm.BootDiskSizeGib))
	if vm.PublicImageId != "" {
		state.BootDisk.ImageID = types.StringValue(vm.PublicImageId)
	}
	if vm.PrivateImageId != "" {
		state.BootDisk.ImageID = types.StringValue(vm.PrivateImageId)
	}
	state.MachineType = types.StringValue(vm.MachineType)
	for i, nic := range state.Networks {
		nic.ExternalIPAddress = types.StringValue(vm.Nics[i].ExternalIpAddress)
		nic.InternalIPAddress = types.StringValue(vm.Nics[i].InternalIpAddress)
	}
	state.GPUModel = types.StringValue(vm.GpuModel)
	state.ID = types.StringValue(vm.Id)
	state.InternalIPAddress = types.StringValue(vm.InternalIpAddress)
	state.ExternalIPAddress = types.StringValue(vm.ExternalIpAddress)
	state.PriceHr = types.StringValue(fmt.Sprintf("%0.2f", vm.PriceHr))
	state.RenewableEnergy = types.BoolValue(vm.RenewableEnergy)
}
