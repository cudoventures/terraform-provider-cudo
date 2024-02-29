package provider

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client/virtual_machines"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"
	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
	"github.com/hashicorp/go-retryablehttp"
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

func waitForVmAvailable(ctx context.Context, projectID string, vmID string, c virtual_machines.ClientService) (*virtual_machines.GetVMOK, error) {
	refreshFunc := func() (interface{}, string, error) {
		params := virtual_machines.NewGetVMParamsWithContext(ctx)
		params.ID = vmID
		params.ProjectID = projectID
		res, err := c.GetVM(params)
		if err != nil {
			if apiErr, ok := err.(*virtual_machines.GetVMDefault); ok && apiErr.IsCode(404) {
				tflog.Debug(ctx, fmt.Sprintf("VM %s in project %s not found: ", vmID, projectID))
				return res, "done", nil
			}
			return nil, "", err
		}

		tflog.Trace(ctx, fmt.Sprintf("pending VM %s in project %s state: %s", vmID, projectID, res.Payload.VM.ShortState))
		return res, res.Payload.VM.ShortState, nil
	}

	tflog.Debug(ctx, fmt.Sprintf("waiting for VM %s in project %s ", vmID, projectID))

	stateConf := &helper.StateChangeConf{
		Pending:    []string{"boot", "clea", "clon", "dsrz", "epil", "hold", "hotp", "init", "migr", "pend", "prol", "save", "shut", "snap", "unkn"},
		Target:     []string{"done", "fail", "poff", "runn", "stop", "susp", "unde"},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for VM %s in project %s to become available: %w", vmID, projectID, err)
	} else if vm, ok := res.(*virtual_machines.GetVMOK); ok {
		var shortState string
		if vm != nil && vm.Payload != nil && vm.Payload.VM != nil {
			shortState = vm.Payload.VM.ShortState
		}
		tflog.Trace(ctx, fmt.Sprintf("completed waiting for VM %s in project %s (%s)", vmID, projectID, shortState))
		return vm, nil
	} else {
		return nil, fmt.Errorf("error waiting for VM: %v", res)
	}
}

func waitForVmDelete(ctx context.Context, projectID string, vmID string, c virtual_machines.ClientService) (*virtual_machines.GetVMOK, error) {
	refreshFunc := func() (interface{}, string, error) {
		params := virtual_machines.NewGetVMParamsWithContext(ctx)
		params.ID = vmID
		params.ProjectID = projectID
		res, err := c.GetVM(params)
		if err != nil {
			if apiErr, ok := err.(*virtual_machines.GetVMDefault); ok && apiErr.IsCode(404) {
				tflog.Debug(ctx, fmt.Sprintf("VM %s in project %s is done: ", vmID, projectID))
				return res, "done", nil
			}
			tflog.Error(ctx, fmt.Sprintf("error getting VM %s in project %s: %v", vmID, projectID, err))
			return nil, "", err
		}

		tflog.Trace(ctx, fmt.Sprintf("pending VM %s in project %s state: %s", vmID, projectID, res.Payload.VM.ShortState))
		return res, res.Payload.VM.ShortState, nil
	}

	tflog.Debug(ctx, fmt.Sprintf("waiting for VM %s in project %s ", vmID, projectID))

	stateConf := &helper.StateChangeConf{
		Pending:    []string{"fail", "poff", "runn", "stop", "susp", "unde", "boot", "clea", "clon", "dsrz", "epil", "hold", "hotp", "init", "migr", "pend", "prol", "save", "shut", "snap", "unkn"},
		Target:     []string{"done"},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for VM %s in project %s to become done: %w", vmID, projectID, err)
	} else if vm, ok := res.(*virtual_machines.GetVMOK); ok {
		var shortState string
		if vm != nil && vm.Payload != nil && vm.Payload.VM != nil {
			shortState = vm.Payload.VM.ShortState
		}
		tflog.Trace(ctx, fmt.Sprintf("completed waiting for VM %s in project %s (%s)", vmID, projectID, shortState))
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

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	params := virtual_machines.NewCreateVMParamsWithContext(ctx).WithHTTPClient(retryClient.StandardClient())
	params.ProjectID = r.client.DefaultProjectID
	if !state.ProjectID.IsNull() {
		params.ProjectID = state.ProjectID.ValueString()
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
	params.Body = &models.CreateVMBody{
		BootDisk:         &bootDisk,
		DataCenterID:     state.DataCenterID.ValueString(),
		Gpus:             int32(state.GPUs.ValueInt64()),
		MachineType:      state.MachineType.ValueString(),
		MaxPriceHr:       maxPriceHr,
		MemoryGib:        int32(state.MemoryGib.ValueInt64()),
		Nics:             nics,
		BootDiskImageID:  state.BootDisk.ImageID.ValueStringPointer(),
		Password:         state.Password.ValueString(),
		Vcpus:            int32(state.VCPUs.ValueInt64()),
		VMID:             state.ID.ValueStringPointer(),
		SecurityGroupIds: securityGroupIDs,
		SSHKeySource:     models.SSHKeySource(sshKeySource).Pointer(),
		CustomSSHKeys:    customKeys,
		StartScript:      state.StartScript.ValueString(),
	}

	_, err := r.client.Client.VirtualMachines.CreateVM(params)
	if err != nil {
		if apiErr, ok := err.(*virtual_machines.CreateVMDefault); ok {
			if apiErr.Code() != 409 {
				resp.Diagnostics.AddError(
					"Error creating VM resource",
					"Could not create VM, unexpected error: "+err.Error(),
				)
				return
			}
		}
	}

	vm, err := waitForVmAvailable(ctx, params.ProjectID, state.ID.ValueString(), r.client.Client.VirtualMachines)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating VM resource",
			"Could not wait for VM resource to become available: "+err.Error(),
		)
		return
	}

	state.DataCenterID = types.StringValue(vm.Payload.VM.DatacenterID)
	state.CPUModel = types.StringValue(vm.Payload.VM.CPUModel)
	state.GPUs = types.Int64Value(vm.Payload.VM.GpuQuantity)
	state.BootDisk.SizeGib = types.Int64Value(vm.Payload.VM.BootDiskSizeGib)
	if vm.Payload.VM.PublicImageID != "" {
		state.BootDisk.ImageID = types.StringValue(vm.Payload.VM.PublicImageID)
	}
	if vm.Payload.VM.PrivateImageID != "" {
		state.BootDisk.ImageID = types.StringValue(vm.Payload.VM.PrivateImageID)
	}
	state.MachineType = types.StringValue(vm.Payload.VM.MachineType)
	for i, nic := range state.Networks {
		nic.ExternalIPAddress = types.StringValue(vm.Payload.VM.Nics[i].ExternalIPAddress)
		nic.InternalIPAddress = types.StringValue(vm.Payload.VM.Nics[i].InternalIPAddress)
	}
	state.GPUModel = types.StringValue(vm.Payload.VM.GpuModel)
	state.ID = types.StringValue(vm.Payload.VM.ID)
	state.InternalIPAddress = types.StringValue(vm.Payload.VM.InternalIPAddress)
	state.ExternalIPAddress = types.StringValue(vm.Payload.VM.ExternalIPAddress)
	state.PriceHr = types.StringValue(fmt.Sprintf("%0.2f", vm.Payload.VM.PriceHr))
	state.RenewableEnergy = types.BoolValue(vm.Payload.VM.RenewableEnergy)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *VMResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *VMResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := virtual_machines.NewGetVMParamsWithContext(ctx)
	params.ProjectID = r.client.DefaultProjectID
	params.ID = state.ID.ValueString()

	res, err := r.client.Client.VirtualMachines.GetVM(params)
	if err != nil {
		if apiErr, ok := err.(*virtual_machines.GetVMDefault); ok && apiErr.IsCode(404) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to read VM resource",
			err.Error(),
		)
		return
	}

	state.DataCenterID = types.StringValue(res.Payload.VM.DatacenterID)
	state.CPUModel = types.StringValue(res.Payload.VM.CPUModel)
	state.GPUs = types.Int64Value(res.Payload.VM.GpuQuantity)
	state.BootDisk.SizeGib = types.Int64Value(res.Payload.VM.BootDiskSizeGib)
	if res.Payload.VM.PublicImageID != "" {
		state.BootDisk.ImageID = types.StringValue(res.Payload.VM.PublicImageID)
	}
	if res.Payload.VM.PrivateImageID != "" {
		state.BootDisk.ImageID = types.StringValue(res.Payload.VM.PrivateImageID)
	}
	state.MachineType = types.StringValue(res.Payload.VM.MachineType)
	for i, nic := range state.Networks {
		nic.ExternalIPAddress = types.StringValue(res.Payload.VM.Nics[i].ExternalIPAddress)
		nic.InternalIPAddress = types.StringValue(res.Payload.VM.Nics[i].InternalIPAddress)
	}
	state.GPUModel = types.StringValue(res.Payload.VM.GpuModel)
	state.ID = types.StringValue(res.Payload.VM.ID)
	state.InternalIPAddress = types.StringValue(res.Payload.VM.InternalIPAddress)
	state.ExternalIPAddress = types.StringValue(res.Payload.VM.ExternalIPAddress)
	state.PriceHr = types.StringValue(fmt.Sprintf("%0.2f", res.Payload.VM.PriceHr))
	state.RenewableEnergy = types.BoolValue(res.Payload.VM.RenewableEnergy)

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

	params := virtual_machines.NewTerminateVMParamsWithContext(ctx)
	params.ProjectID = r.client.DefaultProjectID
	params.ID = state.ID.ValueString()

	if _, err := waitForVmAvailable(ctx, params.ProjectID, params.ID, r.client.Client.VirtualMachines); err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for VM resource to be available",
			err.Error(),
		)
		return
	}

	_, err := r.client.Client.VirtualMachines.TerminateVM(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete VM resource",
			err.Error(),
		)
		return
	}

	_, err = waitForVmDelete(ctx, params.ProjectID, params.ID, r.client.Client.VirtualMachines)
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
