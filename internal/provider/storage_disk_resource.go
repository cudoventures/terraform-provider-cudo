package provider

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/vm"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/grpc/codes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &StorageDiskResource{}
var _ resource.ResourceWithConfigure = &StorageDiskResource{}
var _ resource.ResourceWithImportState = &StorageDiskResource{}
var _ resource.ResourceWithModifyPlan = &StorageDiskResource{}

func NewStorageDiskResource() resource.Resource {
	return &StorageDiskResource{}
}

// DiskResource defines the resource implementation.
type StorageDiskResource struct {
	client *CudoClientData
}

// SecurityGroupResourceModel describes the resource data model.
type StorageDiskResourceModel struct {
	DataCenterID types.String `tfsdk:"data_center_id"`
	ID           types.String `tfsdk:"id"`
	ProjectID    types.String `tfsdk:"project_id"`
	SizeGib      types.Int64  `tfsdk:"size_gib"`
}

func (r *StorageDiskResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_storage_disk"
}

func (r *StorageDiskResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Storage disk resource",
		Attributes: map[string]schema.Attribute{
			"data_center_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The unique identifier of the datacenter where the disk is located.",
				Required:            true,
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"project_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The project the storage disk is in.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid RFC1034 resource id")},
			},
			"id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The unique identifier of the storage disk",
				Required:    true,
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"size_gib": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Size of the storage disk in GiB",
				Required:    true,
			},
		},
	}
}

func (r *StorageDiskResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *StorageDiskResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var projectID types.String

	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("project_id"), &projectID)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if projectID.IsUnknown() && r.client.DefaultProjectID != "" {
		projectID = types.StringValue(r.client.DefaultProjectID)
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("project_id"), projectID)...)
	}
}

func (r *StorageDiskResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan StorageDiskResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.VMClient.CreateStorageDisk(ctx, &vm.CreateStorageDiskRequest{
		Disk: &vm.Disk{
			DataCenterId: plan.DataCenterID.ValueString(),
			Id:           string(plan.ID.ValueString()),
			ProjectId:    plan.ProjectID.ValueString(),
			SizeGib:      int32(plan.SizeGib.ValueInt64()),
		},
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create storage disk resource",
			err.Error(),
		)
		return
	}

	if err := waitForDiskCreate(ctx, r.client.VMClient, plan.ProjectID.ValueString(), plan.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for Disk resource to be created",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func waitForDiskCreate(ctx context.Context, c vm.VMServiceClient, projectID, diskID string) error {
	refreshFunc := func() (interface{}, string, error) {
		res, err := c.GetDisk(ctx, &vm.GetDiskRequest{
			ProjectId: projectID,
			Id:        diskID,
		})
		if err != nil {
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				return res, vm.Disk_CREATING.String(), nil
			}
			return nil, vm.Disk_UNKNOWN.String(), err
		}

		return res, res.Disk.DiskState.String(), nil
	}

	stateConf := &helper.StateChangeConf{
		Pending: []string{
			vm.Disk_ATTACHED.String(),
			vm.Disk_CLONING.String(),
			vm.Disk_CREATING.String(),
			vm.Disk_DELETING.String(),
			vm.Disk_DISABLED.String(),
			vm.Disk_FAILED.String(),
			vm.Disk_UPDATING.String(),
		},
		Target: []string{
			vm.Disk_READY.String(),
			vm.Disk_UNKNOWN.String(),
		},
		Refresh:      refreshFunc,
		Timeout:      20 * time.Minute,
		Delay:        1 * time.Second,
		MinTimeout:   3 * time.Second,
		PollInterval: 5 * time.Second,
	}

	_, err := stateConf.WaitForState(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *StorageDiskResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *StorageDiskResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.ProjectID.IsNull() {
		state.ProjectID = types.StringValue(r.client.DefaultProjectID)
	}

	res, err := r.client.VMClient.GetDisk(ctx, &vm.GetDiskRequest{
		Id:        state.ID.ValueString(),
		ProjectId: state.ProjectID.ValueString(),
	})
	if err != nil {
		if helper.IsErrCode(err, codes.NotFound) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to read disk resource",
			err.Error(),
		)
		return
	}

	state.DataCenterID = types.StringValue(res.Disk.DataCenterId)
	state.ID = types.StringValue(res.Disk.Id)
	state.SizeGib = types.Int64Value(int64(res.Disk.SizeGib))

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *StorageDiskResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Unable to update storage disk",
		"Updating a storage disk is not supported",
	)
}

func (r *StorageDiskResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *StorageDiskResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectID := state.ProjectID.ValueString()
	if state.ProjectID.IsNull() {
		projectID = r.client.DefaultProjectID
	}

	_, err := r.client.VMClient.DeleteStorageDisk(ctx, &vm.DeleteStorageDiskRequest{
		ProjectId: projectID,
		Id:        state.ID.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete storage disk",
			err.Error(),
		)
		return
	}

	if err := waitForDiskDelete(ctx, r.client.VMClient, projectID, state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for Disk resource to be deleted",
			err.Error(),
		)
		return
	}
}

func waitForDiskDelete(ctx context.Context, c vm.VMServiceClient, projectID, diskID string) error {
	refreshFunc := func() (interface{}, string, error) {
		res, err := c.GetDisk(ctx, &vm.GetDiskRequest{
			Id:        diskID,
			ProjectId: projectID,
		})
		if err != nil {
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				return res, "NOT_FOUND", nil
			}
			return nil, vm.Disk_UNKNOWN.String(), err
		}

		return res, res.Disk.DiskState.String(), nil
	}

	stateConf := &helper.StateChangeConf{
		Pending: []string{
			vm.Disk_ATTACHED.String(),
			vm.Disk_CLONING.String(),
			vm.Disk_CREATING.String(),
			vm.Disk_DELETING.String(),
			vm.Disk_DISABLED.String(),
			vm.Disk_READY.String(),
			vm.Disk_UPDATING.String(),
		},
		Target: []string{
			vm.Disk_FAILED.String(),
			vm.Disk_UNKNOWN.String(),
			"NOT_FOUND",
		},
		Refresh:      refreshFunc,
		Timeout:      20 * time.Minute,
		Delay:        1 * time.Second,
		MinTimeout:   3 * time.Second,
		PollInterval: 5 * time.Second,
	}

	_, err := stateConf.WaitForState(ctx)
	if err != nil {
		return err
	}

	return nil
}

var storageDiskImportIDRegExp = regexp.MustCompile("projects/(.+)/disks/(.+)")

func (r *StorageDiskResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var projectID, ID string
	if parts := storageDiskImportIDRegExp.FindStringSubmatch(req.ID); parts != nil {
		projectID = parts[1]
		ID = parts[2]
	}

	if projectID == "" || ID == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: \"projects/<project_id>/disks/<id>\". Got: %q", req.ID),
		)
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("project_id"), projectID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), ID)...)
}
