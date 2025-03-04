// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: svc/compute/network/network_types.proto

package network

import (
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	decimal "google.golang.org/genproto/googleapis/type/decimal"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VRouterSize int32

const (
	VRouterSize_VROUTER_INSTANCE_UNKNOWN VRouterSize = 0
	VRouterSize_VROUTER_INSTANCE_SMALL   VRouterSize = 1
	VRouterSize_VROUTER_INSTANCE_MEDIUM  VRouterSize = 2
	VRouterSize_VROUTER_INSTANCE_LARGE   VRouterSize = 3
)

// Enum value maps for VRouterSize.
var (
	VRouterSize_name = map[int32]string{
		0: "VROUTER_INSTANCE_UNKNOWN",
		1: "VROUTER_INSTANCE_SMALL",
		2: "VROUTER_INSTANCE_MEDIUM",
		3: "VROUTER_INSTANCE_LARGE",
	}
	VRouterSize_value = map[string]int32{
		"VROUTER_INSTANCE_UNKNOWN": 0,
		"VROUTER_INSTANCE_SMALL":   1,
		"VROUTER_INSTANCE_MEDIUM":  2,
		"VROUTER_INSTANCE_LARGE":   3,
	}
)

func (x VRouterSize) Enum() *VRouterSize {
	p := new(VRouterSize)
	*p = x
	return p
}

func (x VRouterSize) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VRouterSize) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_compute_network_network_types_proto_enumTypes[0].Descriptor()
}

func (VRouterSize) Type() protoreflect.EnumType {
	return &file_svc_compute_network_network_types_proto_enumTypes[0]
}

func (x VRouterSize) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VRouterSize.Descriptor instead.
func (VRouterSize) EnumDescriptor() ([]byte, []int) {
	return file_svc_compute_network_network_types_proto_rawDescGZIP(), []int{0}
}

type SecurityGroup_Rule_Protocol int32

const (
	SecurityGroup_Rule_PROTOCOL_UNKNOWN SecurityGroup_Rule_Protocol = 0
	SecurityGroup_Rule_PROTOCOL_ALL     SecurityGroup_Rule_Protocol = 1
	SecurityGroup_Rule_PROTOCOL_TCP     SecurityGroup_Rule_Protocol = 2
	SecurityGroup_Rule_PROTOCOL_UDP     SecurityGroup_Rule_Protocol = 3
	SecurityGroup_Rule_PROTOCOL_ICMP    SecurityGroup_Rule_Protocol = 4
	SecurityGroup_Rule_PROTOCOL_ICMPv6  SecurityGroup_Rule_Protocol = 5
	SecurityGroup_Rule_PROTOCOL_IPSEC   SecurityGroup_Rule_Protocol = 6
)

// Enum value maps for SecurityGroup_Rule_Protocol.
var (
	SecurityGroup_Rule_Protocol_name = map[int32]string{
		0: "PROTOCOL_UNKNOWN",
		1: "PROTOCOL_ALL",
		2: "PROTOCOL_TCP",
		3: "PROTOCOL_UDP",
		4: "PROTOCOL_ICMP",
		5: "PROTOCOL_ICMPv6",
		6: "PROTOCOL_IPSEC",
	}
	SecurityGroup_Rule_Protocol_value = map[string]int32{
		"PROTOCOL_UNKNOWN": 0,
		"PROTOCOL_ALL":     1,
		"PROTOCOL_TCP":     2,
		"PROTOCOL_UDP":     3,
		"PROTOCOL_ICMP":    4,
		"PROTOCOL_ICMPv6":  5,
		"PROTOCOL_IPSEC":   6,
	}
)

func (x SecurityGroup_Rule_Protocol) Enum() *SecurityGroup_Rule_Protocol {
	p := new(SecurityGroup_Rule_Protocol)
	*p = x
	return p
}

func (x SecurityGroup_Rule_Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecurityGroup_Rule_Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_compute_network_network_types_proto_enumTypes[1].Descriptor()
}

func (SecurityGroup_Rule_Protocol) Type() protoreflect.EnumType {
	return &file_svc_compute_network_network_types_proto_enumTypes[1]
}

func (x SecurityGroup_Rule_Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecurityGroup_Rule_Protocol.Descriptor instead.
func (SecurityGroup_Rule_Protocol) EnumDescriptor() ([]byte, []int) {
	return file_svc_compute_network_network_types_proto_rawDescGZIP(), []int{0, 0, 0}
}

type SecurityGroup_Rule_RuleType int32

const (
	SecurityGroup_Rule_RULE_TYPE_UNKNOWN  SecurityGroup_Rule_RuleType = 0
	SecurityGroup_Rule_RULE_TYPE_INBOUND  SecurityGroup_Rule_RuleType = 1
	SecurityGroup_Rule_RULE_TYPE_OUTBOUND SecurityGroup_Rule_RuleType = 2
)

// Enum value maps for SecurityGroup_Rule_RuleType.
var (
	SecurityGroup_Rule_RuleType_name = map[int32]string{
		0: "RULE_TYPE_UNKNOWN",
		1: "RULE_TYPE_INBOUND",
		2: "RULE_TYPE_OUTBOUND",
	}
	SecurityGroup_Rule_RuleType_value = map[string]int32{
		"RULE_TYPE_UNKNOWN":  0,
		"RULE_TYPE_INBOUND":  1,
		"RULE_TYPE_OUTBOUND": 2,
	}
)

func (x SecurityGroup_Rule_RuleType) Enum() *SecurityGroup_Rule_RuleType {
	p := new(SecurityGroup_Rule_RuleType)
	*p = x
	return p
}

func (x SecurityGroup_Rule_RuleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecurityGroup_Rule_RuleType) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_compute_network_network_types_proto_enumTypes[2].Descriptor()
}

func (SecurityGroup_Rule_RuleType) Type() protoreflect.EnumType {
	return &file_svc_compute_network_network_types_proto_enumTypes[2]
}

func (x SecurityGroup_Rule_RuleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecurityGroup_Rule_RuleType.Descriptor instead.
func (SecurityGroup_Rule_RuleType) EnumDescriptor() ([]byte, []int) {
	return file_svc_compute_network_network_types_proto_rawDescGZIP(), []int{0, 0, 1}
}

type Network_NetworkState int32

const (
	Network_CLONING            Network_NetworkState = 0
	Network_CREATING_SNAPSHOT  Network_NetworkState = 1
	Network_DELETED            Network_NetworkState = 2
	Network_DELETING           Network_NetworkState = 3
	Network_DELETING_SNAPSHOT  Network_NetworkState = 4
	Network_FAILED             Network_NetworkState = 5
	Network_HOTPLUGGING        Network_NetworkState = 6
	Network_MIGRATING          Network_NetworkState = 7
	Network_RECREATING         Network_NetworkState = 8
	Network_REVERTING_SNAPSHOT Network_NetworkState = 9
	Network_RESIZING           Network_NetworkState = 10
	Network_RESIZING_DISK      Network_NetworkState = 11
	Network_ACTIVE             Network_NetworkState = 12
	Network_STARTING           Network_NetworkState = 13
	Network_STOPPED            Network_NetworkState = 14
	Network_STOPPING           Network_NetworkState = 15
	Network_SUSPENDED          Network_NetworkState = 16
	Network_SUSPENDING         Network_NetworkState = 17
	Network_UNKNOWN            Network_NetworkState = 18
)

// Enum value maps for Network_NetworkState.
var (
	Network_NetworkState_name = map[int32]string{
		0:  "CLONING",
		1:  "CREATING_SNAPSHOT",
		2:  "DELETED",
		3:  "DELETING",
		4:  "DELETING_SNAPSHOT",
		5:  "FAILED",
		6:  "HOTPLUGGING",
		7:  "MIGRATING",
		8:  "RECREATING",
		9:  "REVERTING_SNAPSHOT",
		10: "RESIZING",
		11: "RESIZING_DISK",
		12: "ACTIVE",
		13: "STARTING",
		14: "STOPPED",
		15: "STOPPING",
		16: "SUSPENDED",
		17: "SUSPENDING",
		18: "UNKNOWN",
	}
	Network_NetworkState_value = map[string]int32{
		"CLONING":            0,
		"CREATING_SNAPSHOT":  1,
		"DELETED":            2,
		"DELETING":           3,
		"DELETING_SNAPSHOT":  4,
		"FAILED":             5,
		"HOTPLUGGING":        6,
		"MIGRATING":          7,
		"RECREATING":         8,
		"REVERTING_SNAPSHOT": 9,
		"RESIZING":           10,
		"RESIZING_DISK":      11,
		"ACTIVE":             12,
		"STARTING":           13,
		"STOPPED":            14,
		"STOPPING":           15,
		"SUSPENDED":          16,
		"SUSPENDING":         17,
		"UNKNOWN":            18,
	}
)

func (x Network_NetworkState) Enum() *Network_NetworkState {
	p := new(Network_NetworkState)
	*p = x
	return p
}

func (x Network_NetworkState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Network_NetworkState) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_compute_network_network_types_proto_enumTypes[3].Descriptor()
}

func (Network_NetworkState) Type() protoreflect.EnumType {
	return &file_svc_compute_network_network_types_proto_enumTypes[3]
}

func (x Network_NetworkState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Network_NetworkState.Descriptor instead.
func (Network_NetworkState) EnumDescriptor() ([]byte, []int) {
	return file_svc_compute_network_network_types_proto_rawDescGZIP(), []int{1, 0}
}

type SecurityGroup struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProjectId     string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	DataCenterId  string                 `protobuf:"bytes,2,opt,name=data_center_id,json=dataCenterId,proto3" json:"data_center_id,omitempty"`
	Id            string                 `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Rules         []*SecurityGroup_Rule  `protobuf:"bytes,5,rep,name=rules,proto3" json:"rules,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SecurityGroup) Reset() {
	*x = SecurityGroup{}
	mi := &file_svc_compute_network_network_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecurityGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityGroup) ProtoMessage() {}

func (x *SecurityGroup) ProtoReflect() protoreflect.Message {
	mi := &file_svc_compute_network_network_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityGroup.ProtoReflect.Descriptor instead.
func (*SecurityGroup) Descriptor() ([]byte, []int) {
	return file_svc_compute_network_network_types_proto_rawDescGZIP(), []int{0}
}

func (x *SecurityGroup) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *SecurityGroup) GetDataCenterId() string {
	if x != nil {
		return x.DataCenterId
	}
	return ""
}

func (x *SecurityGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SecurityGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SecurityGroup) GetRules() []*SecurityGroup_Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type Network struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Id                string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DataCenterId      string                 `protobuf:"bytes,2,opt,name=data_center_id,json=dataCenterId,proto3" json:"data_center_id,omitempty"`
	IpRange           string                 `protobuf:"bytes,3,opt,name=ip_range,json=ipRange,proto3" json:"ip_range,omitempty"`
	Gateway           string                 `protobuf:"bytes,4,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Size              VRouterSize            `protobuf:"varint,5,opt,name=size,proto3,enum=org.cudo.compute.v1.VRouterSize" json:"size,omitempty"`
	PriceHr           *decimal.Decimal       `protobuf:"bytes,6,opt,name=price_hr,json=priceHr,proto3" json:"price_hr,omitempty"`
	ExternalIpAddress string                 `protobuf:"bytes,7,opt,name=external_ip_address,json=externalIpAddress,proto3" json:"external_ip_address,omitempty"`
	InternalIpAddress string                 `protobuf:"bytes,8,opt,name=internal_ip_address,json=internalIpAddress,proto3" json:"internal_ip_address,omitempty"`
	ShortState        string                 `protobuf:"bytes,11,opt,name=short_state,json=shortState,proto3" json:"short_state,omitempty"`
	State             Network_NetworkState   `protobuf:"varint,12,opt,name=state,proto3,enum=org.cudo.compute.v1.Network_NetworkState" json:"state,omitempty"`
	CreateTime        *timestamppb.Timestamp `protobuf:"bytes,31,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Network) Reset() {
	*x = Network{}
	mi := &file_svc_compute_network_network_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network) ProtoMessage() {}

func (x *Network) ProtoReflect() protoreflect.Message {
	mi := &file_svc_compute_network_network_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network.ProtoReflect.Descriptor instead.
func (*Network) Descriptor() ([]byte, []int) {
	return file_svc_compute_network_network_types_proto_rawDescGZIP(), []int{1}
}

func (x *Network) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Network) GetDataCenterId() string {
	if x != nil {
		return x.DataCenterId
	}
	return ""
}

func (x *Network) GetIpRange() string {
	if x != nil {
		return x.IpRange
	}
	return ""
}

func (x *Network) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *Network) GetSize() VRouterSize {
	if x != nil {
		return x.Size
	}
	return VRouterSize_VROUTER_INSTANCE_UNKNOWN
}

func (x *Network) GetPriceHr() *decimal.Decimal {
	if x != nil {
		return x.PriceHr
	}
	return nil
}

func (x *Network) GetExternalIpAddress() string {
	if x != nil {
		return x.ExternalIpAddress
	}
	return ""
}

func (x *Network) GetInternalIpAddress() string {
	if x != nil {
		return x.InternalIpAddress
	}
	return ""
}

func (x *Network) GetShortState() string {
	if x != nil {
		return x.ShortState
	}
	return ""
}

func (x *Network) GetState() Network_NetworkState {
	if x != nil {
		return x.State
	}
	return Network_CLONING
}

func (x *Network) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type SecurityGroup_Rule struct {
	state    protoimpl.MessageState      `protogen:"open.v1"`
	Id       string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Protocol SecurityGroup_Rule_Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=org.cudo.compute.v1.SecurityGroup_Rule_Protocol" json:"protocol,omitempty"`
	Ports    string                      `protobuf:"bytes,3,opt,name=ports,proto3" json:"ports,omitempty"`
	RuleType SecurityGroup_Rule_RuleType `protobuf:"varint,4,opt,name=rule_type,json=ruleType,proto3,enum=org.cudo.compute.v1.SecurityGroup_Rule_RuleType" json:"rule_type,omitempty"`
	// single IP or CIDR format range to apply rule to
	IpRangeCidr   string `protobuf:"bytes,5,opt,name=ip_range_cidr,json=ipRangeCidr,proto3" json:"ip_range_cidr,omitempty"`
	IcmpType      string `protobuf:"bytes,6,opt,name=icmp_type,json=icmpType,proto3" json:"icmp_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SecurityGroup_Rule) Reset() {
	*x = SecurityGroup_Rule{}
	mi := &file_svc_compute_network_network_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecurityGroup_Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityGroup_Rule) ProtoMessage() {}

func (x *SecurityGroup_Rule) ProtoReflect() protoreflect.Message {
	mi := &file_svc_compute_network_network_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityGroup_Rule.ProtoReflect.Descriptor instead.
func (*SecurityGroup_Rule) Descriptor() ([]byte, []int) {
	return file_svc_compute_network_network_types_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SecurityGroup_Rule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SecurityGroup_Rule) GetProtocol() SecurityGroup_Rule_Protocol {
	if x != nil {
		return x.Protocol
	}
	return SecurityGroup_Rule_PROTOCOL_UNKNOWN
}

func (x *SecurityGroup_Rule) GetPorts() string {
	if x != nil {
		return x.Ports
	}
	return ""
}

func (x *SecurityGroup_Rule) GetRuleType() SecurityGroup_Rule_RuleType {
	if x != nil {
		return x.RuleType
	}
	return SecurityGroup_Rule_RULE_TYPE_UNKNOWN
}

func (x *SecurityGroup_Rule) GetIpRangeCidr() string {
	if x != nil {
		return x.IpRangeCidr
	}
	return ""
}

func (x *SecurityGroup_Rule) GetIcmpType() string {
	if x != nil {
		return x.IcmpType
	}
	return ""
}

var File_svc_compute_network_network_types_proto protoreflect.FileDescriptor

var file_svc_compute_network_network_types_proto_rawDesc = string([]byte{
	0x0a, 0x27, 0x73, 0x76, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x63,
	0x75, 0x64, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x6d,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3,
	0x05, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x29, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x64, 0x61,
	0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x75, 0x64, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x1a, 0xf1, 0x03, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4c, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x63, 0x75, 0x64, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x52, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x4d, 0x0a,
	0x09, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x30, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x75, 0x64, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x69, 0x70, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x70, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x69, 0x64, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x63, 0x6d, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x63, 0x6d, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0x92, 0x01,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52,
	0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x41, 0x4c, 0x4c,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x54,
	0x43, 0x50, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c,
	0x5f, 0x55, 0x44, 0x50, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43,
	0x4f, 0x4c, 0x5f, 0x49, 0x43, 0x4d, 0x50, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x4f,
	0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x43, 0x4d, 0x50, 0x76, 0x36, 0x10, 0x05, 0x12, 0x12,
	0x0a, 0x0e, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x50, 0x53, 0x45, 0x43,
	0x10, 0x06, 0x22, 0x50, 0x0a, 0x08, 0x52, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15,
	0x0a, 0x11, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12,
	0x52, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x42, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x02, 0x22, 0xd1, 0x06, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x08, 0x69, 0x70, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x69, 0x70, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x1d, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12,
	0x39, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x63, 0x75, 0x64, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x5f, 0x68, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d,
	0x61, 0x6c, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x48, 0x72,
	0x12, 0x33, 0x0a, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x44, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x29, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x75, 0x64, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xe0, 0x41, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xba, 0x02, 0x0a, 0x0c,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x4c, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54,
	0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0f,
	0x0a, 0x0b, 0x48, 0x4f, 0x54, 0x50, 0x4c, 0x55, 0x47, 0x47, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12,
	0x0d, 0x0a, 0x09, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x0e,
	0x0a, 0x0a, 0x52, 0x45, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x08, 0x12, 0x16,
	0x0a, 0x12, 0x52, 0x45, 0x56, 0x45, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x4e, 0x41, 0x50,
	0x53, 0x48, 0x4f, 0x54, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x49, 0x5a, 0x49,
	0x4e, 0x47, 0x10, 0x0a, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x53, 0x49, 0x5a, 0x49, 0x4e, 0x47,
	0x5f, 0x44, 0x49, 0x53, 0x4b, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10,
	0x0d, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x0e, 0x12, 0x0c,
	0x0a, 0x08, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x0f, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x10, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x11, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x12, 0x2a, 0x80, 0x01, 0x0a, 0x0b, 0x56, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x56, 0x52, 0x4f, 0x55,
	0x54, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x56, 0x52, 0x4f, 0x55, 0x54, 0x45,
	0x52, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x53, 0x4d, 0x41, 0x4c, 0x4c,
	0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x56, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x49, 0x4e,
	0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12,
	0x1a, 0x0a, 0x16, 0x56, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41,
	0x4e, 0x43, 0x45, 0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x03, 0x42, 0x41, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x64, 0x6f, 0x56, 0x65,
	0x6e, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x64, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x2d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_svc_compute_network_network_types_proto_rawDescOnce sync.Once
	file_svc_compute_network_network_types_proto_rawDescData []byte
)

func file_svc_compute_network_network_types_proto_rawDescGZIP() []byte {
	file_svc_compute_network_network_types_proto_rawDescOnce.Do(func() {
		file_svc_compute_network_network_types_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_svc_compute_network_network_types_proto_rawDesc), len(file_svc_compute_network_network_types_proto_rawDesc)))
	})
	return file_svc_compute_network_network_types_proto_rawDescData
}

var file_svc_compute_network_network_types_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_svc_compute_network_network_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_svc_compute_network_network_types_proto_goTypes = []any{
	(VRouterSize)(0),                 // 0: org.cudo.compute.v1.VRouterSize
	(SecurityGroup_Rule_Protocol)(0), // 1: org.cudo.compute.v1.SecurityGroup.Rule.Protocol
	(SecurityGroup_Rule_RuleType)(0), // 2: org.cudo.compute.v1.SecurityGroup.Rule.RuleType
	(Network_NetworkState)(0),        // 3: org.cudo.compute.v1.Network.NetworkState
	(*SecurityGroup)(nil),            // 4: org.cudo.compute.v1.SecurityGroup
	(*Network)(nil),                  // 5: org.cudo.compute.v1.Network
	(*SecurityGroup_Rule)(nil),       // 6: org.cudo.compute.v1.SecurityGroup.Rule
	(*decimal.Decimal)(nil),          // 7: google.type.Decimal
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
}
var file_svc_compute_network_network_types_proto_depIdxs = []int32{
	6, // 0: org.cudo.compute.v1.SecurityGroup.rules:type_name -> org.cudo.compute.v1.SecurityGroup.Rule
	0, // 1: org.cudo.compute.v1.Network.size:type_name -> org.cudo.compute.v1.VRouterSize
	7, // 2: org.cudo.compute.v1.Network.price_hr:type_name -> google.type.Decimal
	3, // 3: org.cudo.compute.v1.Network.state:type_name -> org.cudo.compute.v1.Network.NetworkState
	8, // 4: org.cudo.compute.v1.Network.create_time:type_name -> google.protobuf.Timestamp
	1, // 5: org.cudo.compute.v1.SecurityGroup.Rule.protocol:type_name -> org.cudo.compute.v1.SecurityGroup.Rule.Protocol
	2, // 6: org.cudo.compute.v1.SecurityGroup.Rule.rule_type:type_name -> org.cudo.compute.v1.SecurityGroup.Rule.RuleType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_svc_compute_network_network_types_proto_init() }
func file_svc_compute_network_network_types_proto_init() {
	if File_svc_compute_network_network_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_svc_compute_network_network_types_proto_rawDesc), len(file_svc_compute_network_network_types_proto_rawDesc)),
			NumEnums:      4,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_svc_compute_network_network_types_proto_goTypes,
		DependencyIndexes: file_svc_compute_network_network_types_proto_depIdxs,
		EnumInfos:         file_svc_compute_network_network_types_proto_enumTypes,
		MessageInfos:      file_svc_compute_network_network_types_proto_msgTypes,
	}.Build()
	File_svc_compute_network_network_types_proto = out.File
	file_svc_compute_network_network_types_proto_goTypes = nil
	file_svc_compute_network_network_types_proto_depIdxs = nil
}
