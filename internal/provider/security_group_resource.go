package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/network"
	"github.com/CudoVentures/terraform-provider-cudo/internal/compute/securitygroup"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"google.golang.org/grpc/codes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SecurityGroupResource{}
var _ resource.ResourceWithConfigure = &SecurityGroupResource{}
var _ resource.ResourceWithImportState = &SecurityGroupResource{}
var _ resource.ResourceWithModifyPlan = &SecurityGroupResource{}

func NewSecurityGroupResource() resource.Resource {
	return &SecurityGroupResource{}
}

// SecurityGroupResource defines the resource implementation.
type SecurityGroupResource struct {
	client *CudoClientData
}

type RuleModel struct {
	IcmpType    types.String `tfsdk:"icmp_type"`
	Id          types.String `tfsdk:"id"`
	IpRangeCidr types.String `tfsdk:"ip_range"`
	Ports       types.String `tfsdk:"ports"`
	Protocol    types.String `tfsdk:"protocol"`
	RuleType    types.String `tfsdk:"rule_type"`
}

// SecurityGroupResourceModel describes the resource data model.
type SecurityGroupResourceModel struct {
	DataCenterID types.String `tfsdk:"data_center_id"`
	Description  types.String `tfsdk:"description"`
	ID           types.String `tfsdk:"id"`
	ProjectID    types.String `tfsdk:"project_id"`
	Rules        []RuleModel  `tfsdk:"rules"`
}

func (r *SecurityGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_group"
}

func (r *SecurityGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Security group resource",
		Attributes: map[string]schema.Attribute{
			"data_center_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the datacenter where the network is located.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the security group",
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.RegexMatches(securityGroupDescriptionRegex,
					"must be a valid security group description up to 255 characters, commas, periods, & spaces")},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Security Group ID",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"project_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The project the security group is in.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid RFC1034 resource id")},
			},
			"rules": schema.ListNestedAttribute{
				Description: "List of security group rules",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"icmp_type": schema.StringAttribute{
							Description: "Specific ICMP type of the rule. If a type has multiple codes, it includes all the codes within. This can only be used with ICMP",
							Optional:    true,
							Validators: []validator.String{stringvalidator.RegexMatches(
								icmpTypesRegex, "must be a valid icmp type i.e. 0,3,4,5,8,9,10,11,12,13,14,17,18")},
						},
						"id": schema.StringAttribute{
							Description: "The unique identifier of the rule",
							Computed:    true,
						},
						"ip_range": schema.StringAttribute{
							Description: "A single IP address or CIDR format range to apply rule to",
							Optional:    true,
						},
						"ports": schema.StringAttribute{
							Description: "A comma separated list of ports (80,443,8080) or a single port range (1024:2048)",
							Optional:    true,
							Validators:  []validator.String{portListValidator{}},
						},
						"protocol": schema.StringAttribute{
							Description: "Protocol for rule, use one of: all, tcp, udp, icmp, icmpv6, ipsec",
							Computed:    true,
							Optional:    true,
							Validators:  []validator.String{stringvalidator.OneOf("all", "tcp", "udp", "icmp", "icmpv6", "ipsec")},
						},
						"rule_type": schema.StringAttribute{
							Description: "Type for rule either 'inbound' or 'outbound'",
							Required:    true,
							Validators:  []validator.String{stringvalidator.OneOf("inbound", "outbound")},
						},
					},
				},
			},
		},
	}
}

func (r *SecurityGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SecurityGroupResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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

func getNullableString(value string) basetypes.StringValue {
	var result = basetypes.StringValue{}
	if value != "" {
		result = types.StringValue(value)
	} else {
		result = types.StringNull()
	}

	return result
}

func getRulePlan(rules []*securitygroup.SecurityGroup_Rule) []RuleModel {
	var ruleModels []RuleModel
	for _, rule := range rules {
		protocol := ""
		switch rule.Protocol {
		case securitygroup.SecurityGroup_Rule_PROTOCOL_ALL:
			protocol = "all"
		case securitygroup.SecurityGroup_Rule_PROTOCOL_ICMP:
			protocol = "icmp"
		case securitygroup.SecurityGroup_Rule_PROTOCOL_ICMPv6:
			protocol = "icmpv6"
		case securitygroup.SecurityGroup_Rule_PROTOCOL_IPSEC:
			protocol = "ipsec"
		case securitygroup.SecurityGroup_Rule_PROTOCOL_TCP:
			protocol = "tcp"
		case securitygroup.SecurityGroup_Rule_PROTOCOL_UDP:
			protocol = "udp"
		}

		ruleType := ""
		switch rule.RuleType {
		case securitygroup.SecurityGroup_Rule_RULE_TYPE_INBOUND:
			ruleType = "inbound"
		case securitygroup.SecurityGroup_Rule_RULE_TYPE_OUTBOUND:
			ruleType = "outbound"
		}

		ruleModel := RuleModel{
			IcmpType:    getNullableString(rule.IcmpType),
			Id:          types.StringValue(rule.Id),
			IpRangeCidr: getNullableString(rule.IpRangeCidr),
			Ports:       getNullableString(rule.Ports),
			Protocol:    getNullableString(protocol),
			RuleType:    getNullableString(ruleType),
		}

		ruleModels = append(ruleModels, ruleModel)
	}

	return ruleModels
}

func getRuleParams(stateRules []RuleModel) []*securitygroup.SecurityGroup_Rule {
	var rules []*securitygroup.SecurityGroup_Rule

	for _, r := range stateRules {
		var protocol securitygroup.SecurityGroup_Rule_Protocol
		switch r.Protocol.ValueString() {
		case "tcp":
			protocol = securitygroup.SecurityGroup_Rule_PROTOCOL_TCP
		case "udp":
			protocol = securitygroup.SecurityGroup_Rule_PROTOCOL_UDP
		case "icmp":
			protocol = securitygroup.SecurityGroup_Rule_PROTOCOL_ICMP
		case "icmpv6":
			protocol = securitygroup.SecurityGroup_Rule_PROTOCOL_ICMPv6
		case "ipsec":
			protocol = securitygroup.SecurityGroup_Rule_PROTOCOL_IPSEC
		default:
			protocol = securitygroup.SecurityGroup_Rule_PROTOCOL_ALL
		}

		var ruleType securitygroup.SecurityGroup_Rule_RuleType
		switch r.RuleType.ValueString() {
		case "inbound":
			ruleType = securitygroup.SecurityGroup_Rule_RULE_TYPE_INBOUND
		case "outbound":
			ruleType = securitygroup.SecurityGroup_Rule_RULE_TYPE_OUTBOUND
		}

		rule := securitygroup.SecurityGroup_Rule{
			IcmpType:    r.IcmpType.ValueString(),
			Id:          r.Id.ValueString(),
			IpRangeCidr: r.IpRangeCidr.ValueString(),
			Ports:       r.Ports.ValueString(),
			Protocol:    protocol,
			RuleType:    ruleType,
		}
		rules = append(rules, &rule)
	}

	return rules
}

func (r *SecurityGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan *SecurityGroupResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.NetworkClient.CreateSecurityGroup(ctx, &network.CreateSecurityGroupRequest{
		SecurityGroup: &securitygroup.SecurityGroup{
			ProjectId:    plan.ProjectID.ValueString(),
			DataCenterId: plan.DataCenterID.ValueString(),
			Description:  plan.Description.ValueString(),
			Id:           plan.ID.ValueString(),
			Rules:        getRuleParams(plan.Rules),
		},
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create security group",
			err.Error(),
		)
		return
	}

	plan.Rules = getRulePlan(res.Rules)

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SecurityGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SecurityGroupResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.ProjectID.IsNull() {
		state.ProjectID = types.StringValue(r.client.DefaultProjectID)
	}

	res, err := r.client.NetworkClient.GetSecurityGroup(ctx, &network.GetSecurityGroupRequest{
		Id:        state.ID.ValueString(),
		ProjectId: state.ProjectID.ValueString(),
	})

	if err != nil {
		if ok := helper.IsErrCode(err, codes.NotFound); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to read security group",
			err.Error(),
		)
		return
	}

	state.ID = types.StringValue(res.Id)
	state.Description = types.StringValue(res.Description)
	state.DataCenterID = types.StringValue(res.DataCenterId)
	state.Rules = getRulePlan(res.Rules)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SecurityGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan *SecurityGroupResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.NetworkClient.UpdateSecurityGroup(ctx, &network.UpdateSecurityGroupRequest{
		SecurityGroup: &securitygroup.SecurityGroup{
			DataCenterId: plan.DataCenterID.ValueString(),
			Description:  plan.Description.ValueString(),
			Id:           plan.ID.ValueString(),
			ProjectId:    plan.ProjectID.ValueString(),
			Rules:        getRuleParams(plan.Rules),
		},
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update security group resource",
			err.Error(),
		)
		return
	}

	plan.Rules = getRulePlan(res.Rules)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SecurityGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *SecurityGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	_, err := r.client.NetworkClient.DeleteSecurityGroup(ctx, &network.DeleteSecurityGroupRequest{
		ProjectId: r.client.DefaultProjectID,
		Id:        state.ID.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete security group resource",
			err.Error(),
		)
		return
	}
}

var securityGroupImportIDRegExp = regexp.MustCompile("projects/(.+)/security-groups/(.+)")

func (r *SecurityGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var projectID, ID string
	if parts := securityGroupImportIDRegExp.FindStringSubmatch(req.ID); parts != nil {
		projectID = parts[1]
		ID = parts[2]
	}

	if projectID == "" || ID == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: \"projects/<project_id>/security-groups/<id>\". Got: %q", req.ID),
		)
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("project_id"), projectID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), ID)...)
}

var securityGroupDescriptionRegex = regexp.MustCompile(`^[A-Za-z0-9,'.\s]{1,255}$`)
var icmpTypesRegex = regexp.MustCompile("^(0|3|4|5|8|9|10|11|12|13|14|17|18)$")

type portListValidator struct {
}

func (p portListValidator) Description(ctx context.Context) string {
	return "must be a comma separated list of ports 80,443,8080 or a single port range 1024:2048"
}

func (p portListValidator) MarkdownDescription(ctx context.Context) string {
	return "must be a comma separated list of ports 80,443,8080 or a single port range 1024:2048"
}

func (p portListValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	portString := req.ConfigValue.ValueString()

	_, ok := validatePorts(portString)

	if !ok {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid port range",
			"String length must be a single range i.e. 1024:2048 or a comma separated list i.e. 80,443,8080",
		)
	}

}

func validatePorts(ports string) ([]uint16, bool) {
	var invalid bool
	portRanges := strings.Split(ports, ",")
	ranges := make([]uint16, 0, len(portRanges))
	for _, portRange := range portRanges {
		portRange = strings.Trim(portRange, " ")
		startStr, endStr, found := strings.Cut(portRange, ":")
		if !found {
			endStr = startStr
		}
		start, err := strconv.ParseUint(startStr, 10, 16)
		if err != nil {
			invalid = true
			continue
		}
		end, err := strconv.ParseUint(endStr, 10, 16)
		if err != nil {
			invalid = true
			continue
		}
		if end < start {
			invalid = true
			continue
		}
		ranges = append(ranges, uint16(start), uint16(end))
	}
	return ranges, !invalid
}
