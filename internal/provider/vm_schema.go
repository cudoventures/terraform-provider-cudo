package provider

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *VMResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "VM resource",
		Attributes: map[string]schema.Attribute{
			"boot_disk": schema.SingleNestedAttribute{
				MarkdownDescription: "Specification for boot disk",
				Attributes: map[string]schema.Attribute{
					"size_gib": schema.Int64Attribute{
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
						Computed:            true,
						Optional:            true,
						MarkdownDescription: "Size of boot disk in Gib",
					},
					"image_id": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						MarkdownDescription: "ID of OS image on boot disk",
						Required:            true,
						Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
					},
				},
				Required: true,
			},
			"cpu_model": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The model of the CPU.",
				Optional:            true,
				Computed:            true,
			},
			"data_center_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The id of the datacenter where the VM instance is located.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"external_ip_address": schema.StringAttribute{
				MarkdownDescription: "The external IP address of the VM instance.",
				Computed:            true,
			},
			"gpu_model": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The model of the GPU.",
				Optional:            true,
				Computed:            true,
			},
			"gpus": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Number of GPUs",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
			},
			"id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "ID for VM within project",
				Required:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id e.g. my-vm")},
			},
			"internal_ip_address": schema.StringAttribute{
				MarkdownDescription: "The internal IP address of the VM instance.",
				Computed:            true,
			},
			"machine_type": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "VM machine type, from machine type data source",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"max_price_hr": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The maximum price per hour for the VM instance.",
				Optional:            true,
			},
			"memory_gib": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Amount of VM memory in GiB",
				Optional:            true,
			},
			"networks": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Network adapters for private networks",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_id": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							MarkdownDescription: "ID of private network to attach the NIC to",
							Required:            true,
						},
						"assign_public_ip": schema.BoolAttribute{
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.RequiresReplace(),
							},
							MarkdownDescription: "Assign a public IP to the NIC",
							Optional:            true,
						},
						"external_ip_address": schema.StringAttribute{
							MarkdownDescription: "The external IP address of the NIC.",
							Computed:            true,
						},
						"internal_ip_address": schema.StringAttribute{
							MarkdownDescription: "The internal IP address of the NIC.",
							Computed:            true,
						},
						"security_group_ids": schema.SetAttribute{
							ElementType:         types.StringType,
							Optional:            true,
							MarkdownDescription: "Security groups to assign to the NIC",
						},
					},
				},
			},
			"password": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Root password for linux, or Admin password for windows",
				Optional:            true,
				Sensitive:           true,
				Validators:          []validator.String{stringvalidator.LengthBetween(6, 64)},
			},
			"price_hr": schema.StringAttribute{
				MarkdownDescription: "The current price per hour for the VM instance.",
				Computed:            true,
			},
			"project_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The project the VM instance is in.",
				Optional:            true,
			},
			"renewable_energy": schema.BoolAttribute{
				MarkdownDescription: "Whether the VM instance is powered by renewable energy",
				Computed:            true,
			},
			"security_group_ids": schema.SetAttribute{
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Security groups to assign to the VM when using public networking",
			},
			"ssh_key_source": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Which SSH keys to add to the VM: project (default), user or custom",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("project", "user", "custom")},
			},
			"ssh_keys": schema.ListAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				ElementType:         types.StringType,
				MarkdownDescription: "List of SSH keys to add to the VM, ssh_key_source must be set to custom",
				Optional:            true,
			},
			"start_script": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "A script to run when VM boots",
				Optional:            true,
			},
			"vcpus": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Number of VCPUs",
				Optional:            true,
				Validators:          []validator.Int64{int64validator.AtMost(100)},
			},
		},
	}
}
