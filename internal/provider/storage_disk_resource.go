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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"google.golang.org/grpc/codes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &StorageDiskResource{}
var _ resource.ResourceWithImportState = &StorageDiskResource{}

func NewStorageDiskResource() resource.Resource {
	return &StorageDiskResource{}
}

// DiskResource defines the resource implementation.
type StorageDiskResource struct {
	client *CudoClientData
}

// SecurityGroupResourceModel describes the resource data model.
type StorageDiskResourceModel struct {
	ProjectID    types.String `tfsdk:"project_id"`
	DataCenterID types.String `tfsdk:"data_center_id"`
	Id           types.String `tfsdk:"id"`
	SizeGib      types.Int64  `tfsdk:"size_gib"`
}

func (r *StorageDiskResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "cudo_storage_disk"
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

func (r *StorageDiskResource) waitForDiskDelete(ctx context.Context, diskID, projectID string) error {
	refreshFunc := func() (interface{}, string, error) {
		res, err := r.client.VMClient.GetDisk(ctx, &vm.GetDiskRequest{
			Id:        diskID,
			ProjectId: projectID,
		})
		if err != nil {
			// if not found resource has been deleted
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				// tflog.Debug(ctx, fmt.Sprintf("VM %s in project %s not found: ", vmID, projectID))
				return res, "done", nil
			}
			return nil, "", err
		}

		// tflog.Trace(ctx, fmt.Sprintf("pending VM %s in project %s state: %s", vmID, projectID, res.Payload.VM.ShortState))
		return res, res.Disk.DiskState.String(), nil
	}

	tflog.Debug(ctx, fmt.Sprintf("waiting for Disk %s in project %s ", diskID, projectID))

	pendingStates := []string{
		vm.Disk_DISK_STATE_INIT.String(),
		vm.Disk_DISK_STATE_READY.String(),
		vm.Disk_DISK_STATE_USED.String(),
		vm.Disk_DISK_STATE_DISABLED.String(),
		vm.Disk_DISK_STATE_LOCKED.String(),
		vm.Disk_DISK_STATE_ERROR.String(),
		vm.Disk_DISK_STATE_CLONE.String(),
		vm.Disk_DISK_STATE_DELETE.String(),
	}
	stateConf := &helper.StateChangeConf{
		Pending:      pendingStates,
		Target:       []string{"done"},
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

func (r *StorageDiskResource) waitForDiskCreate(ctx context.Context, diskID, projectID string) error {
	refreshFunc := func() (interface{}, string, error) {
		res, err := r.client.VMClient.GetDisk(ctx, &vm.GetDiskRequest{
			ProjectId: projectID,
			Id:        diskID,
		})
		if err != nil {
			// if not found assume resource is initializing
			if ok := helper.IsErrCode(err, codes.NotFound); ok {
				// tflog.Debug(ctx, fmt.Sprintf("VM %s in project %s not found: ", vmID, projectID))
				return res, vm.Disk_DISK_STATE_INIT.String(), nil
			}
			return nil, "", err
		}

		// tflog.Trace(ctx, fmt.Sprintf("pending VM %s in project %s state: %s", vmID, projectID, res.Payload.VM.ShortState))
		return res, res.Disk.DiskState.String(), nil
	}

	tflog.Debug(ctx, fmt.Sprintf("waiting for Disk %s in project %s ", diskID, projectID))

	pendingStates := []string{
		vm.Disk_DISK_STATE_INIT.String(),
		vm.Disk_DISK_STATE_USED.String(),
		vm.Disk_DISK_STATE_DISABLED.String(),
		vm.Disk_DISK_STATE_LOCKED.String(),
		vm.Disk_DISK_STATE_ERROR.String(),
		vm.Disk_DISK_STATE_CLONE.String(),
		vm.Disk_DISK_STATE_DELETE.String(),
	}

	stateConf := &helper.StateChangeConf{
		Pending:      pendingStates,
		Target:       []string{vm.Disk_DISK_STATE_READY.String()},
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

func (r *StorageDiskResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state StorageDiskResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectID := r.client.DefaultProjectID
	if !state.ProjectID.IsNull() {
		projectID = state.ProjectID.ValueString()
	}

	_, err := r.client.VMClient.CreateStorageDisk(ctx, &vm.CreateStorageDiskRequest{
		DataCenterId: state.DataCenterID.ValueString(),
		ProjectId:    projectID,
		Disk: &vm.Disk{
			Id:      string(state.Id.ValueString()),
			SizeGib: int32(state.SizeGib.ValueInt64()),
		},
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create storage disk resource",
			err.Error(),
		)
		return
	}

	if err := r.waitForDiskCreate(ctx, state.Id.ValueString(), projectID); err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for Disk resource to be created",
			err.Error(),
		)
		return
	}

	// state.DataCenterID = types.StringValue()
	// state.Id = types.StringPointerValue(params.Body.Disk.ID)
	// state.SizeGib = types.Int64Value(int64(*params.Body.Disk.SizeGib))

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *StorageDiskResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *StorageDiskResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectId := r.client.DefaultProjectID
	if !state.ProjectID.IsNull() {
		projectId = state.ProjectID.ValueString()
	}

	res, err := r.client.VMClient.GetDisk(ctx, &vm.GetDiskRequest{
		Id:        state.Id.ValueString(),
		ProjectId: projectId,
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
	state.Id = types.StringValue(res.Disk.Id)
	state.SizeGib = types.Int64Value(int64(res.Disk.SizeGib))

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *StorageDiskResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan StorageDiskResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Error getting storage disk plan",
			"Error getting storage disk plan",
		)
		return
	}

	// Read Terraform state data into the model
	var state StorageDiskResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *StorageDiskResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *StorageDiskResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	projectId := r.client.DefaultProjectID
	if !state.ProjectID.IsNull() {
		projectId = state.ProjectID.ValueString()
	}

	_, err := r.client.VMClient.DeleteStorageDisk(ctx, &vm.DeleteStorageDiskRequest{
		ProjectId: projectId,
		Id:        state.Id.ValueString(),
	})

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete storage disk",
			err.Error(),
		)
		return
	}

	if err := r.waitForDiskDelete(ctx, state.Id.ValueString(), projectId); err != nil {
		resp.Diagnostics.AddError(
			"Unable to wait for Disk resource to be deleted",
			err.Error(),
		)
		return
	}

	tflog.Trace(ctx, "deleted storage disk")
}

func (r *StorageDiskResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
