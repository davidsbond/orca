// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: orca/worker/service/v1/service.proto

package workersvcv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type RunWorkflowRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The unique identifier of the workflow run.
	WorkflowRunId string `protobuf:"bytes,1,opt,name=workflow_run_id,json=workflowRunId,proto3" json:"workflow_run_id,omitempty"`
	// The name of the workflow to run.
	WorkflowName string `protobuf:"bytes,2,opt,name=workflow_name,json=workflowName,proto3" json:"workflow_name,omitempty"`
	// Encoded input to pass to the workflow.
	Input         []byte `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RunWorkflowRequest) Reset() {
	*x = RunWorkflowRequest{}
	mi := &file_orca_worker_service_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RunWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunWorkflowRequest) ProtoMessage() {}

func (x *RunWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orca_worker_service_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunWorkflowRequest.ProtoReflect.Descriptor instead.
func (*RunWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_orca_worker_service_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *RunWorkflowRequest) GetWorkflowRunId() string {
	if x != nil {
		return x.WorkflowRunId
	}
	return ""
}

func (x *RunWorkflowRequest) GetWorkflowName() string {
	if x != nil {
		return x.WorkflowName
	}
	return ""
}

func (x *RunWorkflowRequest) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

type RunWorkflowResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RunWorkflowResponse) Reset() {
	*x = RunWorkflowResponse{}
	mi := &file_orca_worker_service_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RunWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunWorkflowResponse) ProtoMessage() {}

func (x *RunWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orca_worker_service_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunWorkflowResponse.ProtoReflect.Descriptor instead.
func (*RunWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_orca_worker_service_v1_service_proto_rawDescGZIP(), []int{1}
}

type RunTaskRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The unique identifier of the task run.
	TaskRunId string `protobuf:"bytes,1,opt,name=task_run_id,json=taskRunId,proto3" json:"task_run_id,omitempty"`
	// The name of the task to run.
	TaskName string `protobuf:"bytes,2,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	// Encoded input to pass to the task.
	Input         []byte `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RunTaskRequest) Reset() {
	*x = RunTaskRequest{}
	mi := &file_orca_worker_service_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RunTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunTaskRequest) ProtoMessage() {}

func (x *RunTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orca_worker_service_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunTaskRequest.ProtoReflect.Descriptor instead.
func (*RunTaskRequest) Descriptor() ([]byte, []int) {
	return file_orca_worker_service_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *RunTaskRequest) GetTaskRunId() string {
	if x != nil {
		return x.TaskRunId
	}
	return ""
}

func (x *RunTaskRequest) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *RunTaskRequest) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

type RunTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RunTaskResponse) Reset() {
	*x = RunTaskResponse{}
	mi := &file_orca_worker_service_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RunTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunTaskResponse) ProtoMessage() {}

func (x *RunTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orca_worker_service_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunTaskResponse.ProtoReflect.Descriptor instead.
func (*RunTaskResponse) Descriptor() ([]byte, []int) {
	return file_orca_worker_service_v1_service_proto_rawDescGZIP(), []int{3}
}

var File_orca_worker_service_v1_service_proto protoreflect.FileDescriptor

var file_orca_worker_service_v1_service_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x6f, 0x72, 0x63, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x77,
	0x0a, 0x12, 0x52, 0x75, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x75, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63,
	0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd3, 0x01, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x2a, 0x2e, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x75, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5a, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x26, 0x2e, 0x6f, 0x72,
	0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4e, 0x5a, 0x4c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x76, 0x69, 0x64,
	0x73, 0x62, 0x6f, 0x6e, 0x64, 0x2f, 0x6f, 0x72, 0x63, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x63, 0x61, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x76, 0x63, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_orca_worker_service_v1_service_proto_rawDescOnce sync.Once
	file_orca_worker_service_v1_service_proto_rawDescData []byte
)

func file_orca_worker_service_v1_service_proto_rawDescGZIP() []byte {
	file_orca_worker_service_v1_service_proto_rawDescOnce.Do(func() {
		file_orca_worker_service_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_orca_worker_service_v1_service_proto_rawDesc), len(file_orca_worker_service_v1_service_proto_rawDesc)))
	})
	return file_orca_worker_service_v1_service_proto_rawDescData
}

var file_orca_worker_service_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orca_worker_service_v1_service_proto_goTypes = []any{
	(*RunWorkflowRequest)(nil),  // 0: orca.worker.service.v1.RunWorkflowRequest
	(*RunWorkflowResponse)(nil), // 1: orca.worker.service.v1.RunWorkflowResponse
	(*RunTaskRequest)(nil),      // 2: orca.worker.service.v1.RunTaskRequest
	(*RunTaskResponse)(nil),     // 3: orca.worker.service.v1.RunTaskResponse
}
var file_orca_worker_service_v1_service_proto_depIdxs = []int32{
	0, // 0: orca.worker.service.v1.WorkerService.RunWorkflow:input_type -> orca.worker.service.v1.RunWorkflowRequest
	2, // 1: orca.worker.service.v1.WorkerService.RunTask:input_type -> orca.worker.service.v1.RunTaskRequest
	1, // 2: orca.worker.service.v1.WorkerService.RunWorkflow:output_type -> orca.worker.service.v1.RunWorkflowResponse
	3, // 3: orca.worker.service.v1.WorkerService.RunTask:output_type -> orca.worker.service.v1.RunTaskResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orca_worker_service_v1_service_proto_init() }
func file_orca_worker_service_v1_service_proto_init() {
	if File_orca_worker_service_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_orca_worker_service_v1_service_proto_rawDesc), len(file_orca_worker_service_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orca_worker_service_v1_service_proto_goTypes,
		DependencyIndexes: file_orca_worker_service_v1_service_proto_depIdxs,
		MessageInfos:      file_orca_worker_service_v1_service_proto_msgTypes,
	}.Build()
	File_orca_worker_service_v1_service_proto = out.File
	file_orca_worker_service_v1_service_proto_goTypes = nil
	file_orca_worker_service_v1_service_proto_depIdxs = nil
}
