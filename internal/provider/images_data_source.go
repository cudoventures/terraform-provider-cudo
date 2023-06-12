package provider

import (
	"context"
	"fmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client/virtual_machines"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ImagesDataSource{}

func NewImagesDataSource() datasource.DataSource {
	return &ImagesDataSource{}
}

// ImagesDataSource defines the data source implementation.
type ImagesDataSource struct {
	client *CudoClientData
}

type imagesModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	SizeGiB     types.String `tfsdk:"size_gib"`
}

// ImagesDataSourceModel describes the data source data model.
type ImagesDataSourceModel struct {
	Images []imagesModel `tfsdk:"images"`
	ID     types.String  `tfsdk:"id"`
}

func (d *ImagesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "cudo_images"
}

func (d *ImagesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

func (d *ImagesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ImagesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ImagesDataSourceModel

	res, err := d.client.Client.VirtualMachines.ListPublicVMImages(virtual_machines.NewListPublicVMImagesParams())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read images",
			err.Error(),
		)
		return
	}

	//tflog.Debug(ctx, fmt.Sprintf("length of images %d", len(res.Payload.Images)))

	for _, image := range res.Payload.Images {
		imageState := imagesModel{
			Id:          types.StringValue(*image.ID),
			Name:        types.StringValue(*image.Name),
			Description: types.StringValue(*image.Description),
			SizeGiB:     types.StringValue(fmt.Sprintf("%d", *image.SizeGib)),
		}

		state.Images = append(state.Images, imageState)
		//tflog.Debug(ctx, image.Name)

	}

	state.ID = types.StringValue("placeholder")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
