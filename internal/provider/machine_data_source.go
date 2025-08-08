package provider

import (
	"context"
	"fmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/baremetal"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MachineDataSource{}

func NewMachineDataSource() datasource.DataSource {
	return &MachineDataSource{}
}

// MachineDataSource defines the data source implementation.
type MachineDataSource struct {
	client *CudoClientData
}

type MachineDataSourceModel struct {
	DataCenterID types.String `tfsdk:"data_center_id"`
	ID           types.String `tfsdk:"id"`
	MachineType  types.String `tfsdk:"machine_type"`
	ProjectID    types.String `tfsdk:"project_id"`
	OS           types.String `tfsdk:"os"`

	ExternalIPAddresses types.List   `tfsdk:"external_ip_addresses"`
	State               types.String `tfsdk:"state"`
	PowerState          types.String `tfsdk:"power_state"`
}

func (d *MachineDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_machine"
}

func (d *MachineDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Machine data source",
		Description:         "Gets a machine",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the machine within the project.",
				Required:            true,
			},
			"data_center_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the datacenter where the machine is located.",
				Computed:            true,
			},
			"machine_type": schema.StringAttribute{
				MarkdownDescription: "Machine type of machine.",
				Optional:            true,
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the project the machine is in.",
				Optional:            true,
			},
			"os": schema.StringAttribute{
				MarkdownDescription: "Operating system deployed to machine.",
				Optional:            true,
			},
			"external_ip_addresses": schema.ListAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "The external IP addresses allocated to the machine.",
				Computed:            true,
			},
			"state": schema.StringAttribute{
				MarkdownDescription: "State of the machine.",
				Computed:            true,
			},
			"power_state": schema.StringAttribute{
				MarkdownDescription: "The power state of the machine.",
				Computed:            true,
			},
		},
	}
}

func (d *MachineDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *MachineDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state MachineDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	projectID := state.ProjectID.ValueString()
	if state.ProjectID.IsNull() {
		state.ProjectID = types.StringValue(d.client.DefaultProjectID)
		projectID = d.client.DefaultProjectID
	}

	machine, err := d.client.BareMetalClient.GetMachine(ctx, &baremetal.GetMachineRequest{
		ProjectId: projectID,
		Id:        state.ID.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read machine",
			err.Error(),
		)
		return
	}

	state.DataCenterID = types.StringValue(machine.DataCenterId)
	state.MachineType = types.StringValue(machine.MachineTypeId)
	ipAddresses := make([]attr.Value, len(machine.ExternalIpAddresses))
	for i, ipAddress := range machine.ExternalIpAddresses {
		ipAddresses[i] = types.StringValue(ipAddress)
	}
	var diags diag.Diagnostics
	state.ExternalIPAddresses, diags = types.ListValue(types.StringType, ipAddresses)
	if diags != nil {
		resp.Diagnostics.Append(diags...)
	}
	state.OS = types.StringValue(machine.Os)
	state.PowerState = types.StringValue(machine.PowerState.String())
	state.State = types.StringValue(machine.State.String())

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
