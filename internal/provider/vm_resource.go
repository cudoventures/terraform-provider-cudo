package provider

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute"
	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/sshkey"
	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/vm"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	"google.golang.org/grpc/codes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VMResource{}
var _ resource.ResourceWithConfigure = &VMResource{}
var _ resource.ResourceWithImportState = &VMResource{}
var _ resource.ResourceWithModifyPlan = &VMResource{}

func NewVMResource() resource.Resource {
	return &VMResource{}
}

func (r *VMResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vm"
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
			"commitment_months": schema.Int32Attribute{
				MarkdownDescription: "The minimum length of time to commit to the VM instance. It cannot be deleted before the commitment end date.",
				Optional:            true,
				Validators:          []validator.Int32{int32validator.OneOf(1, 3, 6, 12, 24, 36)},
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
				MarkdownDescription: "VM machine type",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"memory_gib": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Amount of memory in GiB",
				Optional:            true,
			},
			"metadata": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Metadata values to associate with the VM instance",
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
			"project_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The project the VM instance is in.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid RFC1034 resource id")},
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
			"storage_disks": schema.SetNestedAttribute{
				MarkdownDescription: "Specification for storage disks",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"disk_id": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							MarkdownDescription: "ID of storage disk to attach to vm",
							Required:            true,
							Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
						},
					},
				},
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

// VMResource defines the resource implementation.
type VMResource struct {
	client *CudoClientData
}

// VMResourceModel describes the resource data model.
type VMResourceModel struct {
	BootDisk          *VMBootDiskResourceModel      `tfsdk:"boot_disk"`
	CommitmentMonths  types.Int32                   `tfsdk:"commitment_months"`
	DataCenterID      types.String                  `tfsdk:"data_center_id"`
	ExternalIPAddress types.String                  `tfsdk:"external_ip_address"`
	GPUs              types.Int64                   `tfsdk:"gpus"`
	ID                types.String                  `tfsdk:"id"`
	InternalIPAddress types.String                  `tfsdk:"internal_ip_address"`
	MachineType       types.String                  `tfsdk:"machine_type"`
	MemoryGib         types.Int64                   `tfsdk:"memory_gib"`
	Metadata          types.Map                     `tfsdk:"metadata"`
	Networks          []*VMNICResourceModel         `tfsdk:"networks"`
	Password          types.String                  `tfsdk:"password"`
	ProjectID         types.String                  `tfsdk:"project_id"`
	SecurityGroupIDs  types.Set                     `tfsdk:"security_group_ids"`
	SSHKeys           []types.String                `tfsdk:"ssh_keys"`
	SSHKeySource      types.String                  `tfsdk:"ssh_key_source"`
	StartScript       types.String                  `tfsdk:"start_script"`
	StorageDisks      []*VMStorageDiskResourceModel `tfsdk:"storage_disks"`
	VCPUs             types.Int64                   `tfsdk:"vcpus"`
}

type VMBootDiskResourceModel struct {
	ImageID types.String `tfsdk:"image_id"`
	SizeGib types.Int64  `tfsdk:"size_gib"`
}

type VMNICResourceModel struct {
	AssignPublicIP    types.Bool   `tfsdk:"assign_public_ip"`
	ExternalIPAddress types.String `tfsdk:"external_ip_address"`
	InternalIPAddress types.String `tfsdk:"internal_ip_address"`
	NetworkID         types.String `tfsdk:"network_id"`
	SecurityGroupIDs  types.Set    `tfsdk:"security_group_ids"`
}

type VMStorageDiskResourceModel struct {
	DiskID types.String `tfsdk:"disk_id"`
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

func (r *VMResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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

func (r *VMResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VMResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sshKeySource := sshkey.SshKeySource_SSH_KEY_SOURCE_PROJECT
	switch plan.SSHKeySource.ValueString() {
	case "personal", "user":
		sshKeySource = sshkey.SshKeySource_SSH_KEY_SOURCE_USER
	case "custom", "none":
		sshKeySource = sshkey.SshKeySource_SSH_KEY_SOURCE_NONE
	}

	var customKeys []string
	if sshKeySource == sshkey.SshKeySource_SSH_KEY_SOURCE_NONE {
		for _, key := range plan.SSHKeys {
			customKeys = append(customKeys, key.ValueString())
		}
	}

	var bootDisk vm.Disk
	if !plan.BootDisk.SizeGib.IsNull() {
		sizeGib := int32(plan.BootDisk.SizeGib.ValueInt64())
		bootDisk.SizeGib = sizeGib
	}

	nics := make([]*vm.CreateVMRequest_NIC, len(plan.Networks))
	for i, nic := range plan.Networks {
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
	resp.Diagnostics.Append(plan.SecurityGroupIDs.ElementsAs(ctx, &securityGroupIDs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	storageDiskIds := make([]string, len(plan.StorageDisks))
	for i, diskResource := range plan.StorageDisks {
		storageDiskIds[i] = diskResource.DiskID.ValueString()
	}

	metadataMap := make(map[string]string)
	diag := plan.Metadata.ElementsAs(ctx, &metadataMap, false)
	if diag.HasError() {
		resp.Diagnostics.Append(diag...)
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

	params := &vm.CreateVMRequest{
		BootDisk:         &bootDisk,
		BootDiskImageId:  plan.BootDisk.ImageID.ValueString(),
		CommitmentTerm:   commitmentTerm,
		CustomSshKeys:    customKeys,
		DataCenterId:     plan.DataCenterID.ValueString(),
		Gpus:             int32(plan.GPUs.ValueInt64()),
		MachineType:      plan.MachineType.ValueString(),
		MemoryGib:        int32(plan.MemoryGib.ValueInt64()),
		Metadata:         metadataMap,
		Nics:             nics,
		Password:         plan.Password.ValueString(),
		ProjectId:        plan.ProjectID.ValueString(),
		SecurityGroupIds: securityGroupIDs,
		SshKeySource:     sshKeySource,
		StartScript:      plan.StartScript.ValueString(),
		StorageDiskIds:   storageDiskIds,
		Vcpus:            int32(plan.VCPUs.ValueInt64()),
		VmId:             plan.ID.ValueString(),
	}

	_, err := r.client.VMClient.CreateVM(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create VM instance",
			err.Error(),
		)
		return
	}

	vm, err := waitForVmAvailable(ctx, r.client.VMClient, plan.ProjectID.ValueString(), plan.ID.ValueString(), plan.StorageDisks)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for VM to become available",
			err.Error(),
		)
		return
	}

	// if the vm is created and returned update the state.
	resp.Diagnostics.Append(appendVmState(vm, &plan)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *VMResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VMResourceModel

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
		if ok := helper.IsErrCode(err, codes.NotFound); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to read VM instance",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(appendVmState(vm.VM, &state)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *VMResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Unable to update VM instance",
		"Updating a VM instance is not supported",
	)
}

func (r *VMResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VMResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	projectID := state.ProjectID.ValueString()
	if state.ProjectID.IsNull() {
		projectID = r.client.DefaultProjectID
	}
	vmID := state.ID.ValueString()

	_, err := r.client.VMClient.TerminateVM(ctx, &vm.TerminateVMRequest{
		ProjectId: projectID,
		Id:        vmID,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete VM resource",
			err.Error(),
		)
		return
	}

	_, err = waitForVmDelete(ctx, r.client.VMClient, projectID, vmID)
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

func waitForVmAvailable(ctx context.Context, c vm.VMServiceClient, projectId string, vmID string, storageDisks []*VMStorageDiskResourceModel) (*vm.VM, error) {
	refreshFunc := func() (interface{}, string, error) {
		params := &vm.GetVMRequest{
			Id:        vmID,
			ProjectId: projectId,
		}
		res, err := c.GetVM(ctx, params)
		if err != nil {
			return res, "target", err
		}
		// if there are storage disks in the plan, wait for them to attach to the vm
		if len(res.VM.StorageDisks) != len(storageDisks) {
			return res, "pending", nil
		}
		switch res.VM.State {
		case vm.VM_ACTIVE,
			vm.VM_DELETED,
			vm.VM_FAILED,
			vm.VM_STOPPED,
			vm.VM_SUSPENDED:
			return res, "target", nil
		default:
			return res, "pending", nil
		}
	}

	stateConf := &helper.StateChangeConf{
		Pending:    []string{"pending"},
		Target:     []string{"target"},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for VM %s in project %s to become available: %w", vmID, projectId, err)
	} else if vm, ok := res.(*vm.GetVMResponse); ok {
		return vm.VM, nil
	} else {
		return nil, fmt.Errorf("error waiting for VM: %v", res)
	}
}

func waitForVmDelete(ctx context.Context, c vm.VMServiceClient, projectId string, vmID string) (*vm.GetVMResponse, error) {
	refreshFunc := func() (interface{}, string, error) {
		res, err := c.GetVM(ctx, &vm.GetVMRequest{
			Id:        vmID,
			ProjectId: projectId,
		})
		if err != nil {
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				return res, "deleted", nil
			}
			return nil, "unknown", err
		}

		return res, "pending", nil
	}

	stateConf := &helper.StateChangeConf{
		Pending:    []string{"pending"},
		Target:     []string{"deleted", "unknown"},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for VM %s in project %s to become done: %w", vmID, projectId, err)
	} else if vm, ok := res.(*vm.GetVMResponse); ok {
		return vm, nil
	} else {
		return nil, fmt.Errorf("error waiting for VM: %v", res)
	}
}

func appendVmState(instance *vm.VM, state *VMResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	state.BootDisk.SizeGib = types.Int64Value(int64(instance.BootDiskSizeGib))
	if instance.PublicImageId != "" {
		state.BootDisk.ImageID = types.StringValue(instance.PublicImageId)
	}
	if instance.PrivateImageId != "" {
		state.BootDisk.ImageID = types.StringValue(instance.PrivateImageId)
	}

	switch instance.CommitmentTerm {
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

	state.DataCenterID = types.StringValue(instance.DatacenterId)
	state.ExternalIPAddress = types.StringValue(instance.ExternalIpAddress)
	state.GPUs = types.Int64Value(int64(instance.GpuQuantity))
	state.ID = types.StringValue(instance.Id)
	state.InternalIPAddress = types.StringValue(instance.InternalIpAddress)
	state.MachineType = types.StringValue(instance.MachineType)
	state.MemoryGib = types.Int64Value(int64(instance.Memory))

	var networks []*VMNICResourceModel
	for _, network := range instance.Nics {
		securityGroupIDs := make([]attr.Value, 0, len(network.SecurityGroupIds))
		for _, securityGroupID := range network.SecurityGroupIds {
			securityGroupIDs = append(securityGroupIDs, types.StringValue(securityGroupID))
		}
		securityGroupSetValue := types.SetNull(types.StringType)
		if len(securityGroupIDs) > 0 {
			var d diag.Diagnostics
			securityGroupSetValue, d = types.SetValue(types.StringType, securityGroupIDs)
			diags.Append(d...)
		}

		networks = append(networks, &VMNICResourceModel{
			AssignPublicIP:    types.BoolNull(),
			ExternalIPAddress: types.StringValue(network.ExternalIpAddress),
			InternalIPAddress: types.StringValue(network.InternalIpAddress),
			NetworkID:         types.StringValue(network.NetworkId),
			SecurityGroupIDs:  securityGroupSetValue,
		})
	}
	state.Networks = networks

	for i, nic := range state.Networks {
		nic.ExternalIPAddress = types.StringValue(instance.Nics[i].ExternalIpAddress)
		nic.InternalIPAddress = types.StringValue(instance.Nics[i].InternalIpAddress)
	}

	state.ProjectID = types.StringValue(instance.ProjectId)

	var storageDisks []*VMStorageDiskResourceModel
	for _, vmDisk := range instance.StorageDisks {
		storageDisks = append(storageDisks, &VMStorageDiskResourceModel{
			DiskID: types.StringValue(vmDisk.Id),
		})
	}
	state.StorageDisks = storageDisks

	mdElems := make(map[string]attr.Value)
	for k, v := range instance.Metadata {
		mdElems[k] = types.StringValue(v)
	}
	if len(mdElems) > 0 {
		md, mapDiags := types.MapValue(types.StringType, mdElems)
		diags.Append(mapDiags...)
		state.Metadata = md
	}

	switch instance.SshKeySource {
	case sshkey.SshKeySource_SSH_KEY_SOURCE_NONE:
		state.SSHKeySource = types.StringValue("none")
		for authorizedKey := range strings.Lines(instance.AuthorizedSshKeys) {
			state.SSHKeys = append(state.SSHKeys, types.StringValue(authorizedKey))
		}
	case sshkey.SshKeySource_SSH_KEY_SOURCE_PROJECT,
		sshkey.SshKeySource_SSH_KEY_SOURCE_UNKNOWN:
		state.SSHKeySource = types.StringNull()
	case sshkey.SshKeySource_SSH_KEY_SOURCE_USER:
		state.SSHKeySource = types.StringValue("user")
	}

	state.VCPUs = types.Int64Value(int64(instance.Vcpus))

	return diags
}
