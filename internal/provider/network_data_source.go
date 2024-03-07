package provider

import (
	"context"
	"fmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/network"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NetworkDataSource{}

func NewNetworkDataSource() datasource.DataSource {
	return &NetworkDataSource{}
}

// NetworkDataSource defines the data source implementation.
type NetworkDataSource struct {
	client *CudoClientData
}

// NetworkDataSourceModel describes the data source data model.
type NetworkDataSourceModel struct {
	Id                types.String `tfsdk:"id"`
	DataCenterId      types.String `tfsdk:"data_center_id"`
	IPRange           types.String `tfsdk:"ip_range"`
	Gateway           types.String `tfsdk:"gateway"`
	ExternalIPAddress types.String `tfsdk:"external_ip_address"`
	InternalIPAddress types.String `tfsdk:"internal_ip_address"`
}

func (d *NetworkDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "cudo_network"
}

func (d *NetworkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Network data source",
		Description:         "Fetches a network",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"data_center_id": schema.StringAttribute{
				MarkdownDescription: "The id of the datacenter where the network is located.",
				Computed:            true,
			},
			"ip_range": schema.StringAttribute{
				MarkdownDescription: "IP Range in CIDR format e.g 192.168.0.0/24",
				Computed:            true},
			"gateway": schema.StringAttribute{
				MarkdownDescription: "IP of the gateway for the network",
				Computed:            true,
			},
			"external_ip_address": schema.StringAttribute{
				MarkdownDescription: "External IP of the network router",
				Computed:            true,
			},
			"internal_ip_address": schema.StringAttribute{
				MarkdownDescription: "Internal IP of the network router",
				Computed:            true,
			},
		},
	}
}

func (d *NetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *NetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state NetworkDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	res, err := d.client.NetworkClient.GetNetwork(ctx, &network.GetNetworkRequest{
		ProjectId: d.client.DefaultProjectID,
		Id:        state.Id.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read networks",
			err.Error(),
		)
		return
	}

	// Save data into Terraform state

	state.Id = types.StringValue(res.Network.Id)
	state.DataCenterId = types.StringValue(res.Network.DataCenterId)
	state.IPRange = types.StringValue(res.Network.IpRange)
	state.Gateway = types.StringValue(res.Network.Gateway)
	state.ExternalIPAddress = types.StringValue(res.Network.ExternalIpAddress)
	state.InternalIPAddress = types.StringValue(res.Network.InternalIpAddress)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
