// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: external/ingest/job_scheduler.proto

package ingest

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	request "github.com/chef/automate/api/external/ingest/request"
	response "github.com/chef/automate/api/external/ingest/response"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_external_ingest_job_scheduler_proto protoreflect.FileDescriptor

var file_external_ingest_job_scheduler_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa5, 0x08, 0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0xdd, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4a, 0x6f, 0x62, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12,
	0x37, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a, 0x6f, 0x62, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a, 0x35, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4a, 0x6f, 0x62,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x54, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30,
	0x2f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x8a, 0xb5, 0x18, 0x11, 0x0a, 0x0f, 0x72, 0x65, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x8a, 0xb5, 0x18, 0x15,
	0x12, 0x13, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x3a, 0x67, 0x65, 0x74, 0x12, 0x86, 0x02, 0x0a, 0x1e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x41, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x4d, 0x69,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0x6e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x22, 0x2c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f,
	0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2f,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x3a, 0x01, 0x2a, 0x8a, 0xb5, 0x18,
	0x11, 0x0a, 0x0f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x8a, 0xb5, 0x18, 0x18, 0x12, 0x16, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x3a, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x3a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x83,
	0x02, 0x0a, 0x1d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x12, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x40, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x22, 0x2b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2d, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x03,
	0x3a, 0x01, 0x2a, 0x8a, 0xb5, 0x18, 0x11, 0x0a, 0x0f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x3a, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x8a, 0xb5, 0x18, 0x18, 0x12, 0x16, 0x72, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x3a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0xa5, 0x02, 0x0a, 0x29, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x46, 0x6f,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x12, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x4c, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x46,
	0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x22, 0x77, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x22, 0x35, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x30, 0x2f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x2f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x2d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x3a, 0x01, 0x2a, 0x8a, 0xb5, 0x18, 0x11, 0x0a,
	0x0f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x8a, 0xb5, 0x18, 0x18, 0x12, 0x16, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x3a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x2e, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_external_ingest_job_scheduler_proto_goTypes = []interface{}{
	(*request.GetStatusJobScheduler)(nil),                      // 0: chef.automate.api.ingest.request.GetStatusJobScheduler
	(*request.SchedulerConfig)(nil),                            // 1: chef.automate.api.ingest.request.SchedulerConfig
	(*response.JobSchedulerStatus)(nil),                        // 2: chef.automate.api.ingest.response.JobSchedulerStatus
	(*response.ConfigureNodesMissingScheduler)(nil),            // 3: chef.automate.api.ingest.response.ConfigureNodesMissingScheduler
	(*response.ConfigureDeleteNodesScheduler)(nil),             // 4: chef.automate.api.ingest.response.ConfigureDeleteNodesScheduler
	(*response.ConfigureMissingNodesForDeletionScheduler)(nil), // 5: chef.automate.api.ingest.response.ConfigureMissingNodesForDeletionScheduler
}
var file_external_ingest_job_scheduler_proto_depIdxs = []int32{
	0, // 0: chef.automate.api.ingest.JobScheduler.GetStatusJobScheduler:input_type -> chef.automate.api.ingest.request.GetStatusJobScheduler
	1, // 1: chef.automate.api.ingest.JobScheduler.ConfigureNodesMissingScheduler:input_type -> chef.automate.api.ingest.request.SchedulerConfig
	1, // 2: chef.automate.api.ingest.JobScheduler.ConfigureDeleteNodesScheduler:input_type -> chef.automate.api.ingest.request.SchedulerConfig
	1, // 3: chef.automate.api.ingest.JobScheduler.ConfigureMissingNodesForDeletionScheduler:input_type -> chef.automate.api.ingest.request.SchedulerConfig
	2, // 4: chef.automate.api.ingest.JobScheduler.GetStatusJobScheduler:output_type -> chef.automate.api.ingest.response.JobSchedulerStatus
	3, // 5: chef.automate.api.ingest.JobScheduler.ConfigureNodesMissingScheduler:output_type -> chef.automate.api.ingest.response.ConfigureNodesMissingScheduler
	4, // 6: chef.automate.api.ingest.JobScheduler.ConfigureDeleteNodesScheduler:output_type -> chef.automate.api.ingest.response.ConfigureDeleteNodesScheduler
	5, // 7: chef.automate.api.ingest.JobScheduler.ConfigureMissingNodesForDeletionScheduler:output_type -> chef.automate.api.ingest.response.ConfigureMissingNodesForDeletionScheduler
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_ingest_job_scheduler_proto_init() }
func file_external_ingest_job_scheduler_proto_init() {
	if File_external_ingest_job_scheduler_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_ingest_job_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_ingest_job_scheduler_proto_goTypes,
		DependencyIndexes: file_external_ingest_job_scheduler_proto_depIdxs,
	}.Build()
	File_external_ingest_job_scheduler_proto = out.File
	file_external_ingest_job_scheduler_proto_rawDesc = nil
	file_external_ingest_job_scheduler_proto_goTypes = nil
	file_external_ingest_job_scheduler_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobSchedulerClient is the client API for JobScheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobSchedulerClient interface {
	GetStatusJobScheduler(ctx context.Context, in *request.GetStatusJobScheduler, opts ...grpc.CallOption) (*response.JobSchedulerStatus, error)
	ConfigureNodesMissingScheduler(ctx context.Context, in *request.SchedulerConfig, opts ...grpc.CallOption) (*response.ConfigureNodesMissingScheduler, error)
	ConfigureDeleteNodesScheduler(ctx context.Context, in *request.SchedulerConfig, opts ...grpc.CallOption) (*response.ConfigureDeleteNodesScheduler, error)
	ConfigureMissingNodesForDeletionScheduler(ctx context.Context, in *request.SchedulerConfig, opts ...grpc.CallOption) (*response.ConfigureMissingNodesForDeletionScheduler, error)
}

type jobSchedulerClient struct {
	cc grpc.ClientConnInterface
}

func NewJobSchedulerClient(cc grpc.ClientConnInterface) JobSchedulerClient {
	return &jobSchedulerClient{cc}
}

func (c *jobSchedulerClient) GetStatusJobScheduler(ctx context.Context, in *request.GetStatusJobScheduler, opts ...grpc.CallOption) (*response.JobSchedulerStatus, error) {
	out := new(response.JobSchedulerStatus)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.JobScheduler/GetStatusJobScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobSchedulerClient) ConfigureNodesMissingScheduler(ctx context.Context, in *request.SchedulerConfig, opts ...grpc.CallOption) (*response.ConfigureNodesMissingScheduler, error) {
	out := new(response.ConfigureNodesMissingScheduler)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.JobScheduler/ConfigureNodesMissingScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobSchedulerClient) ConfigureDeleteNodesScheduler(ctx context.Context, in *request.SchedulerConfig, opts ...grpc.CallOption) (*response.ConfigureDeleteNodesScheduler, error) {
	out := new(response.ConfigureDeleteNodesScheduler)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.JobScheduler/ConfigureDeleteNodesScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobSchedulerClient) ConfigureMissingNodesForDeletionScheduler(ctx context.Context, in *request.SchedulerConfig, opts ...grpc.CallOption) (*response.ConfigureMissingNodesForDeletionScheduler, error) {
	out := new(response.ConfigureMissingNodesForDeletionScheduler)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.JobScheduler/ConfigureMissingNodesForDeletionScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobSchedulerServer is the server API for JobScheduler service.
type JobSchedulerServer interface {
	GetStatusJobScheduler(context.Context, *request.GetStatusJobScheduler) (*response.JobSchedulerStatus, error)
	ConfigureNodesMissingScheduler(context.Context, *request.SchedulerConfig) (*response.ConfigureNodesMissingScheduler, error)
	ConfigureDeleteNodesScheduler(context.Context, *request.SchedulerConfig) (*response.ConfigureDeleteNodesScheduler, error)
	ConfigureMissingNodesForDeletionScheduler(context.Context, *request.SchedulerConfig) (*response.ConfigureMissingNodesForDeletionScheduler, error)
}

// UnimplementedJobSchedulerServer can be embedded to have forward compatible implementations.
type UnimplementedJobSchedulerServer struct {
}

func (*UnimplementedJobSchedulerServer) GetStatusJobScheduler(context.Context, *request.GetStatusJobScheduler) (*response.JobSchedulerStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatusJobScheduler not implemented")
}
func (*UnimplementedJobSchedulerServer) ConfigureNodesMissingScheduler(context.Context, *request.SchedulerConfig) (*response.ConfigureNodesMissingScheduler, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureNodesMissingScheduler not implemented")
}
func (*UnimplementedJobSchedulerServer) ConfigureDeleteNodesScheduler(context.Context, *request.SchedulerConfig) (*response.ConfigureDeleteNodesScheduler, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureDeleteNodesScheduler not implemented")
}
func (*UnimplementedJobSchedulerServer) ConfigureMissingNodesForDeletionScheduler(context.Context, *request.SchedulerConfig) (*response.ConfigureMissingNodesForDeletionScheduler, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureMissingNodesForDeletionScheduler not implemented")
}

func RegisterJobSchedulerServer(s *grpc.Server, srv JobSchedulerServer) {
	s.RegisterService(&_JobScheduler_serviceDesc, srv)
}

func _JobScheduler_GetStatusJobScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetStatusJobScheduler)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobSchedulerServer).GetStatusJobScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.JobScheduler/GetStatusJobScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobSchedulerServer).GetStatusJobScheduler(ctx, req.(*request.GetStatusJobScheduler))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobScheduler_ConfigureNodesMissingScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SchedulerConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobSchedulerServer).ConfigureNodesMissingScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.JobScheduler/ConfigureNodesMissingScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobSchedulerServer).ConfigureNodesMissingScheduler(ctx, req.(*request.SchedulerConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobScheduler_ConfigureDeleteNodesScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SchedulerConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobSchedulerServer).ConfigureDeleteNodesScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.JobScheduler/ConfigureDeleteNodesScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobSchedulerServer).ConfigureDeleteNodesScheduler(ctx, req.(*request.SchedulerConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobScheduler_ConfigureMissingNodesForDeletionScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SchedulerConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobSchedulerServer).ConfigureMissingNodesForDeletionScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.JobScheduler/ConfigureMissingNodesForDeletionScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobSchedulerServer).ConfigureMissingNodesForDeletionScheduler(ctx, req.(*request.SchedulerConfig))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobScheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.ingest.JobScheduler",
	HandlerType: (*JobSchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatusJobScheduler",
			Handler:    _JobScheduler_GetStatusJobScheduler_Handler,
		},
		{
			MethodName: "ConfigureNodesMissingScheduler",
			Handler:    _JobScheduler_ConfigureNodesMissingScheduler_Handler,
		},
		{
			MethodName: "ConfigureDeleteNodesScheduler",
			Handler:    _JobScheduler_ConfigureDeleteNodesScheduler_Handler,
		},
		{
			MethodName: "ConfigureMissingNodesForDeletionScheduler",
			Handler:    _JobScheduler_ConfigureMissingNodesForDeletionScheduler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external/ingest/job_scheduler.proto",
}
