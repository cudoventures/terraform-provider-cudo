package provider

import (
	"context"
	"fmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/vm"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	Id                types.String  `tfsdk:"id"`
	BootDiskSizeGib   types.Int64   `tfsdk:"boot_disk_size_gib"`
	CPUModel          types.String  `tfsdk:"cpu_model"`
	DatacenterID      types.String  `tfsdk:"data_center_id"`
	GpuModel          types.String  `tfsdk:"gpu_model"`
	Gpus              types.Int64   `tfsdk:"gpus"`
	ImageID           types.String  `tfsdk:"image_id"`
	InternalIPAddress types.String  `tfsdk:"internal_ip_address"`
	ExternalIPAddress types.String  `tfsdk:"external_ip_address"`
	Memory            types.Int64   `tfsdk:"memory_gib"`
	Metadata          types.Map     `tfsdk:"metadata"`
	PriceHr           types.Float64 `tfsdk:"price_hr"`
	ProjectID         types.String  `tfsdk:"project_id"`
	Vcpus             types.Int64   `tfsdk:"vcpus"`
}

func (d *VMDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "cudo_vm"
}

func (d *VMDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "VM data source",
		Description:         "Gets a VM",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the VM instance.",
				Required:            true,
			},
			"boot_disk_size_gib": schema.Int64Attribute{
				MarkdownDescription: "The size of the boot disk in gibibytes (GiB).",
				Computed:            true,
			},
			"cpu_model": schema.StringAttribute{
				MarkdownDescription: "The model of the CPU.",
				Computed:            true,
			},
			"data_center_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the datacenter where the VM instance is located.",
				Computed:            true,
			},
			"gpu_model": schema.StringAttribute{
				MarkdownDescription: "The model of the GPU.",
				Computed:            true,
			},
			"gpus": schema.Int64Attribute{
				MarkdownDescription: "The number of GPUs attached to the VM instance.",
				Computed:            true,
			},
			"image_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the image used to create the VM instance.",
				Computed:            true,
			},
			"internal_ip_address": schema.StringAttribute{
				MarkdownDescription: "The internal IP address of the VM instance.",
				Computed:            true,
			},
			"external_ip_address": schema.StringAttribute{
				MarkdownDescription: "The external IP address of the VM instance.",
				Computed:            true,
			},
			"memory_gib": schema.Int64Attribute{
				MarkdownDescription: "The amount of memory allocated to the VM instance.",
				Computed:            true,
			},
			"metadata": schema.MapAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "The amount of memory allocated to the VM instance.",
				Optional:            true,
			},
			"price_hr": schema.Float64Attribute{
				MarkdownDescription: "The price per hour for the VM instance.",
				Computed:            true,
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

	projectId := state.ProjectID.ValueString()
	vmId := state.Id.ValueString()
	if state.ProjectID.IsNull() {
		state.ProjectID = types.StringValue(d.client.DefaultProjectID)
		projectId = d.client.DefaultProjectID
	}

	res, err := d.client.VMClient.GetVM(ctx, &vm.GetVMRequest{
		ProjectId: projectId,
		Id:        vmId,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read VM instance",
			err.Error(),
		)
		return
	}

	imageID := res.VM.PrivateImageId
	if imageID == "" {
		imageID = res.VM.PublicImageId
	}

	metadataMap, diag := types.MapValueFrom(ctx, types.StringType, res.VM.Metadata)
	if diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	state.BootDiskSizeGib = types.Int64Value(int64(res.VM.BootDiskSizeGib))
	state.CPUModel = types.StringValue(res.VM.CpuModel)
	state.DatacenterID = types.StringValue(res.VM.DatacenterId)
	state.GpuModel = types.StringValue(res.VM.GpuModel)
	state.Gpus = types.Int64Value(int64(res.VM.GpuQuantity))
	state.ImageID = types.StringValue(imageID)
	state.InternalIPAddress = types.StringValue(res.VM.InternalIpAddress)
	state.ExternalIPAddress = types.StringValue(res.VM.ExternalIpAddress)
	state.Metadata = metadataMap
	state.Memory = types.Int64Value(int64(res.VM.Memory))
	state.PriceHr = types.Float64Value(float64(res.VM.PriceHr))
	state.Vcpus = types.Int64Value(int64(res.VM.Vcpus))

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
