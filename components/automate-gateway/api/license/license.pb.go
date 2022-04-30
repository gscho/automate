// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: automate-gateway/api/license/license.proto

package license

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApplyLicenseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	License string `protobuf:"bytes,1,opt,name=license,proto3" json:"license,omitempty"`
}

func (x *ApplyLicenseReq) Reset() {
	*x = ApplyLicenseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_automate_gateway_api_license_license_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyLicenseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyLicenseReq) ProtoMessage() {}

func (x *ApplyLicenseReq) ProtoReflect() protoreflect.Message {
	mi := &file_automate_gateway_api_license_license_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyLicenseReq.ProtoReflect.Descriptor instead.
func (*ApplyLicenseReq) Descriptor() ([]byte, []int) {
	return file_automate_gateway_api_license_license_proto_rawDescGZIP(), []int{0}
}

func (x *ApplyLicenseReq) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

type ApplyLicenseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *GetStatusResp `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ApplyLicenseResp) Reset() {
	*x = ApplyLicenseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_automate_gateway_api_license_license_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyLicenseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyLicenseResp) ProtoMessage() {}

func (x *ApplyLicenseResp) ProtoReflect() protoreflect.Message {
	mi := &file_automate_gateway_api_license_license_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyLicenseResp.ProtoReflect.Descriptor instead.
func (*ApplyLicenseResp) Descriptor() ([]byte, []int) {
	return file_automate_gateway_api_license_license_proto_rawDescGZIP(), []int{1}
}

func (x *ApplyLicenseResp) GetStatus() *GetStatusResp {
	if x != nil {
		return x.Status
	}
	return nil
}

type GetStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStatusReq) Reset() {
	*x = GetStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_automate_gateway_api_license_license_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusReq) ProtoMessage() {}

func (x *GetStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_automate_gateway_api_license_license_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusReq.ProtoReflect.Descriptor instead.
func (*GetStatusReq) Descriptor() ([]byte, []int) {
	return file_automate_gateway_api_license_license_proto_rawDescGZIP(), []int{2}
}

type GetStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LicenseId      string                   `protobuf:"bytes,1,opt,name=license_id,json=licenseId,proto3" json:"license_id,omitempty"`
	ConfiguredAt   *timestamppb.Timestamp   `protobuf:"bytes,2,opt,name=configured_at,json=configuredAt,proto3" json:"configured_at,omitempty"`
	LicensedPeriod *GetStatusResp_DateRange `protobuf:"bytes,3,opt,name=licensed_period,json=licensedPeriod,proto3" json:"licensed_period,omitempty"`
	CustomerName   string                   `protobuf:"bytes,4,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
}

func (x *GetStatusResp) Reset() {
	*x = GetStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_automate_gateway_api_license_license_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusResp) ProtoMessage() {}

func (x *GetStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_automate_gateway_api_license_license_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusResp.ProtoReflect.Descriptor instead.
func (*GetStatusResp) Descriptor() ([]byte, []int) {
	return file_automate_gateway_api_license_license_proto_rawDescGZIP(), []int{3}
}

func (x *GetStatusResp) GetLicenseId() string {
	if x != nil {
		return x.LicenseId
	}
	return ""
}

func (x *GetStatusResp) GetConfiguredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ConfiguredAt
	}
	return nil
}

func (x *GetStatusResp) GetLicensedPeriod() *GetStatusResp_DateRange {
	if x != nil {
		return x.LicensedPeriod
	}
	return nil
}

func (x *GetStatusResp) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

type RequestLicenseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	GdprAgree bool   `protobuf:"varint,4,opt,name=gdpr_agree,json=gdprAgree,proto3" json:"gdpr_agree,omitempty"`
}

func (x *RequestLicenseReq) Reset() {
	*x = RequestLicenseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_automate_gateway_api_license_license_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLicenseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLicenseReq) ProtoMessage() {}

func (x *RequestLicenseReq) ProtoReflect() protoreflect.Message {
	mi := &file_automate_gateway_api_license_license_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLicenseReq.ProtoReflect.Descriptor instead.
func (*RequestLicenseReq) Descriptor() ([]byte, []int) {
	return file_automate_gateway_api_license_license_proto_rawDescGZIP(), []int{4}
}

func (x *RequestLicenseReq) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *RequestLicenseReq) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *RequestLicenseReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RequestLicenseReq) GetGdprAgree() bool {
	if x != nil {
		return x.GdprAgree
	}
	return false
}

type RequestLicenseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	License string         `protobuf:"bytes,1,opt,name=license,proto3" json:"license,omitempty"`
	Status  *GetStatusResp `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RequestLicenseResp) Reset() {
	*x = RequestLicenseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_automate_gateway_api_license_license_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLicenseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLicenseResp) ProtoMessage() {}

func (x *RequestLicenseResp) ProtoReflect() protoreflect.Message {
	mi := &file_automate_gateway_api_license_license_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLicenseResp.ProtoReflect.Descriptor instead.
func (*RequestLicenseResp) Descriptor() ([]byte, []int) {
	return file_automate_gateway_api_license_license_proto_rawDescGZIP(), []int{5}
}

// Deprecated: Do not use.
func (x *RequestLicenseResp) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

func (x *RequestLicenseResp) GetStatus() *GetStatusResp {
	if x != nil {
		return x.Status
	}
	return nil
}

type GetStatusResp_DateRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetStatusResp_DateRange) Reset() {
	*x = GetStatusResp_DateRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_automate_gateway_api_license_license_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusResp_DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusResp_DateRange) ProtoMessage() {}

func (x *GetStatusResp_DateRange) ProtoReflect() protoreflect.Message {
	mi := &file_automate_gateway_api_license_license_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusResp_DateRange.ProtoReflect.Descriptor instead.
func (*GetStatusResp_DateRange) Descriptor() ([]byte, []int) {
	return file_automate_gateway_api_license_license_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetStatusResp_DateRange) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *GetStatusResp_DateRange) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

var File_automate_gateway_api_license_license_proto protoreflect.FileDescriptor

var file_automate_gateway_api_license_license_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x22,
	0x54, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x22, 0xde, 0x02, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x5b, 0x0a, 0x0f, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x0e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x64, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x6b, 0x0a, 0x09, 0x44, 0x61, 0x74,
	0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x64, 0x70, 0x72, 0x5f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x67, 0x64, 0x70, 0x72, 0x41, 0x67, 0x72, 0x65, 0x65, 0x22, 0x74, 0x0a,
	0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0xbf, 0x04, 0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12,
	0xbd, 0x01, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x54, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x3a, 0x01,
	0x2a, 0x8a, 0xb5, 0x18, 0x10, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x8a, 0xb5, 0x18, 0x16, 0x12, 0x14, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x3a, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x3a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x12,
	0xa9, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x49, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x30, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x8a, 0xb5, 0x18, 0x0f, 0x0a, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x8a, 0xb5, 0x18, 0x14, 0x12, 0x12, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x3a, 0x67, 0x65, 0x74, 0x12, 0xc7, 0x01, 0x0a, 0x0e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2d, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x58, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x03, 0x3a, 0x01, 0x2a, 0x8a, 0xb5, 0x18, 0x10, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x3a, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x8a, 0xb5, 0x18, 0x18, 0x12, 0x16, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x3a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_automate_gateway_api_license_license_proto_rawDescOnce sync.Once
	file_automate_gateway_api_license_license_proto_rawDescData = file_automate_gateway_api_license_license_proto_rawDesc
)

func file_automate_gateway_api_license_license_proto_rawDescGZIP() []byte {
	file_automate_gateway_api_license_license_proto_rawDescOnce.Do(func() {
		file_automate_gateway_api_license_license_proto_rawDescData = protoimpl.X.CompressGZIP(file_automate_gateway_api_license_license_proto_rawDescData)
	})
	return file_automate_gateway_api_license_license_proto_rawDescData
}

var file_automate_gateway_api_license_license_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_automate_gateway_api_license_license_proto_goTypes = []interface{}{
	(*ApplyLicenseReq)(nil),         // 0: chef.automate.api.license.ApplyLicenseReq
	(*ApplyLicenseResp)(nil),        // 1: chef.automate.api.license.ApplyLicenseResp
	(*GetStatusReq)(nil),            // 2: chef.automate.api.license.GetStatusReq
	(*GetStatusResp)(nil),           // 3: chef.automate.api.license.GetStatusResp
	(*RequestLicenseReq)(nil),       // 4: chef.automate.api.license.RequestLicenseReq
	(*RequestLicenseResp)(nil),      // 5: chef.automate.api.license.RequestLicenseResp
	(*GetStatusResp_DateRange)(nil), // 6: chef.automate.api.license.GetStatusResp.DateRange
	(*timestamppb.Timestamp)(nil),   // 7: google.protobuf.Timestamp
}
var file_automate_gateway_api_license_license_proto_depIdxs = []int32{
	3, // 0: chef.automate.api.license.ApplyLicenseResp.status:type_name -> chef.automate.api.license.GetStatusResp
	7, // 1: chef.automate.api.license.GetStatusResp.configured_at:type_name -> google.protobuf.Timestamp
	6, // 2: chef.automate.api.license.GetStatusResp.licensed_period:type_name -> chef.automate.api.license.GetStatusResp.DateRange
	3, // 3: chef.automate.api.license.RequestLicenseResp.status:type_name -> chef.automate.api.license.GetStatusResp
	7, // 4: chef.automate.api.license.GetStatusResp.DateRange.start:type_name -> google.protobuf.Timestamp
	7, // 5: chef.automate.api.license.GetStatusResp.DateRange.end:type_name -> google.protobuf.Timestamp
	0, // 6: chef.automate.api.license.License.ApplyLicense:input_type -> chef.automate.api.license.ApplyLicenseReq
	2, // 7: chef.automate.api.license.License.GetStatus:input_type -> chef.automate.api.license.GetStatusReq
	4, // 8: chef.automate.api.license.License.RequestLicense:input_type -> chef.automate.api.license.RequestLicenseReq
	1, // 9: chef.automate.api.license.License.ApplyLicense:output_type -> chef.automate.api.license.ApplyLicenseResp
	3, // 10: chef.automate.api.license.License.GetStatus:output_type -> chef.automate.api.license.GetStatusResp
	5, // 11: chef.automate.api.license.License.RequestLicense:output_type -> chef.automate.api.license.RequestLicenseResp
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_automate_gateway_api_license_license_proto_init() }
func file_automate_gateway_api_license_license_proto_init() {
	if File_automate_gateway_api_license_license_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_automate_gateway_api_license_license_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyLicenseReq); i {
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
		file_automate_gateway_api_license_license_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyLicenseResp); i {
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
		file_automate_gateway_api_license_license_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusReq); i {
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
		file_automate_gateway_api_license_license_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusResp); i {
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
		file_automate_gateway_api_license_license_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLicenseReq); i {
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
		file_automate_gateway_api_license_license_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLicenseResp); i {
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
		file_automate_gateway_api_license_license_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusResp_DateRange); i {
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
			RawDescriptor: file_automate_gateway_api_license_license_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_automate_gateway_api_license_license_proto_goTypes,
		DependencyIndexes: file_automate_gateway_api_license_license_proto_depIdxs,
		MessageInfos:      file_automate_gateway_api_license_license_proto_msgTypes,
	}.Build()
	File_automate_gateway_api_license_license_proto = out.File
	file_automate_gateway_api_license_license_proto_rawDesc = nil
	file_automate_gateway_api_license_license_proto_goTypes = nil
	file_automate_gateway_api_license_license_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LicenseClient is the client API for License service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LicenseClient interface {
	ApplyLicense(ctx context.Context, in *ApplyLicenseReq, opts ...grpc.CallOption) (*ApplyLicenseResp, error)
	GetStatus(ctx context.Context, in *GetStatusReq, opts ...grpc.CallOption) (*GetStatusResp, error)
	RequestLicense(ctx context.Context, in *RequestLicenseReq, opts ...grpc.CallOption) (*RequestLicenseResp, error)
}

type licenseClient struct {
	cc grpc.ClientConnInterface
}

func NewLicenseClient(cc grpc.ClientConnInterface) LicenseClient {
	return &licenseClient{cc}
}

func (c *licenseClient) ApplyLicense(ctx context.Context, in *ApplyLicenseReq, opts ...grpc.CallOption) (*ApplyLicenseResp, error) {
	out := new(ApplyLicenseResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.license.License/ApplyLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseClient) GetStatus(ctx context.Context, in *GetStatusReq, opts ...grpc.CallOption) (*GetStatusResp, error) {
	out := new(GetStatusResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.license.License/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseClient) RequestLicense(ctx context.Context, in *RequestLicenseReq, opts ...grpc.CallOption) (*RequestLicenseResp, error) {
	out := new(RequestLicenseResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.license.License/RequestLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LicenseServer is the server API for License service.
type LicenseServer interface {
	ApplyLicense(context.Context, *ApplyLicenseReq) (*ApplyLicenseResp, error)
	GetStatus(context.Context, *GetStatusReq) (*GetStatusResp, error)
	RequestLicense(context.Context, *RequestLicenseReq) (*RequestLicenseResp, error)
}

// UnimplementedLicenseServer can be embedded to have forward compatible implementations.
type UnimplementedLicenseServer struct {
}

func (*UnimplementedLicenseServer) ApplyLicense(context.Context, *ApplyLicenseReq) (*ApplyLicenseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyLicense not implemented")
}
func (*UnimplementedLicenseServer) GetStatus(context.Context, *GetStatusReq) (*GetStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (*UnimplementedLicenseServer) RequestLicense(context.Context, *RequestLicenseReq) (*RequestLicenseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLicense not implemented")
}

func RegisterLicenseServer(s *grpc.Server, srv LicenseServer) {
	s.RegisterService(&_License_serviceDesc, srv)
}

func _License_ApplyLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyLicenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServer).ApplyLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.license.License/ApplyLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServer).ApplyLicense(ctx, req.(*ApplyLicenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _License_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.license.License/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServer).GetStatus(ctx, req.(*GetStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _License_RequestLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestLicenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServer).RequestLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.license.License/RequestLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServer).RequestLicense(ctx, req.(*RequestLicenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _License_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.license.License",
	HandlerType: (*LicenseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyLicense",
			Handler:    _License_ApplyLicense_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _License_GetStatus_Handler,
		},
		{
			MethodName: "RequestLicense",
			Handler:    _License_RequestLicense_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "automate-gateway/api/license/license.proto",
}
