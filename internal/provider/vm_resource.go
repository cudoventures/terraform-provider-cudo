package provider

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client/compute/vm"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"
	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

func (r *VMResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "cudo_vm"
}

func (r *VMResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "VM resource",
		Attributes: map[string]schema.Attribute{
			"boot_disk": schema.SingleNestedAttribute{
				MarkdownDescription: "Specification for boot disk",
				Attributes: map[string]schema.Attribute{
					"size_gib": schema.Int64Attribute{
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
						Computed:            true,
						Optional:            true,
						MarkdownDescription: "Size of boot disk in Gib",
					},
					"image_id": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						MarkdownDescription: "ID of OS image on boot disk",
						Required:            true,
						Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
					},
				},
				Required: true,
			},
			"cpu_model": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The model of the CPU.",
				Optional:            true,
				Computed:            true,
			},
			"data_center_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The id of the datacenter where the VM instance is located.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"external_ip_address": schema.StringAttribute{
				MarkdownDescription: "The external IP address of the VM instance.",
				Computed:            true,
			},
			"gpu_model": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The model of the GPU.",
				Optional:            true,
				Computed:            true,
			},
			"gpus": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Number of GPUs",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
			},
			"id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "ID for VM within project",
				Required:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id e.g. my-vm")},
			},
			"internal_ip_address": schema.StringAttribute{
				MarkdownDescription: "The internal IP address of the VM instance.",
				Computed:            true,
			},
			"machine_type": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "VM machine type, from machine type data source",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"max_price_hr": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The maximum price per hour for the VM instance.",
				Optional:            true,
			},
			"memory_gib": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Amount of VM memory in GiB",
				Optional:            true,
			},
			"networks": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Network adapters for private networks",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_id": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							MarkdownDescription: "ID of private network to attach the NIC to",
							Required:            true,
						},
						"assign_public_ip": schema.BoolAttribute{
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.RequiresReplace(),
							},
							MarkdownDescription: "Assign a public IP to the NIC",
							Optional:            true,
						},
						"external_ip_address": schema.StringAttribute{
							MarkdownDescription: "The external IP address of the NIC.",
							Computed:            true,
						},
						"internal_ip_address": schema.StringAttribute{
							MarkdownDescription: "The internal IP address of the NIC.",
							Computed:            true,
						},
						"security_group_ids": schema.SetAttribute{
							ElementType:         types.StringType,
							Optional:            true,
							MarkdownDescription: "Security groups to assign to the NIC",
						},
					},
				},
			},
			"password": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Root password for linux, or Admin password for windows",
				Optional:            true,
				Sensitive:           true,
				Validators:          []validator.String{stringvalidator.LengthBetween(6, 64)},
			},
			"price_hr": schema.StringAttribute{
				MarkdownDescription: "The current price per hour for the VM instance.",
				Computed:            true,
			},
			"project_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The project the VM instance is in.",
				Optional:            true,
			},
			"renewable_energy": schema.BoolAttribute{
				MarkdownDescription: "Whether the VM instance is powered by renewable energy",
				Computed:            true,
			},
			"security_group_ids": schema.SetAttribute{
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Security groups to assign to the VM when using public networking",
			},
			"ssh_key_source": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Which SSH keys to add to the VM: project (default), user or custom",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("project", "user", "custom")},
			},
			"ssh_keys": schema.ListAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				ElementType:         types.StringType,
				MarkdownDescription: "List of SSH keys to add to the VM, ssh_key_source must be set to custom",
				Optional:            true,
			},
			"start_script": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "A script to run when VM boots",
				Optional:            true,
			},
			"vcpus": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Number of VCPUs",
				Optional:            true,
				Validators:          []validator.Int64{int64validator.AtMost(100)},
			},
		},
	}
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

func waitForVmAvailable(ctx context.Context, projectId string, vmID string, c vm.VMServiceClient) (*vm.GetVMResponse, error) {
	refreshFunc := func() (interface{}, string, error) {
		params := &vm.GetVMRequest{
			Id:        vmID,
			ProjectId: projectId,
		}
		res, err := c.GetVM(ctx, params)
		if ok, err := isErrCode(err, codes.NotFound); err != nil {
			if ok {
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
		if vm != nil && vm != nil && vm.VM != nil {
			shortState = vm.VM.ShortState
		}
		tflog.Trace(ctx, fmt.Sprintf("completed waiting for VM %s in project %s (%s)", vmID, projectId, shortState))
		return vm, nil
	} else {
		return nil, fmt.Errorf("error waiting for VM: %v", res)
	}
}

func isErrCode(err error, wantCode codes.Code) (bool, error) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == wantCode {
				return true, err
			}
		}

		return false, err
	}
	return false, nil
}

func waitForVmDelete(ctx context.Context, projectId string, vmID string, c vm.VMServiceClient) (*vm.GetVMResponse, error) {
	refreshFunc := func() (interface{}, string, error) {
		params := &vm.GetVMRequest{
			Id:        vmID,
			ProjectId: projectId,
		}
		res, err := c.GetVM(ctx, params)
		if ok, err := isErrCode(err, codes.NotFound); err != nil {
			if ok {
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
		if vm != nil && vm != nil && vm.VM != nil {
			shortState = vm.VM.ShortState
		}
		tflog.Trace(ctx, fmt.Sprintf("completed waiting for VM %s in project %s (%s)", vmID, projectId, shortState))
		return vm, nil
	} else {
		return nil, fmt.Errorf("error waiting for VM: %v", res)
	}
}

func (r *VMResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state *VMResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	sshKeySource := models.SSHKeySourceSSHKEYSOURCEPROJECT
	switch state.SSHKeySource.ValueString() {
	case "user":
		sshKeySource = models.SSHKeySourceSSHKEYSOURCEUSER
	case "custom":
		sshKeySource = models.SSHKeySourceSSHKEYSOURCENONE
	}

	var customKeys []string
	if sshKeySource == models.SSHKeySourceSSHKEYSOURCENONE {
		for _, key := range state.SSHKeys {
			customKeys = append(customKeys, key.ValueString())
		}
	}

	// retryClient := retryablehttp.NewClient()
	// retryClient.RetryMax = 10
	params := &vm.CreateVMRequest{}
	params.ProjectId = r.client.DefaultProjectID
	if !state.ProjectID.IsNull() {
		params.ProjectId = state.ProjectID.ValueString()
	}

	var bootDisk models.Disk
	if !state.BootDisk.SizeGib.IsNull() {
		sizeGib := int32(state.BootDisk.SizeGib.ValueInt64())
		bootDisk.SizeGib = &sizeGib
	}
	nics := make([]*models.CreateVMRequestNIC, len(state.Networks))

	for i, nic := range state.Networks {
		var securityGroupIDS []string
		if !nic.SecurityGroupIDs.IsNull() {
			resp.Diagnostics.Append(nic.SecurityGroupIDs.ElementsAs(ctx, &securityGroupIDS, false)...)
			if resp.Diagnostics.HasError() {
				return
			}
		}
		nics[i] = &models.CreateVMRequestNIC{
			AssignPublicIP:   nic.AssignPublicIP.ValueBool(),
			NetworkID:        nic.NetworkID.ValueString(),
			SecurityGroupIds: securityGroupIDS,
		}
	}

	var securityGroupIDs []string
	resp.Diagnostics.Append(state.SecurityGroupIDs.ElementsAs(ctx, &securityGroupIDs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var maxPriceHr *models.Decimal
	if !state.MaxPriceHr.IsNull() {
		maxPriceHr = &models.Decimal{Value: state.MaxPriceHr.ValueString()}
	}

	params.BootDisk = &bootDisk
	params.DataCenterId = state.DataCenterID.ValueString()
	params.Gpus = int32(state.GPUs.ValueInt64())
	params.MachineType = state.MachineType.ValueString()
	params.MaxPriceHr = maxPriceHr
	params.MemoryGib = int32(state.MemoryGib.ValueInt64())
	params.Nics = nics
	params.BootDiskImageId = state.BootDisk.ImageID.ValueString()
	params.Password = state.Password.ValueString()
	params.Vcpus = int32(state.VCPUs.ValueInt64())
	params.VmId = state.ID.ValueString()
	params.SecurityGroupIds = securityGroupIDs
	params.SshKeySource = models.SSHKeySource(sshKeySource).Pointer()
	params.CustomSshKeys = customKeys
	params.StartScript = state.StartScript.ValueString()

	_, err := r.client.VMClient.CreateVM(ctx, params)
	if err != nil {
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

	state.DataCenterID = types.StringValue(vm.VM.DatacenterId)
	state.CPUModel = types.StringValue(vm.VM.CpuModel)
	state.GPUs = types.Int64Value(int64(vm.VM.GpuQuantity))
	state.BootDisk.SizeGib = types.Int64Value(int64(vm.VM.BootDiskSizeGib))
	if vm.VM.PublicImageId != "" {
		state.BootDisk.ImageID = types.StringValue(vm.VM.PublicImageId)
	}
	if vm.VM.PrivateImageId != "" {
		state.BootDisk.ImageID = types.StringValue(vm.VM.PrivateImageId)
	}
	state.MachineType = types.StringValue(vm.VM.MachineType)
	for i, nic := range state.Networks {
		nic.ExternalIPAddress = types.StringValue(vm.VM.Nics[i].ExternalIpAddress)
		nic.InternalIPAddress = types.StringValue(vm.VM.Nics[i].InternalIpAddress)
	}
	state.GPUModel = types.StringValue(vm.VM.GpuModel)
	state.ID = types.StringValue(vm.VM.Id)
	state.InternalIPAddress = types.StringValue(vm.VM.InternalIpAddress)
	state.ExternalIPAddress = types.StringValue(vm.VM.ExternalIpAddress)
	state.PriceHr = types.StringValue(fmt.Sprintf("%0.2f", vm.VM.PriceHr))
	state.RenewableEnergy = types.BoolValue(vm.VM.RenewableEnergy)

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

	res, err := r.client.VMClient.GetVM(ctx, params)
	if ok, err := isErrCode(err, codes.NotFound); err != nil {
		if ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to read VM resource",
			err.Error(),
		)
		return
	}

	state.DataCenterID = types.StringValue(res.VM.DatacenterId)
	state.CPUModel = types.StringValue(res.VM.CpuModel)
	state.GPUs = types.Int64Value(int64(res.VM.GpuQuantity))
	state.BootDisk.SizeGib = types.Int64Value(int64(res.VM.BootDiskSizeGib))
	if res.VM.PublicImageId != "" {
		state.BootDisk.ImageID = types.StringValue(res.VM.PublicImageId)
	}
	if res.VM.PrivateImageId != "" {
		state.BootDisk.ImageID = types.StringValue(res.VM.PrivateImageId)
	}
	state.MachineType = types.StringValue(res.VM.MachineType)
	for i, nic := range state.Networks {
		nic.ExternalIPAddress = types.StringValue(res.VM.Nics[i].ExternalIpAddress)
		nic.InternalIPAddress = types.StringValue(res.VM.Nics[i].InternalIpAddress)
	}
	state.GPUModel = types.StringValue(res.VM.GpuModel)
	state.ID = types.StringValue(res.VM.Id)
	state.InternalIPAddress = types.StringValue(res.VM.InternalIpAddress)
	state.ExternalIPAddress = types.StringValue(res.VM.ExternalIpAddress)
	state.PriceHr = types.StringValue(fmt.Sprintf("%0.2f", res.VM.PriceHr))
	state.RenewableEnergy = types.BoolValue(res.VM.RenewableEnergy)

	// Save updated data into Terraform state
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
