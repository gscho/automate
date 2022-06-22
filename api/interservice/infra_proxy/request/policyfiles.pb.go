// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: interservice/infra_proxy/request/policyfiles.proto

package request

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type Policyfiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
}

func (x *Policyfiles) Reset() {
	*x = Policyfiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policyfiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policyfiles) ProtoMessage() {}

func (x *Policyfiles) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policyfiles.ProtoReflect.Descriptor instead.
func (*Policyfiles) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_policyfiles_proto_rawDescGZIP(), []int{0}
}

func (x *Policyfiles) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Policyfiles) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type Policyfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Policyfile name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Policyfile revision ID.
	RevisionId string `protobuf:"bytes,4,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty" toml:"revision_id,omitempty" mapstructure:"revision_id,omitempty"`
}

func (x *Policyfile) Reset() {
	*x = Policyfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policyfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policyfile) ProtoMessage() {}

func (x *Policyfile) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policyfile.ProtoReflect.Descriptor instead.
func (*Policyfile) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_policyfiles_proto_rawDescGZIP(), []int{1}
}

func (x *Policyfile) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Policyfile) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Policyfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Policyfile) GetRevisionId() string {
	if x != nil {
		return x.RevisionId
	}
	return ""
}

type DeletePolicyfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Policyfile name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
}

func (x *DeletePolicyfile) Reset() {
	*x = DeletePolicyfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicyfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyfile) ProtoMessage() {}

func (x *DeletePolicyfile) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyfile.ProtoReflect.Descriptor instead.
func (*DeletePolicyfile) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_policyfiles_proto_rawDescGZIP(), []int{2}
}

func (x *DeletePolicyfile) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *DeletePolicyfile) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *DeletePolicyfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PolicyfileRevisions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Policyfile name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
}

func (x *PolicyfileRevisions) Reset() {
	*x = PolicyfileRevisions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyfileRevisions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyfileRevisions) ProtoMessage() {}

func (x *PolicyfileRevisions) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyfileRevisions.ProtoReflect.Descriptor instead.
func (*PolicyfileRevisions) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_policyfiles_proto_rawDescGZIP(), []int{3}
}

func (x *PolicyfileRevisions) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *PolicyfileRevisions) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *PolicyfileRevisions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Policygroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Policygroup name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
}

func (x *Policygroup) Reset() {
	*x = Policygroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policygroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policygroup) ProtoMessage() {}

func (x *Policygroup) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policygroup.ProtoReflect.Descriptor instead.
func (*Policygroup) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_policyfiles_proto_rawDescGZIP(), []int{4}
}

func (x *Policygroup) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Policygroup) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Policygroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_interservice_infra_proxy_request_policyfiles_proto protoreflect.FileDescriptor

var file_interservice_infra_proxy_request_policyfiles_proto_rawDesc = []byte{
	0x0a, 0x32, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0b,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6f,
	0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x75, 0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x13, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x55, 0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_interservice_infra_proxy_request_policyfiles_proto_rawDescOnce sync.Once
	file_interservice_infra_proxy_request_policyfiles_proto_rawDescData = file_interservice_infra_proxy_request_policyfiles_proto_rawDesc
)

func file_interservice_infra_proxy_request_policyfiles_proto_rawDescGZIP() []byte {
	file_interservice_infra_proxy_request_policyfiles_proto_rawDescOnce.Do(func() {
		file_interservice_infra_proxy_request_policyfiles_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_infra_proxy_request_policyfiles_proto_rawDescData)
	})
	return file_interservice_infra_proxy_request_policyfiles_proto_rawDescData
}

var file_interservice_infra_proxy_request_policyfiles_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_interservice_infra_proxy_request_policyfiles_proto_goTypes = []interface{}{
	(*Policyfiles)(nil),         // 0: chef.automate.domain.infra_proxy.request.Policyfiles
	(*Policyfile)(nil),          // 1: chef.automate.domain.infra_proxy.request.Policyfile
	(*DeletePolicyfile)(nil),    // 2: chef.automate.domain.infra_proxy.request.DeletePolicyfile
	(*PolicyfileRevisions)(nil), // 3: chef.automate.domain.infra_proxy.request.PolicyfileRevisions
	(*Policygroup)(nil),         // 4: chef.automate.domain.infra_proxy.request.Policygroup
}
var file_interservice_infra_proxy_request_policyfiles_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interservice_infra_proxy_request_policyfiles_proto_init() }
func file_interservice_infra_proxy_request_policyfiles_proto_init() {
	if File_interservice_infra_proxy_request_policyfiles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policyfiles); i {
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
		file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policyfile); i {
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
		file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicyfile); i {
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
		file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyfileRevisions); i {
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
		file_interservice_infra_proxy_request_policyfiles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policygroup); i {
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
			RawDescriptor: file_interservice_infra_proxy_request_policyfiles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_infra_proxy_request_policyfiles_proto_goTypes,
		DependencyIndexes: file_interservice_infra_proxy_request_policyfiles_proto_depIdxs,
		MessageInfos:      file_interservice_infra_proxy_request_policyfiles_proto_msgTypes,
	}.Build()
	File_interservice_infra_proxy_request_policyfiles_proto = out.File
	file_interservice_infra_proxy_request_policyfiles_proto_rawDesc = nil
	file_interservice_infra_proxy_request_policyfiles_proto_goTypes = nil
	file_interservice_infra_proxy_request_policyfiles_proto_depIdxs = nil
}
