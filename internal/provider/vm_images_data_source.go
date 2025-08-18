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
var _ datasource.DataSource = &VMImagesDataSource{}

func NewVMImagesDataSource() datasource.DataSource {
	return &VMImagesDataSource{}
}

// VMImagesDataSource defines the data source implementation.
type VMImagesDataSource struct {
	client *CudoClientData
}

type vmImagesModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	SizeGiB     types.String `tfsdk:"size_gib"`
}

// VMImagesDataSourceModel describes the data source data model.
type VMImagesDataSourceModel struct {
	VmImages []vmImagesModel `tfsdk:"images"`
	ID       types.String    `tfsdk:"id"`
}

func (d *VMImagesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vm_images"
}

func (d *VMImagesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Images data source",
		Description:         "Fetches the list of images",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
			"images": schema.ListNestedAttribute{
				Description: "List of images.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Image identifier",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Image name",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Image description",
							Computed:            true,
						},
						"size_gib": schema.StringAttribute{
							MarkdownDescription: "Image size in GiB",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *VMImagesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *VMImagesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state VMImagesDataSourceModel

	res, err := d.client.VMClient.ListPublicVMImages(ctx, &vm.ListPublicVMImagesRequest{})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read images",
			err.Error(),
		)
		return
	}

	for _, image := range res.Images {
		imageState := vmImagesModel{
			Id:          types.StringValue(image.Id),
			Name:        types.StringValue(image.Name),
			Description: types.StringValue(image.Description),
			SizeGiB:     types.StringValue(fmt.Sprintf("%d", image.SizeGib)),
		}

		state.VmImages = append(state.VmImages, imageState)
	}

	state.ID = types.StringValue("placeholder")

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
