// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: external/cfgmgmt/request/stats.proto

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

type NodesCounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of filters to be applied to the node count results.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	// Earliest node check-in.
	Start string `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	// Latest node check-in
	End string `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *NodesCounts) Reset() {
	*x = NodesCounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_request_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodesCounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodesCounts) ProtoMessage() {}

func (x *NodesCounts) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_request_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodesCounts.ProtoReflect.Descriptor instead.
func (*NodesCounts) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_request_stats_proto_rawDescGZIP(), []int{0}
}

func (x *NodesCounts) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *NodesCounts) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *NodesCounts) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type RunsCounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of filters to be applied to the run count results.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	// Earliest (in history) run information to return for the run counts.
	Start string `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	// Latest (in history) run information to return for the run counts.
	End string `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	// Node id associated with the run.
	NodeId string `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *RunsCounts) Reset() {
	*x = RunsCounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_request_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunsCounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunsCounts) ProtoMessage() {}

func (x *RunsCounts) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_request_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunsCounts.ProtoReflect.Descriptor instead.
func (*RunsCounts) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_request_stats_proto_rawDescGZIP(), []int{1}
}

func (x *RunsCounts) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *RunsCounts) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *RunsCounts) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *RunsCounts) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type CheckInCountsTimeSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of filters to be applied to the time series.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	// Number of past days to create the time series
	DaysAgo int32 `protobuf:"varint,2,opt,name=days_ago,json=daysAgo,proto3" json:"days_ago,omitempty"`
}

func (x *CheckInCountsTimeSeries) Reset() {
	*x = CheckInCountsTimeSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_request_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInCountsTimeSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInCountsTimeSeries) ProtoMessage() {}

func (x *CheckInCountsTimeSeries) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_request_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInCountsTimeSeries.ProtoReflect.Descriptor instead.
func (*CheckInCountsTimeSeries) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_request_stats_proto_rawDescGZIP(), []int{2}
}

func (x *CheckInCountsTimeSeries) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *CheckInCountsTimeSeries) GetDaysAgo() int32 {
	if x != nil {
		return x.DaysAgo
	}
	return 0
}

var File_external_cfgmgmt_request_stats_proto protoreflect.FileDescriptor

var file_external_cfgmgmt_request_stats_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x3a, 0x3d, 0x92, 0x41, 0x3a, 0x0a, 0x09, 0xd2, 0x01,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x32, 0x2d, 0x12, 0x2b, 0x7b, 0x22, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x6d, 0x79, 0x53,
	0x4f, 0x2a, 0x22, 0x2c, 0x22, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x75, 0x62,
	0x75, 0x6e, 0x2a, 0x22, 0x5d, 0x7d, 0x22, 0xc4, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x3a, 0x5d,
	0x92, 0x41, 0x5a, 0x0a, 0x0a, 0xd2, 0x01, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x32,
	0x4c, 0x12, 0x4a, 0x7b, 0x22, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22,
	0x38, 0x32, 0x31, 0x66, 0x66, 0x66, 0x30, 0x37, 0x2d, 0x61, 0x62, 0x63, 0x39, 0x2d, 0x34, 0x31,
	0x36, 0x30, 0x2d, 0x39, 0x36, 0x62, 0x31, 0x2d, 0x38, 0x33, 0x64, 0x36, 0x38, 0x61, 0x65, 0x35,
	0x63, 0x66, 0x64, 0x64, 0x22, 0x2c, 0x20, 0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x22, 0x3a, 0x20,
	0x22, 0x32, 0x30, 0x31, 0x39, 0x2d, 0x31, 0x31, 0x2d, 0x30, 0x32, 0x22, 0x7d, 0x22, 0x4c, 0x0a,
	0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x61, 0x67, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x64, 0x61, 0x79, 0x73, 0x41, 0x67, 0x6f, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_cfgmgmt_request_stats_proto_rawDescOnce sync.Once
	file_external_cfgmgmt_request_stats_proto_rawDescData = file_external_cfgmgmt_request_stats_proto_rawDesc
)

func file_external_cfgmgmt_request_stats_proto_rawDescGZIP() []byte {
	file_external_cfgmgmt_request_stats_proto_rawDescOnce.Do(func() {
		file_external_cfgmgmt_request_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_cfgmgmt_request_stats_proto_rawDescData)
	})
	return file_external_cfgmgmt_request_stats_proto_rawDescData
}

var file_external_cfgmgmt_request_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_external_cfgmgmt_request_stats_proto_goTypes = []interface{}{
	(*NodesCounts)(nil),             // 0: chef.automate.api.cfgmgmt.request.NodesCounts
	(*RunsCounts)(nil),              // 1: chef.automate.api.cfgmgmt.request.RunsCounts
	(*CheckInCountsTimeSeries)(nil), // 2: chef.automate.api.cfgmgmt.request.CheckInCountsTimeSeries
}
var file_external_cfgmgmt_request_stats_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_cfgmgmt_request_stats_proto_init() }
func file_external_cfgmgmt_request_stats_proto_init() {
	if File_external_cfgmgmt_request_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_cfgmgmt_request_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodesCounts); i {
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
		file_external_cfgmgmt_request_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunsCounts); i {
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
		file_external_cfgmgmt_request_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInCountsTimeSeries); i {
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
			RawDescriptor: file_external_cfgmgmt_request_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_cfgmgmt_request_stats_proto_goTypes,
		DependencyIndexes: file_external_cfgmgmt_request_stats_proto_depIdxs,
		MessageInfos:      file_external_cfgmgmt_request_stats_proto_msgTypes,
	}.Build()
	File_external_cfgmgmt_request_stats_proto = out.File
	file_external_cfgmgmt_request_stats_proto_rawDesc = nil
	file_external_cfgmgmt_request_stats_proto_goTypes = nil
	file_external_cfgmgmt_request_stats_proto_depIdxs = nil
}
