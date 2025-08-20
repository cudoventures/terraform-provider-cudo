package provider

import (
	"context"
	"fmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/vm"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &VMDataSource{}

func NewVMDataSource() datasource.DataSource {
	return &VMDataSource{}
}

// VMDataSource defines the data source implementation.
type VMDataSource struct {
	client *CudoClientData
}

type VMDataSourceModel struct {
	BootDiskSizeGib   types.Int64  `tfsdk:"boot_disk_size_gib"`
	DatacenterID      types.String `tfsdk:"data_center_id"`
	ExternalIPAddress types.String `tfsdk:"external_ip_address"`
	Gpus              types.Int64  `tfsdk:"gpus"`
	Id                types.String `tfsdk:"id"`
	ImageID           types.String `tfsdk:"image_id"`
	InternalIPAddress types.String `tfsdk:"internal_ip_address"`
	MachineType       types.String `tfsdk:"machine_type"`
	MemoryGib         types.Int64  `tfsdk:"memory_gib"`
	Metadata          types.Map    `tfsdk:"metadata"`
	ProjectID         types.String `tfsdk:"project_id"`
	Vcpus             types.Int64  `tfsdk:"vcpus"`
}

func (d *VMDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vm"
}

func (d *VMDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "VM data source",
		Description:         "Gets a VM",
		Attributes: map[string]schema.Attribute{
			"boot_disk_size_gib": schema.Int64Attribute{
				MarkdownDescription: "The size of the boot disk in gibibytes (GiB).",
				Computed:            true,
			},
			"data_center_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the datacenter where the VM instance is located.",
				Computed:            true,
			},
			"external_ip_address": schema.StringAttribute{
				MarkdownDescription: "The external IP address of the VM instance.",
				Computed:            true,
			},
			"gpus": schema.Int64Attribute{
				MarkdownDescription: "The number of GPUs attached to the VM instance.",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the VM instance.",
				Required:            true,
			},
			"image_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the image used to create the VM instance.",
				Computed:            true,
			},
			"internal_ip_address": schema.StringAttribute{
				MarkdownDescription: "The internal IP address of the VM instance.",
				Computed:            true,
			},
			"machine_type": schema.StringAttribute{
				MarkdownDescription: "The machine type of the VM instance",
				Computed:            true,
			},
			"memory_gib": schema.Int64Attribute{
				MarkdownDescription: "The amount of memory allocated to the VM instance.",
				Computed:            true,
			},
			"metadata": schema.MapAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "The key value pairs to associate with the VM.",
				Optional:            true,
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the project the VM is in.",
				Optional:            true,
			},
			"vcpus": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
		},
	}
}

func (d *VMDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*CudoClientData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *CudoClientData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

func (d *VMDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state VMDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectID := state.ProjectID.ValueString()
	if state.ProjectID.IsNull() {
		state.ProjectID = types.StringValue(d.client.DefaultProjectID)
		projectID = d.client.DefaultProjectID
	}

	ID := state.Id.ValueString()

	vm, err := d.client.VMClient.GetVM(ctx, &vm.GetVMRequest{
		ProjectId: projectID,
		Id:        ID,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read VM instance",
			err.Error(),
		)
		return
	}

	imageID := vm.VM.PrivateImageId
	if imageID == "" {
		imageID = vm.VM.PublicImageId
	}

	metadataMap, diag := types.MapValueFrom(ctx, types.StringType, vm.VM.Metadata)
	resp.Diagnostics.Append(diag...)

	state.BootDiskSizeGib = types.Int64Value(int64(vm.VM.BootDiskSizeGib))
	state.DatacenterID = types.StringValue(vm.VM.DatacenterId)
	state.ExternalIPAddress = types.StringValue(vm.VM.ExternalIpAddress)
	state.Gpus = types.Int64Value(int64(vm.VM.GpuQuantity))
	state.ImageID = types.StringValue(imageID)
	state.InternalIPAddress = types.StringValue(vm.VM.InternalIpAddress)
	state.MachineType = types.StringValue(vm.VM.MachineType)
	state.MemoryGib = types.Int64Value(int64(vm.VM.Memory))
	state.Metadata = metadataMap
	state.Vcpus = types.Int64Value(int64(vm.VM.Vcpus))

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
