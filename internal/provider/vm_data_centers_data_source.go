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
var _ datasource.DataSource = &VMDataCentersDataSource{}

func NewVMDataCentersDataSource() datasource.DataSource {
	return &VMDataCentersDataSource{}
}

// VMDataCentersDataSource defines the data source implementation.
type VMDataCentersDataSource struct {
	client *CudoClientData
}

type VMDataCenterDataSourceModel struct {
	ID       types.String `tfsdk:"id"`
	RegionID types.String `tfsdk:"region_id"`
}

// VMDataCentersDataSourceModelModel describes the data source data model.
type VMDataCentersDataSourceModel struct {
	DataCenters []VMDataCenterDataSourceModel `tfsdk:"data_centers"`
	ID          types.String                  `tfsdk:"id"`
}

func (d *VMDataCentersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "cudo_vm_data_centers"
}

func (d *VMDataCentersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "VM data centers data source",
		Description:         "Fetches the list of data centers that VMs can be deployed into",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier attribute.",
				Computed:    true,
			},
			"data_centers": schema.ListNestedAttribute{
				Description: "List of data centers.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Data center identifier",
							Computed:            true,
						},
						"region_id": schema.StringAttribute{
							MarkdownDescription: "Region identifier",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *VMDataCentersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *VMDataCentersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state VMDataCentersDataSourceModel

	res, err := d.client.VMClient.ListVMDataCenters(ctx, &vm.ListVMDataCentersRequest{})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read VM data centers",
			err.Error(),
		)
		return
	}

	for _, dc := range res.DataCenters {
		dataCenter := VMDataCenterDataSourceModel{
			ID:       types.StringValue(dc.Id),
			RegionID: types.StringValue(dc.RegionId),
		}

		state.DataCenters = append(state.DataCenters, dataCenter)
	}

	state.ID = types.StringValue("vm_data_centers")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
