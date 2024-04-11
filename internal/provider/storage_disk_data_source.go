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
var _ datasource.DataSource = &StorageDiskDataSource{}

func NewStorageDiskDataSource() datasource.DataSource {
	return &StorageDiskDataSource{}
}

// SecurityGroupsDataSource defines the data source implementation.
type StorageDiskDataSource struct {
	client *CudoClientData
}

// SecurityGroupDataSourceModel describes the resource data model.
type StorageDiskDataSourceModel struct {
	ID           types.String `tfsdk:"id"`
	ProjectID    types.String `tfsdk:"project_id"`
	DataCenterID types.String `tfsdk:"data_center_id"`
	SizeGib      types.Int64  `tfsdk:"size_gib"`
}

func (d *StorageDiskDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "cudo_storage_disk"
}

func (d *StorageDiskDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Disk data source",
		Description:         "Gets a Disk",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Storage disk ID.",
				Required:    true,
			},
			"project_id": schema.StringAttribute{
				Description: "The unique identifier of the project the disk is in.",
				Optional:    true,
			},
			"data_center_id": schema.StringAttribute{
				Description: "The unique identifier of the datacenter where the disk is located.",
				Computed:    true,
			},
			"size_gib": schema.Int64Attribute{
				Description: "Size of the storage disk in GiB",
				Computed:    true,
			},
		},
	}
}

func (d *StorageDiskDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*CudoClientData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *CudoClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *StorageDiskDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state StorageDiskDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	projectId := d.client.DefaultProjectID
	if !state.ProjectID.IsNull() {
		projectId = state.ProjectID.ValueString()
	}

	res, err := d.client.VMClient.GetDisk(ctx, &vm.GetDiskRequest{
		ProjectId: projectId,
		Id:        state.ID.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read storage disks",
			err.Error(),
		)
		return
	}

	state.DataCenterID = types.StringValue(res.Disk.DataCenterId)
	state.SizeGib = types.Int64Value(int64(res.Disk.SizeGib))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
