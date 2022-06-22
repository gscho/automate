// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: external/iam/v2/common/rules.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RuleType int32

const (
	RuleType_RULE_TYPE_UNSET RuleType = 0
	RuleType_NODE            RuleType = 1
	RuleType_EVENT           RuleType = 2
)

// Enum value maps for RuleType.
var (
	RuleType_name = map[int32]string{
		0: "RULE_TYPE_UNSET",
		1: "NODE",
		2: "EVENT",
	}
	RuleType_value = map[string]int32{
		"RULE_TYPE_UNSET": 0,
		"NODE":            1,
		"EVENT":           2,
	}
)

func (x RuleType) Enum() *RuleType {
	p := new(RuleType)
	*p = x
	return p
}

func (x RuleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RuleType) Descriptor() protoreflect.EnumDescriptor {
	return file_external_iam_v2_common_rules_proto_enumTypes[0].Descriptor()
}

func (RuleType) Type() protoreflect.EnumType {
	return &file_external_iam_v2_common_rules_proto_enumTypes[0]
}

func (x RuleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RuleType.Descriptor instead.
func (RuleType) EnumDescriptor() ([]byte, []int) {
	return file_external_iam_v2_common_rules_proto_rawDescGZIP(), []int{0}
}

type ConditionAttribute int32

const (
	ConditionAttribute_CONDITION_ATTRIBUTE_UNSET ConditionAttribute = 0
	ConditionAttribute_CHEF_SERVER               ConditionAttribute = 1
	ConditionAttribute_CHEF_ORGANIZATION         ConditionAttribute = 2
	ConditionAttribute_ENVIRONMENT               ConditionAttribute = 3
	ConditionAttribute_CHEF_ROLE                 ConditionAttribute = 4
	ConditionAttribute_CHEF_TAG                  ConditionAttribute = 5
	ConditionAttribute_CHEF_POLICY_GROUP         ConditionAttribute = 6
	ConditionAttribute_CHEF_POLICY_NAME          ConditionAttribute = 7
)

// Enum value maps for ConditionAttribute.
var (
	ConditionAttribute_name = map[int32]string{
		0: "CONDITION_ATTRIBUTE_UNSET",
		1: "CHEF_SERVER",
		2: "CHEF_ORGANIZATION",
		3: "ENVIRONMENT",
		4: "CHEF_ROLE",
		5: "CHEF_TAG",
		6: "CHEF_POLICY_GROUP",
		7: "CHEF_POLICY_NAME",
	}
	ConditionAttribute_value = map[string]int32{
		"CONDITION_ATTRIBUTE_UNSET": 0,
		"CHEF_SERVER":               1,
		"CHEF_ORGANIZATION":         2,
		"ENVIRONMENT":               3,
		"CHEF_ROLE":                 4,
		"CHEF_TAG":                  5,
		"CHEF_POLICY_GROUP":         6,
		"CHEF_POLICY_NAME":          7,
	}
)

func (x ConditionAttribute) Enum() *ConditionAttribute {
	p := new(ConditionAttribute)
	*p = x
	return p
}

func (x ConditionAttribute) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConditionAttribute) Descriptor() protoreflect.EnumDescriptor {
	return file_external_iam_v2_common_rules_proto_enumTypes[1].Descriptor()
}

func (ConditionAttribute) Type() protoreflect.EnumType {
	return &file_external_iam_v2_common_rules_proto_enumTypes[1]
}

func (x ConditionAttribute) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConditionAttribute.Descriptor instead.
func (ConditionAttribute) EnumDescriptor() ([]byte, []int) {
	return file_external_iam_v2_common_rules_proto_rawDescGZIP(), []int{1}
}

type ConditionOperator int32

const (
	ConditionOperator_CONDITION_OPERATOR_UNSET ConditionOperator = 0
	ConditionOperator_MEMBER_OF                ConditionOperator = 1
	ConditionOperator_EQUALS                   ConditionOperator = 2
)

// Enum value maps for ConditionOperator.
var (
	ConditionOperator_name = map[int32]string{
		0: "CONDITION_OPERATOR_UNSET",
		1: "MEMBER_OF",
		2: "EQUALS",
	}
	ConditionOperator_value = map[string]int32{
		"CONDITION_OPERATOR_UNSET": 0,
		"MEMBER_OF":                1,
		"EQUALS":                   2,
	}
)

func (x ConditionOperator) Enum() *ConditionOperator {
	p := new(ConditionOperator)
	*p = x
	return p
}

func (x ConditionOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConditionOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_external_iam_v2_common_rules_proto_enumTypes[2].Descriptor()
}

func (ConditionOperator) Type() protoreflect.EnumType {
	return &file_external_iam_v2_common_rules_proto_enumTypes[2]
}

func (x ConditionOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConditionOperator.Descriptor instead.
func (ConditionOperator) EnumDescriptor() ([]byte, []int) {
	return file_external_iam_v2_common_rules_proto_rawDescGZIP(), []int{2}
}

type RuleStatus int32

const (
	RuleStatus_RULE_STATUS_UNSET RuleStatus = 0
	RuleStatus_STAGED            RuleStatus = 1
	RuleStatus_APPLIED           RuleStatus = 2
)

// Enum value maps for RuleStatus.
var (
	RuleStatus_name = map[int32]string{
		0: "RULE_STATUS_UNSET",
		1: "STAGED",
		2: "APPLIED",
	}
	RuleStatus_value = map[string]int32{
		"RULE_STATUS_UNSET": 0,
		"STAGED":            1,
		"APPLIED":           2,
	}
)

func (x RuleStatus) Enum() *RuleStatus {
	p := new(RuleStatus)
	*p = x
	return p
}

func (x RuleStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RuleStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_external_iam_v2_common_rules_proto_enumTypes[3].Descriptor()
}

func (RuleStatus) Type() protoreflect.EnumType {
	return &file_external_iam_v2_common_rules_proto_enumTypes[3]
}

func (x RuleStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RuleStatus.Descriptor instead.
func (RuleStatus) EnumDescriptor() ([]byte, []int) {
	return file_external_iam_v2_common_rules_proto_rawDescGZIP(), []int{3}
}

type ProjectRulesStatus int32

const (
	ProjectRulesStatus_PROJECT_RULES_STATUS_UNSET ProjectRulesStatus = 0
	ProjectRulesStatus_RULES_APPLIED              ProjectRulesStatus = 1
	ProjectRulesStatus_EDITS_PENDING              ProjectRulesStatus = 2
	ProjectRulesStatus_NO_RULES                   ProjectRulesStatus = 3
)

// Enum value maps for ProjectRulesStatus.
var (
	ProjectRulesStatus_name = map[int32]string{
		0: "PROJECT_RULES_STATUS_UNSET",
		1: "RULES_APPLIED",
		2: "EDITS_PENDING",
		3: "NO_RULES",
	}
	ProjectRulesStatus_value = map[string]int32{
		"PROJECT_RULES_STATUS_UNSET": 0,
		"RULES_APPLIED":              1,
		"EDITS_PENDING":              2,
		"NO_RULES":                   3,
	}
)

func (x ProjectRulesStatus) Enum() *ProjectRulesStatus {
	p := new(ProjectRulesStatus)
	*p = x
	return p
}

func (x ProjectRulesStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectRulesStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_external_iam_v2_common_rules_proto_enumTypes[4].Descriptor()
}

func (ProjectRulesStatus) Type() protoreflect.EnumType {
	return &file_external_iam_v2_common_rules_proto_enumTypes[4]
}

func (x ProjectRulesStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectRulesStatus.Descriptor instead.
func (ProjectRulesStatus) EnumDescriptor() ([]byte, []int) {
	return file_external_iam_v2_common_rules_proto_rawDescGZIP(), []int{4}
}

type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique ID of the project this rule belongs to. Cannot be changed.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Name for the project rule.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Whether the rule applies to ingested `NODE` or `EVENT resources.
	// Cannot be changed.
	Type RuleType `protobuf:"varint,4,opt,name=type,proto3,enum=chef.automate.api.iam.v2.RuleType" json:"type,omitempty"`
	// Conditions that ingested resources must match to belong to the project.
	// Will contain one or more.
	Conditions []*Condition `protobuf:"bytes,5,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// Whether the rule is `STAGED` (not in effect) or `APPLIED` (in effect).
	Status RuleStatus `protobuf:"varint,6,opt,name=status,proto3,enum=chef.automate.api.iam.v2.RuleStatus" json:"status,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_common_rules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_common_rules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_common_rules_proto_rawDescGZIP(), []int{0}
}

func (x *Rule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Rule) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Rule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rule) GetType() RuleType {
	if x != nil {
		return x.Type
	}
	return RuleType_RULE_TYPE_UNSET
}

func (x *Rule) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *Rule) GetStatus() RuleStatus {
	if x != nil {
		return x.Status
	}
	return RuleStatus_RULE_STATUS_UNSET
}

type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents a property of an ingested resource. Depends on the rule type.
	Attribute ConditionAttribute `protobuf:"varint,1,opt,name=attribute,proto3,enum=chef.automate.api.iam.v2.ConditionAttribute" json:"attribute,omitempty"`
	// The value(s) of the attribute that an ingested resource must match.
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	// Whether the attribute matches a single value (`EQUALS`) or
	// matches at least one of a set of values (`MEMBER_OF`).
	Operator ConditionOperator `protobuf:"varint,3,opt,name=operator,proto3,enum=chef.automate.api.iam.v2.ConditionOperator" json:"operator,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_iam_v2_common_rules_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_common_rules_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_common_rules_proto_rawDescGZIP(), []int{1}
}

func (x *Condition) GetAttribute() ConditionAttribute {
	if x != nil {
		return x.Attribute
	}
	return ConditionAttribute_CONDITION_ATTRIBUTE_UNSET
}

func (x *Condition) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Condition) GetOperator() ConditionOperator {
	if x != nil {
		return x.Operator
	}
	return ConditionOperator_CONDITION_OPERATOR_UNSET
}

var File_external_iam_v2_common_rules_proto protoreflect.FileDescriptor

var file_external_iam_v2_common_rules_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76,
	0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x22, 0x84,
	0x02, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x32, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xb8, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2a, 0x34, 0x0a, 0x08, 0x52, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f,
	0x52, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x2a, 0xb6, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x19, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49,
	0x42, 0x55, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x43, 0x48, 0x45, 0x46, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x01, 0x12, 0x15, 0x0a,
	0x11, 0x43, 0x48, 0x45, 0x46, 0x5f, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x56, 0x49, 0x52, 0x4f, 0x4e, 0x4d,
	0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x48, 0x45, 0x46, 0x5f, 0x52, 0x4f,
	0x4c, 0x45, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x48, 0x45, 0x46, 0x5f, 0x54, 0x41, 0x47,
	0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x45, 0x46, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43,
	0x59, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x48, 0x45,
	0x46, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x07, 0x2a,
	0x4c, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4f, 0x46, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x02, 0x2a, 0x3c, 0x0a,
	0x0a, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x52,
	0x55, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x47, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x68, 0x0a, 0x12, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x55, 0x4c,
	0x45, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x55, 0x4c, 0x45, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x44, 0x49, 0x54, 0x53, 0x5f, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x52, 0x55,
	0x4c, 0x45, 0x53, 0x10, 0x03, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_iam_v2_common_rules_proto_rawDescOnce sync.Once
	file_external_iam_v2_common_rules_proto_rawDescData = file_external_iam_v2_common_rules_proto_rawDesc
)

func file_external_iam_v2_common_rules_proto_rawDescGZIP() []byte {
	file_external_iam_v2_common_rules_proto_rawDescOnce.Do(func() {
		file_external_iam_v2_common_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_iam_v2_common_rules_proto_rawDescData)
	})
	return file_external_iam_v2_common_rules_proto_rawDescData
}

var file_external_iam_v2_common_rules_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_external_iam_v2_common_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_external_iam_v2_common_rules_proto_goTypes = []interface{}{
	(RuleType)(0),           // 0: chef.automate.api.iam.v2.RuleType
	(ConditionAttribute)(0), // 1: chef.automate.api.iam.v2.ConditionAttribute
	(ConditionOperator)(0),  // 2: chef.automate.api.iam.v2.ConditionOperator
	(RuleStatus)(0),         // 3: chef.automate.api.iam.v2.RuleStatus
	(ProjectRulesStatus)(0), // 4: chef.automate.api.iam.v2.ProjectRulesStatus
	(*Rule)(nil),            // 5: chef.automate.api.iam.v2.Rule
	(*Condition)(nil),       // 6: chef.automate.api.iam.v2.Condition
}
var file_external_iam_v2_common_rules_proto_depIdxs = []int32{
	0, // 0: chef.automate.api.iam.v2.Rule.type:type_name -> chef.automate.api.iam.v2.RuleType
	6, // 1: chef.automate.api.iam.v2.Rule.conditions:type_name -> chef.automate.api.iam.v2.Condition
	3, // 2: chef.automate.api.iam.v2.Rule.status:type_name -> chef.automate.api.iam.v2.RuleStatus
	1, // 3: chef.automate.api.iam.v2.Condition.attribute:type_name -> chef.automate.api.iam.v2.ConditionAttribute
	2, // 4: chef.automate.api.iam.v2.Condition.operator:type_name -> chef.automate.api.iam.v2.ConditionOperator
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_external_iam_v2_common_rules_proto_init() }
func file_external_iam_v2_common_rules_proto_init() {
	if File_external_iam_v2_common_rules_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_iam_v2_common_rules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_iam_v2_common_rules_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_iam_v2_common_rules_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_iam_v2_common_rules_proto_goTypes,
		DependencyIndexes: file_external_iam_v2_common_rules_proto_depIdxs,
		EnumInfos:         file_external_iam_v2_common_rules_proto_enumTypes,
		MessageInfos:      file_external_iam_v2_common_rules_proto_msgTypes,
	}.Build()
	File_external_iam_v2_common_rules_proto = out.File
	file_external_iam_v2_common_rules_proto_rawDesc = nil
	file_external_iam_v2_common_rules_proto_goTypes = nil
	file_external_iam_v2_common_rules_proto_depIdxs = nil
}
