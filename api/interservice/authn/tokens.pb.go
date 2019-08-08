// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/authn/tokens.proto

package authn

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateTokenReq struct {
	// Match either a completely empty string or a valid id.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	Active               bool     `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty" toml:"active,omitempty" mapstructure:"active,omitempty"`
	Projects             []string `protobuf:"bytes,4,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateTokenReq) Reset()         { *m = CreateTokenReq{} }
func (m *CreateTokenReq) String() string { return proto.CompactTextString(m) }
func (*CreateTokenReq) ProtoMessage()    {}
func (*CreateTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{0}
}

func (m *CreateTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenReq.Unmarshal(m, b)
}
func (m *CreateTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenReq.Marshal(b, m, deterministic)
}
func (m *CreateTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenReq.Merge(m, src)
}
func (m *CreateTokenReq) XXX_Size() int {
	return xxx_messageInfo_CreateTokenReq.Size(m)
}
func (m *CreateTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenReq proto.InternalMessageInfo

func (m *CreateTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateTokenReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateTokenReq) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *CreateTokenReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type CreateTokenWithValueReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	Active               bool     `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty" toml:"active,omitempty" mapstructure:"active,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty" toml:"value,omitempty" mapstructure:"value,omitempty"`
	Projects             []string `protobuf:"bytes,5,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateTokenWithValueReq) Reset()         { *m = CreateTokenWithValueReq{} }
func (m *CreateTokenWithValueReq) String() string { return proto.CompactTextString(m) }
func (*CreateTokenWithValueReq) ProtoMessage()    {}
func (*CreateTokenWithValueReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{1}
}

func (m *CreateTokenWithValueReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenWithValueReq.Unmarshal(m, b)
}
func (m *CreateTokenWithValueReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenWithValueReq.Marshal(b, m, deterministic)
}
func (m *CreateTokenWithValueReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenWithValueReq.Merge(m, src)
}
func (m *CreateTokenWithValueReq) XXX_Size() int {
	return xxx_messageInfo_CreateTokenWithValueReq.Size(m)
}
func (m *CreateTokenWithValueReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenWithValueReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenWithValueReq proto.InternalMessageInfo

func (m *CreateTokenWithValueReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateTokenWithValueReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateTokenWithValueReq) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *CreateTokenWithValueReq) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CreateTokenWithValueReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type UpdateTokenReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	Active               bool     `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty" toml:"active,omitempty" mapstructure:"active,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	Projects             []string `protobuf:"bytes,4,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *UpdateTokenReq) Reset()         { *m = UpdateTokenReq{} }
func (m *UpdateTokenReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTokenReq) ProtoMessage()    {}
func (*UpdateTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{2}
}

func (m *UpdateTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTokenReq.Unmarshal(m, b)
}
func (m *UpdateTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTokenReq.Marshal(b, m, deterministic)
}
func (m *UpdateTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTokenReq.Merge(m, src)
}
func (m *UpdateTokenReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTokenReq.Size(m)
}
func (m *UpdateTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTokenReq proto.InternalMessageInfo

func (m *UpdateTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTokenReq) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *UpdateTokenReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateTokenReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Token struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty" toml:"value,omitempty" mapstructure:"value,omitempty"`
	Active               bool     `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty" toml:"active,omitempty" mapstructure:"active,omitempty"`
	Created              string   `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty" toml:"created,omitempty" mapstructure:"created,omitempty"`
	Updated              string   `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty" toml:"updated,omitempty" mapstructure:"updated,omitempty"`
	Projects             []string `protobuf:"bytes,7,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Token) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Token) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Token) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Token) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

func (m *Token) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Tokens struct {
	Tokens               []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty" toml:"tokens,omitempty" mapstructure:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Tokens) Reset()         { *m = Tokens{} }
func (m *Tokens) String() string { return proto.CompactTextString(m) }
func (*Tokens) ProtoMessage()    {}
func (*Tokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{4}
}

func (m *Tokens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tokens.Unmarshal(m, b)
}
func (m *Tokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tokens.Marshal(b, m, deterministic)
}
func (m *Tokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tokens.Merge(m, src)
}
func (m *Tokens) XXX_Size() int {
	return xxx_messageInfo_Tokens.Size(m)
}
func (m *Tokens) XXX_DiscardUnknown() {
	xxx_messageInfo_Tokens.DiscardUnknown(m)
}

var xxx_messageInfo_Tokens proto.InternalMessageInfo

func (m *Tokens) GetTokens() []*Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type Value struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty" toml:"value,omitempty" mapstructure:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{5}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetTokenReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *GetTokenReq) Reset()         { *m = GetTokenReq{} }
func (m *GetTokenReq) String() string { return proto.CompactTextString(m) }
func (*GetTokenReq) ProtoMessage()    {}
func (*GetTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{6}
}

func (m *GetTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenReq.Unmarshal(m, b)
}
func (m *GetTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenReq.Marshal(b, m, deterministic)
}
func (m *GetTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenReq.Merge(m, src)
}
func (m *GetTokenReq) XXX_Size() int {
	return xxx_messageInfo_GetTokenReq.Size(m)
}
func (m *GetTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenReq proto.InternalMessageInfo

func (m *GetTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTokensReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *GetTokensReq) Reset()         { *m = GetTokensReq{} }
func (m *GetTokensReq) String() string { return proto.CompactTextString(m) }
func (*GetTokensReq) ProtoMessage()    {}
func (*GetTokensReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{7}
}

func (m *GetTokensReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokensReq.Unmarshal(m, b)
}
func (m *GetTokensReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokensReq.Marshal(b, m, deterministic)
}
func (m *GetTokensReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokensReq.Merge(m, src)
}
func (m *GetTokensReq) XXX_Size() int {
	return xxx_messageInfo_GetTokensReq.Size(m)
}
func (m *GetTokensReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokensReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokensReq proto.InternalMessageInfo

type DeleteTokenReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DeleteTokenReq) Reset()         { *m = DeleteTokenReq{} }
func (m *DeleteTokenReq) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenReq) ProtoMessage()    {}
func (*DeleteTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{8}
}

func (m *DeleteTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenReq.Unmarshal(m, b)
}
func (m *DeleteTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenReq.Marshal(b, m, deterministic)
}
func (m *DeleteTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenReq.Merge(m, src)
}
func (m *DeleteTokenReq) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenReq.Size(m)
}
func (m *DeleteTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenReq proto.InternalMessageInfo

func (m *DeleteTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteTokenResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DeleteTokenResp) Reset()         { *m = DeleteTokenResp{} }
func (m *DeleteTokenResp) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenResp) ProtoMessage()    {}
func (*DeleteTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_65ff3d3f7473d71c, []int{9}
}

func (m *DeleteTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenResp.Unmarshal(m, b)
}
func (m *DeleteTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenResp.Marshal(b, m, deterministic)
}
func (m *DeleteTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenResp.Merge(m, src)
}
func (m *DeleteTokenResp) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenResp.Size(m)
}
func (m *DeleteTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateTokenReq)(nil), "chef.automate.domain.authn.CreateTokenReq")
	proto.RegisterType((*CreateTokenWithValueReq)(nil), "chef.automate.domain.authn.CreateTokenWithValueReq")
	proto.RegisterType((*UpdateTokenReq)(nil), "chef.automate.domain.authn.UpdateTokenReq")
	proto.RegisterType((*Token)(nil), "chef.automate.domain.authn.Token")
	proto.RegisterType((*Tokens)(nil), "chef.automate.domain.authn.Tokens")
	proto.RegisterType((*Value)(nil), "chef.automate.domain.authn.Value")
	proto.RegisterType((*GetTokenReq)(nil), "chef.automate.domain.authn.GetTokenReq")
	proto.RegisterType((*GetTokensReq)(nil), "chef.automate.domain.authn.GetTokensReq")
	proto.RegisterType((*DeleteTokenReq)(nil), "chef.automate.domain.authn.DeleteTokenReq")
	proto.RegisterType((*DeleteTokenResp)(nil), "chef.automate.domain.authn.DeleteTokenResp")
}

func init() {
	proto.RegisterFile("api/interservice/authn/tokens.proto", fileDescriptor_65ff3d3f7473d71c)
}

var fileDescriptor_65ff3d3f7473d71c = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0xe3, 0x24, 0x4d, 0xae, 0x3f, 0xe5, 0x13, 0xa3, 0xd0, 0x18, 0x8b, 0x20, 0x33, 0x54,
	0x22, 0x02, 0x62, 0xd3, 0x14, 0x21, 0x75, 0xc3, 0x22, 0x45, 0x62, 0xc5, 0xc6, 0xe2, 0x4f, 0x54,
	0x04, 0x4d, 0xed, 0x21, 0x19, 0x48, 0x3c, 0xae, 0x3d, 0xc9, 0x82, 0x9f, 0x97, 0x60, 0xc9, 0x8b,
	0x20, 0xba, 0xe2, 0x4d, 0x58, 0xf3, 0x16, 0xc8, 0x76, 0x9c, 0x8c, 0xdb, 0x26, 0xb5, 0xda, 0x5d,
	0x26, 0xf7, 0x9e, 0x7b, 0xcf, 0x39, 0x77, 0xae, 0x07, 0xee, 0x90, 0x80, 0xd9, 0xcc, 0x17, 0x34,
	0x8c, 0x68, 0x38, 0x67, 0x2e, 0xb5, 0xc9, 0x4c, 0x8c, 0x7d, 0x5b, 0xf0, 0x4f, 0xd4, 0x8f, 0xac,
	0x20, 0xe4, 0x82, 0x23, 0xc3, 0x1d, 0xd3, 0x0f, 0x16, 0x99, 0x09, 0x3e, 0x25, 0x82, 0x5a, 0x1e,
	0x9f, 0x12, 0xe6, 0x5b, 0x49, 0xa2, 0x71, 0x73, 0xc4, 0xf9, 0x68, 0x42, 0xed, 0xb8, 0x0e, 0xf1,
	0x7d, 0x2e, 0x88, 0x60, 0x3c, 0x43, 0x1a, 0xed, 0x39, 0x99, 0x30, 0x8f, 0x08, 0x6a, 0x67, 0x3f,
	0xd2, 0x00, 0x3e, 0x51, 0xa0, 0x79, 0x10, 0x52, 0x22, 0xe8, 0x8b, 0xb8, 0x93, 0x43, 0x8f, 0x91,
	0x05, 0x65, 0xe6, 0xe9, 0x8a, 0xa9, 0x74, 0x1b, 0x83, 0x5b, 0x27, 0x7f, 0x7f, 0xab, 0x37, 0xc2,
	0x76, 0xff, 0xfa, 0xf0, 0x90, 0xf4, 0x3e, 0x3f, 0xec, 0xed, 0xf7, 0xde, 0xbf, 0xfb, 0xb2, 0xfb,
	0xe0, 0xf1, 0xa3, 0x6f, 0x3b, 0x5f, 0x87, 0x3b, 0x4e, 0x99, 0x79, 0xc8, 0x04, 0xcd, 0xa3, 0x91,
	0x1b, 0xb2, 0x20, 0xee, 0xa8, 0x97, 0x63, 0xa0, 0x23, 0xff, 0x85, 0xb6, 0xa1, 0x46, 0x5c, 0xc1,
	0xe6, 0x54, 0x57, 0x4d, 0xa5, 0x5b, 0x77, 0x16, 0x27, 0xf4, 0x04, 0xea, 0x41, 0xc8, 0x3f, 0x52,
	0x57, 0x44, 0x7a, 0xc5, 0x54, 0xbb, 0x8d, 0x01, 0x8e, 0xfb, 0x75, 0xbe, 0x2b, 0x86, 0xae, 0xe0,
	0xed, 0xb0, 0xd5, 0x47, 0x67, 0xdb, 0x3a, 0x4b, 0x0c, 0xfe, 0xa5, 0x40, 0x5b, 0x22, 0xff, 0x9a,
	0x89, 0xf1, 0x2b, 0x32, 0x99, 0xd1, 0x58, 0x45, 0x73, 0xa5, 0xe2, 0x8a, 0x2c, 0x5b, 0x50, 0x9d,
	0xc7, 0x55, 0xf5, 0x4a, 0x82, 0x49, 0x0f, 0x39, 0xee, 0xd5, 0x4b, 0x70, 0xff, 0xa1, 0x40, 0xf3,
	0x65, 0xe0, 0xc9, 0xc6, 0x9f, 0xa6, 0xbc, 0x22, 0x54, 0xce, 0x11, 0x3a, 0x25, 0x45, 0x3d, 0x2b,
	0xe5, 0xaa, 0xc6, 0xfe, 0x54, 0xa0, 0x9a, 0xd0, 0xba, 0x84, 0x8d, 0x4b, 0xbb, 0x54, 0xd9, 0xae,
	0x95, 0x96, 0x4a, 0x4e, 0x8b, 0x0e, 0x5b, 0x6e, 0x32, 0x41, 0x4f, 0xaf, 0x26, 0xf9, 0xd9, 0x31,
	0x8e, 0xcc, 0x12, 0x7f, 0x3c, 0xbd, 0x96, 0x46, 0x16, 0x47, 0x64, 0x48, 0xea, 0xb6, 0x62, 0x75,
	0x12, 0xf3, 0x03, 0xa8, 0x25, 0xc4, 0x23, 0xb4, 0x0f, 0xb5, 0x74, 0x79, 0x74, 0xc5, 0x54, 0xbb,
	0x5a, 0xff, 0xb6, 0xb5, 0x7e, 0x7b, 0xac, 0x74, 0x06, 0x0b, 0x00, 0xee, 0x40, 0x35, 0xb9, 0x47,
	0x2b, 0x2d, 0x8a, 0xa4, 0x05, 0x77, 0x40, 0x7b, 0x46, 0xc5, 0xba, 0xb1, 0xe1, 0x26, 0xfc, 0x97,
	0x85, 0x23, 0x87, 0x1e, 0x63, 0x13, 0x9a, 0x4f, 0xe9, 0x84, 0xae, 0x1f, 0x34, 0xbe, 0x06, 0xff,
	0xe7, 0x32, 0xa2, 0xa0, 0xff, 0xa7, 0x02, 0x90, 0x96, 0x78, 0x3e, 0x9a, 0x0a, 0x74, 0x08, 0x8d,
	0x65, 0x4d, 0xd4, 0xdd, 0xa4, 0x44, 0x6e, 0x6d, 0xe0, 0x0b, 0x35, 0x47, 0xb8, 0x84, 0x86, 0xa0,
	0x49, 0x5b, 0x84, 0xee, 0x6d, 0x02, 0xe5, 0xbf, 0x15, 0xc6, 0xc5, 0xa6, 0xe2, 0x12, 0x0a, 0xa0,
	0x75, 0xde, 0x96, 0xa2, 0xbd, 0x82, 0x8d, 0xe4, 0xbd, 0x2e, 0xd6, 0x71, 0x08, 0x9a, 0xb4, 0x5b,
	0x9b, 0x15, 0xe5, 0x97, 0xb0, 0x58, 0xfd, 0x37, 0x50, 0xcf, 0x7c, 0x46, 0x77, 0x8b, 0x4c, 0xa3,
	0x70, 0xe5, 0x31, 0x68, 0xd2, 0x55, 0xd8, 0xcc, 0x3c, 0x7f, 0xab, 0x8c, 0xfb, 0x85, 0x73, 0xa3,
	0x00, 0x97, 0x06, 0xbb, 0x6f, 0xed, 0x11, 0x13, 0xe3, 0xd9, 0x91, 0xe5, 0xf2, 0xa9, 0x1d, 0x43,
	0xed, 0x0c, 0x6a, 0x9f, 0xff, 0x18, 0x1d, 0xd5, 0x92, 0x37, 0x63, 0xef, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xeb, 0x15, 0x99, 0xa6, 0xad, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TokensMgmtClient is the client API for TokensMgmt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokensMgmtClient interface {
	GetTokens(ctx context.Context, in *GetTokensReq, opts ...grpc.CallOption) (*Tokens, error)
	CreateToken(ctx context.Context, in *CreateTokenReq, opts ...grpc.CallOption) (*Token, error)
	CreateTokenWithValue(ctx context.Context, in *CreateTokenWithValueReq, opts ...grpc.CallOption) (*Token, error)
	UpdateToken(ctx context.Context, in *UpdateTokenReq, opts ...grpc.CallOption) (*Token, error)
	GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*Token, error)
	DeleteToken(ctx context.Context, in *DeleteTokenReq, opts ...grpc.CallOption) (*DeleteTokenResp, error)
}

type tokensMgmtClient struct {
	cc *grpc.ClientConn
}

func NewTokensMgmtClient(cc *grpc.ClientConn) TokensMgmtClient {
	return &tokensMgmtClient{cc}
}

func (c *tokensMgmtClient) GetTokens(ctx context.Context, in *GetTokensReq, opts ...grpc.CallOption) (*Tokens, error) {
	out := new(Tokens)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/GetTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensMgmtClient) CreateToken(ctx context.Context, in *CreateTokenReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensMgmtClient) CreateTokenWithValue(ctx context.Context, in *CreateTokenWithValueReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/CreateTokenWithValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensMgmtClient) UpdateToken(ctx context.Context, in *UpdateTokenReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/UpdateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensMgmtClient) GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensMgmtClient) DeleteToken(ctx context.Context, in *DeleteTokenReq, opts ...grpc.CallOption) (*DeleteTokenResp, error) {
	out := new(DeleteTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.TokensMgmt/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokensMgmtServer is the server API for TokensMgmt service.
type TokensMgmtServer interface {
	GetTokens(context.Context, *GetTokensReq) (*Tokens, error)
	CreateToken(context.Context, *CreateTokenReq) (*Token, error)
	CreateTokenWithValue(context.Context, *CreateTokenWithValueReq) (*Token, error)
	UpdateToken(context.Context, *UpdateTokenReq) (*Token, error)
	GetToken(context.Context, *GetTokenReq) (*Token, error)
	DeleteToken(context.Context, *DeleteTokenReq) (*DeleteTokenResp, error)
}

func RegisterTokensMgmtServer(s *grpc.Server, srv TokensMgmtServer) {
	s.RegisterService(&_TokensMgmt_serviceDesc, srv)
}

func _TokensMgmt_GetTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokensReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).GetTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/GetTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).GetTokens(ctx, req.(*GetTokensReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokensMgmt_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).CreateToken(ctx, req.(*CreateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokensMgmt_CreateTokenWithValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenWithValueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).CreateTokenWithValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/CreateTokenWithValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).CreateTokenWithValue(ctx, req.(*CreateTokenWithValueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokensMgmt_UpdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).UpdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/UpdateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).UpdateToken(ctx, req.(*UpdateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokensMgmt_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).GetToken(ctx, req.(*GetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokensMgmt_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensMgmtServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.TokensMgmt/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensMgmtServer).DeleteToken(ctx, req.(*DeleteTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokensMgmt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.authn.TokensMgmt",
	HandlerType: (*TokensMgmtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTokens",
			Handler:    _TokensMgmt_GetTokens_Handler,
		},
		{
			MethodName: "CreateToken",
			Handler:    _TokensMgmt_CreateToken_Handler,
		},
		{
			MethodName: "CreateTokenWithValue",
			Handler:    _TokensMgmt_CreateTokenWithValue_Handler,
		},
		{
			MethodName: "UpdateToken",
			Handler:    _TokensMgmt_UpdateToken_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _TokensMgmt_GetToken_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _TokensMgmt_DeleteToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interservice/authn/tokens.proto",
}
