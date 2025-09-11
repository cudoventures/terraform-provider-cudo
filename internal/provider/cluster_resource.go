package provider

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute"
	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/baremetal"
	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/sshkey"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/grpc/codes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ClusterResource{}
var _ resource.ResourceWithConfigure = &ClusterResource{}
var _ resource.ResourceWithImportState = &ClusterResource{}
var _ resource.ResourceWithModifyPlan = &ClusterResource{}

func NewClusterResource() resource.Resource {
	return &ClusterResource{}
}

func (*ClusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (*ClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Cluster resource",
		Attributes: map[string]schema.Attribute{
			"commitment_months": schema.Int32Attribute{
				MarkdownDescription: "The minimum length of time to commit to the cluster. It cannot be deleted before the commitment end date.",
				Optional:            true,
				Validators:          []validator.Int32{int32validator.OneOf(1, 3, 6, 12, 24, 36)},
			},
			"data_center_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The id of the data center where the cluster is located.",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid RFC1034 resource id")},
			},
			"id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "ID for cluster within project.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid RFC1034 resource id")},
			},
			"machine_count": schema.Int32Attribute{
				MarkdownDescription: "Number of machines in cluster",
				Required:            true,
			},
			"machine_type": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Machine type of cluster. See console for valid options",
				Required:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid RFC1034 resource id")},
			},
			"project_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The project the cluster is in.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid RFC1034 resource id")},
			},
			"ssh_key_source": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Which SSH keys to add to the cluster machines: project (default), user or custom",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("project", "user", "custom")},
			},
			"ssh_keys": schema.ListAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				ElementType:         types.StringType,
				MarkdownDescription: "List of SSH keys to add to the cluster machines, ssh_key_source must be set to custom",
				Optional:            true,
			},
			"start_script": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "A script to run when cluster machine boots",
				Optional:            true,
			},
			"state": schema.StringAttribute{
				MarkdownDescription: "State of the cluster.",
				Computed:            true,
			},
		},
	}
}

// ClusterResource defines the resource implementation.
type ClusterResource struct {
	client *CudoClientData
}

// ClusterResourceModel describes the resource data model.
type ClusterResourceModel struct {
	CommitmentMonths types.Int32    `tfsdk:"commitment_months"`
	DataCenterID     types.String   `tfsdk:"data_center_id"`
	ID               types.String   `tfsdk:"id"`
	MachineCount     types.Int32    `tfsdk:"machine_count"`
	MachineType      types.String   `tfsdk:"machine_type"`
	ProjectID        types.String   `tfsdk:"project_id"`
	SSHKeys          []types.String `tfsdk:"ssh_keys"`
	SSHKeySource     types.String   `tfsdk:"ssh_key_source"`
	StartScript      types.String   `tfsdk:"start_script"`

	State types.String `tfsdk:"state"`
}

type ClusterMachineResourceModel struct {
	ID                types.String `tfsdk:"id"`
	ExternalIPAddress types.String `tfsdk:"external_ip_address"`
	InternalIPAddress types.String `tfsdk:"internal_ip_address"`
	PowerState        types.String `tfsdk:"power_state"`
	State             types.String `tfsdk:"state"`
}

func (r *ClusterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ClusterResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var projectID types.String
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("project_id"), &projectID)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if projectID.IsUnknown() && r.client.DefaultProjectID != "" {
		projectID = types.StringValue(r.client.DefaultProjectID)
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("project_id"), projectID)...)
	}
}

func (r *ClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ClusterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
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

	params := &baremetal.CreateClusterRequest{
		Cluster: &baremetal.Cluster{
			CommitmentTerm: commitmentTerm,
			DataCenterId:   plan.DataCenterID.ValueString(),
			Id:             plan.ID.ValueString(),
			MachineTypeId:  plan.MachineType.ValueString(),
			ProjectId:      plan.ProjectID.ValueString(),
			MachineCount:   plan.MachineCount.ValueInt32(),
			SshKeySource:   sshKeySource,
			CustomSshKeys:  customKeys,
			StartScript:    plan.StartScript.ValueString(),
		},
	}

	_, err := r.client.BareMetalClient.CreateCluster(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create cluster",
			err.Error(),
		)
		return
	}

	cluster, err := waitForClusterDeploying(ctx, r.client.BareMetalClient, plan.ProjectID.ValueString(), plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for cluster to deploy",
			err.Error(),
		)
		return
	}

	// if the cluster is created and returned update the state.
	if cluster != nil {
		appendClusterState(cluster, &plan)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func waitForClusterDeploying(ctx context.Context, c baremetal.BareMetalServiceClient, projectID string, ID string) (*baremetal.Cluster, error) {
	refreshFunc := func() (interface{}, string, error) {
		params := &baremetal.GetClusterRequest{
			Id:        ID,
			ProjectId: projectID,
		}
		cluster, err := c.GetCluster(ctx, params)
		if err != nil {
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				return cluster, "NOT_FOUND", nil
			}
			return nil, baremetal.ClusterMachine_STATE_UNSPECIFIED.String(), err
		}

		return cluster, cluster.State.String(), nil
	}

	stateConf := &helper.StateChangeConf{
		Pending: []string{
			baremetal.ClusterMachine_CREATING.String(),
		},
		Target: []string{
			baremetal.ClusterMachine_ACTIVE.String(),
			baremetal.ClusterMachine_DELETING.String(),
			baremetal.ClusterMachine_FAILED.String(),
			baremetal.ClusterMachine_STATE_UNSPECIFIED.String(),
			"NOT_FOUND",
		},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return nil, fmt.Errorf("error waiting for cluster %s in project %s: %w", ID, projectID, err)
	} else if cluster, ok := res.(*baremetal.Cluster); ok {
		return cluster, nil
	} else {
		return nil, fmt.Errorf("error waiting for cluster: %v", res)
	}
}

func (r *ClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ClusterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := &baremetal.GetClusterRequest{
		ProjectId: state.ProjectID.ValueString(),
		Id:        state.ID.ValueString(),
	}

	cluster, err := r.client.BareMetalClient.GetCluster(ctx, params)
	if err != nil {
		if ok := helper.IsErrCode(err, codes.NotFound); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to read cluster",
			err.Error(),
		)
		return
	}

	appendClusterState(cluster, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Unable to update cluster",
		"Updating a cluster is not supported",
	)
}

func (r *ClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ClusterResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	projectID := state.ProjectID.ValueString()
	if state.ProjectID.IsNull() {
		projectID = r.client.DefaultProjectID
	}

	ID := state.ID.ValueString()

	_, err := r.client.BareMetalClient.DeleteCluster(ctx, &baremetal.DeleteClusterRequest{
		ProjectId: projectID,
		Id:        ID,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete cluster",
			err.Error(),
		)
		return
	}

	err = waitForClusterDeleting(ctx, r.client.BareMetalClient, projectID, ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for cluster to be deleted",
			err.Error(),
		)
		return
	}
}

func waitForClusterDeleting(ctx context.Context, c baremetal.BareMetalServiceClient, projectID string, ID string) error {
	refreshFunc := func() (interface{}, string, error) {
		params := &baremetal.GetClusterRequest{
			Id:        ID,
			ProjectId: projectID,
		}
		cluster, err := c.GetCluster(ctx, params)
		if err != nil {
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				return cluster, "NOT_FOUND", nil
			}
			return nil, baremetal.ClusterMachine_STATE_UNSPECIFIED.String(), err
		}

		return cluster, cluster.State.String(), nil
	}

	stateConf := &helper.StateChangeConf{
		Pending: []string{
			baremetal.ClusterMachine_DELETING.String(),
		},
		Target: []string{
			baremetal.ClusterMachine_ACTIVE.String(),
			baremetal.ClusterMachine_CREATING.String(),
			baremetal.ClusterMachine_FAILED.String(),
			baremetal.ClusterMachine_STATE_UNSPECIFIED.String(),
			"NOT_FOUND",
		},
		Refresh:    refreshFunc,
		Timeout:    2 * time.Hour,
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if res, err := stateConf.WaitForState(ctx); err != nil {
		return fmt.Errorf("error waiting for cluster %s in project %s: %w", ID, projectID, err)
	} else if _, ok := res.(*baremetal.Cluster); ok {
		return nil
	} else {
		return fmt.Errorf("error waiting for cluster: %v", res)
	}
}

func appendClusterState(cluster *baremetal.Cluster, state *ClusterResourceModel) {
	switch cluster.CommitmentTerm {
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

	state.DataCenterID = types.StringValue(cluster.DataCenterId)
	state.ID = types.StringValue(cluster.Id)
	state.MachineCount = types.Int32Value(cluster.MachineCount)
	state.MachineType = types.StringValue(cluster.MachineTypeId)
	state.ProjectID = types.StringValue(cluster.ProjectId)
	state.State = types.StringValue(cluster.State.String())
}

var clusterImportIDRegExp = regexp.MustCompile("projects/(.+)/clusters/(.+)")

func (r *ClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var projectID, ID string
	if parts := clusterImportIDRegExp.FindStringSubmatch(req.ID); parts != nil {
		projectID = parts[1]
		ID = parts[2]
	}

	if projectID == "" || ID == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: \"projects/<project_id>/clusters/<id>\". Got: %q", req.ID),
		)
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("project_id"), projectID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), ID)...)
}
