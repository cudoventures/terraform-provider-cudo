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
var _ datasource.DataSource = &SecurityGroupDataSource{}

func NewSecurityGroupDataSource() datasource.DataSource {
	return &SecurityGroupDataSource{}
}

// SecurityGroupsDataSource defines the data source implementation.
type SecurityGroupDataSource struct {
	client *CudoClientData
}

type RuleDataSourceModel struct {
	IcmpType    types.String `tfsdk:"icmp_type"`
	Id          types.String `tfsdk:"id"`
	IpRangeCidr types.String `tfsdk:"ip_range"`
	Ports       types.String `tfsdk:"ports"`
	Protocol    types.String `tfsdk:"protocol"`
	RuleType    types.String `tfsdk:"rule_type"`
}

// SecurityGroupDataSourceModel describes the resource data model.
type SecurityGroupDataSourceModel struct {
	DataCenterID types.String `tfsdk:"data_center_id"`
	Description  types.String `tfsdk:"description"`
	ID           types.String `tfsdk:"id"`
	ProjectID    types.String `tfsdk:"project_id"`
	Rules        []RuleModel  `tfsdk:"rules"`
}

func (d *SecurityGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_group"
}

func (d *SecurityGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Security groups data source",
		Description:         "Fetches the list of security groups",
		Attributes: map[string]schema.Attribute{
			"data_center_id": schema.StringAttribute{
				MarkdownDescription: "Datacenter ID to request security groups from",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Security group description",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Description: "Security Group ID.",
				Required:    true,
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the project the security group is in.",
				Optional:            true,
			},
			"rules": schema.ListNestedAttribute{
				Description: "List of rules in security group",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"icmp_type": schema.StringAttribute{
							MarkdownDescription: "ICMP type",
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: "Rule ID",
							Computed:            true,
						},
						"ip_range": schema.StringAttribute{
							MarkdownDescription: "IP range",
							Computed:            true,
						},
						"ports": schema.StringAttribute{
							MarkdownDescription: "Image size in GiB",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Image size in GiB",
							Computed:            true,
						},
						"rule_type": schema.StringAttribute{
							MarkdownDescription: "Image size in GiB",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SecurityGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *SecurityGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SecurityGroupDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectID := state.ProjectID.ValueString()
	if state.ProjectID.IsNull() {
		state.ProjectID = types.StringValue(d.client.DefaultProjectID)
		projectID = d.client.DefaultProjectID
	}

	sg, err := d.client.NetworkClient.GetSecurityGroup(ctx, &network.GetSecurityGroupRequest{
		ProjectId: projectID,
		Id:        state.ID.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read security groups",
			err.Error(),
		)
		return
	}

	state.ID = types.StringValue(sg.Id)
	state.DataCenterID = types.StringValue(sg.DataCenterId)
	state.Description = types.StringValue(sg.Description)
	state.Rules = getRulePlan(sg.Rules)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
