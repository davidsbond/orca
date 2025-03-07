// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: orca/workflow/v1/workflow.proto

package workflowv1

import (
	v1 "github.com/davidsbond/orca/internal/proto/orca/task/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// The Status enumeration represents the status of a workflow run.
type Status int32

const (
	// The workflow run has an unknown status.
	Status_STATUS_UNSPECIFIED Status = 0
	// The workflow run has been created and is awaiting scheduling.
	Status_STATUS_PENDING Status = 1
	// The workflow run has been scheduled and is awaiting execution.
	Status_STATUS_SCHEDULED Status = 2
	// The workflow run is running.
	Status_STATUS_RUNNING Status = 3
	// The workflow run completed successfully.
	Status_STATUS_COMPLETE Status = 4
	// The workflow run failed.
	Status_STATUS_FAILED Status = 5
	// The workflow run was skipped due to a match on the idempotent key and workflow
	// name with a previous workflow ran.
	Status_STATUS_SKIPPED Status = 6
	// The workflow run was cancelled as its duration took longer than its deadline.
	Status_STATUS_TIMEOUT Status = 7
	// The workflow run was manually cancelled.
	Status_STATUS_CANCELLED Status = 8
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_PENDING",
		2: "STATUS_SCHEDULED",
		3: "STATUS_RUNNING",
		4: "STATUS_COMPLETE",
		5: "STATUS_FAILED",
		6: "STATUS_SKIPPED",
		7: "STATUS_TIMEOUT",
		8: "STATUS_CANCELLED",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_PENDING":     1,
		"STATUS_SCHEDULED":   2,
		"STATUS_RUNNING":     3,
		"STATUS_COMPLETE":    4,
		"STATUS_FAILED":      5,
		"STATUS_SKIPPED":     6,
		"STATUS_TIMEOUT":     7,
		"STATUS_CANCELLED":   8,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_orca_workflow_v1_workflow_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_orca_workflow_v1_workflow_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_orca_workflow_v1_workflow_proto_rawDescGZIP(), []int{0}
}

// The Run type represents the current state of a single workflow run
type Run struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The unique identifier of the workflow run.
	RunId string `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	// The unique identifier of the parent workflow run this workflow was run within. Assumed to be the root if this
	// is not provided.
	ParentWorkflowRunId string `protobuf:"bytes,2,opt,name=parent_workflow_run_id,json=parentWorkflowRunId,proto3" json:"parent_workflow_run_id,omitempty"`
	// The name of the workflow being ran.
	WorkflowName string `protobuf:"bytes,3,opt,name=workflow_name,json=workflowName,proto3" json:"workflow_name,omitempty"`
	// When the workflow run was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// When the workflow run was scheduled.
	ScheduledAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=scheduled_at,json=scheduledAt,proto3" json:"scheduled_at,omitempty"`
	// When the workflow run was started.
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// When the workflow run was completed.
	CompletedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	// When the workflow run was cancelled
	CancelledAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"`
	// The status of the workflow run.
	Status Status `protobuf:"varint,9,opt,name=status,proto3,enum=orca.workflow.v1.Status" json:"status,omitempty"`
	// The JSON encoded input provided to the workflow run.
	Input []byte `protobuf:"bytes,10,opt,name=input,proto3" json:"input,omitempty"`
	// The JSON encoded output of the workflow run.
	Output        []byte `protobuf:"bytes,11,opt,name=output,proto3" json:"output,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Run) Reset() {
	*x = Run{}
	mi := &file_orca_workflow_v1_workflow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Run) ProtoMessage() {}

func (x *Run) ProtoReflect() protoreflect.Message {
	mi := &file_orca_workflow_v1_workflow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Run.ProtoReflect.Descriptor instead.
func (*Run) Descriptor() ([]byte, []int) {
	return file_orca_workflow_v1_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *Run) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *Run) GetParentWorkflowRunId() string {
	if x != nil {
		return x.ParentWorkflowRunId
	}
	return ""
}

func (x *Run) GetWorkflowName() string {
	if x != nil {
		return x.WorkflowName
	}
	return ""
}

func (x *Run) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Run) GetScheduledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledAt
	}
	return nil
}

func (x *Run) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Run) GetCompletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletedAt
	}
	return nil
}

func (x *Run) GetCancelledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CancelledAt
	}
	return nil
}

func (x *Run) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *Run) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Run) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

// The WorkflowRunDescription type represents full details of a workflow run. Including the child workflows and
// tasks that it executed.
type WorkflowRunDescription struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The workflow run data.
	Run *Run `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"`
	// The child workflows and tasks performed during this run. These will be ordered by their creation.
	Actions       []*WorkflowAction `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowRunDescription) Reset() {
	*x = WorkflowRunDescription{}
	mi := &file_orca_workflow_v1_workflow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowRunDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRunDescription) ProtoMessage() {}

func (x *WorkflowRunDescription) ProtoReflect() protoreflect.Message {
	mi := &file_orca_workflow_v1_workflow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRunDescription.ProtoReflect.Descriptor instead.
func (*WorkflowRunDescription) Descriptor() ([]byte, []int) {
	return file_orca_workflow_v1_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowRunDescription) GetRun() *Run {
	if x != nil {
		return x.Run
	}
	return nil
}

func (x *WorkflowRunDescription) GetActions() []*WorkflowAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

// The WorkflowAction type describes an action taken by a workflow. This is typically a child workflow or a task.
type WorkflowAction struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Action:
	//
	//	*WorkflowAction_WorkflowRun
	//	*WorkflowAction_TaskRun
	Action        isWorkflowAction_Action `protobuf_oneof:"action"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowAction) Reset() {
	*x = WorkflowAction{}
	mi := &file_orca_workflow_v1_workflow_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowAction) ProtoMessage() {}

func (x *WorkflowAction) ProtoReflect() protoreflect.Message {
	mi := &file_orca_workflow_v1_workflow_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowAction.ProtoReflect.Descriptor instead.
func (*WorkflowAction) Descriptor() ([]byte, []int) {
	return file_orca_workflow_v1_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *WorkflowAction) GetAction() isWorkflowAction_Action {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *WorkflowAction) GetWorkflowRun() *WorkflowRunDescription {
	if x != nil {
		if x, ok := x.Action.(*WorkflowAction_WorkflowRun); ok {
			return x.WorkflowRun
		}
	}
	return nil
}

func (x *WorkflowAction) GetTaskRun() *v1.TaskRunDescription {
	if x != nil {
		if x, ok := x.Action.(*WorkflowAction_TaskRun); ok {
			return x.TaskRun
		}
	}
	return nil
}

type isWorkflowAction_Action interface {
	isWorkflowAction_Action()
}

type WorkflowAction_WorkflowRun struct {
	// A child workflow the parent ran.
	WorkflowRun *WorkflowRunDescription `protobuf:"bytes,1,opt,name=workflow_run,json=workflowRun,proto3,oneof"`
}

type WorkflowAction_TaskRun struct {
	// A task the parent ran.
	TaskRun *v1.TaskRunDescription `protobuf:"bytes,2,opt,name=task_run,json=taskRun,proto3,oneof"`
}

func (*WorkflowAction_WorkflowRun) isWorkflowAction_Action() {}

func (*WorkflowAction_TaskRun) isWorkflowAction_Action() {}

var File_orca_workflow_v1_workflow_proto protoreflect.FileDescriptor

var file_orca_workflow_v1_workflow_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x6f, 0x72, 0x63, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6f, 0x72, 0x63, 0x61, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x04,
	0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x16,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6f, 0x72, 0x63, 0x61,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x7d, 0x0a, 0x16, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x03, 0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x03, 0x72, 0x75, 0x6e, 0x12, 0x3a, 0x0a, 0x07,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6f, 0x72, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x0e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x0c, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x12, 0x3d, 0x0a, 0x08, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f,
	0x72, 0x63, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x75, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2a, 0xc4, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x4b, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x06,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x08, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x76, 0x69, 0x64, 0x73, 0x62,
	0x6f, 0x6e, 0x64, 0x2f, 0x6f, 0x72, 0x63, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x63, 0x61, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_orca_workflow_v1_workflow_proto_rawDescOnce sync.Once
	file_orca_workflow_v1_workflow_proto_rawDescData []byte
)

func file_orca_workflow_v1_workflow_proto_rawDescGZIP() []byte {
	file_orca_workflow_v1_workflow_proto_rawDescOnce.Do(func() {
		file_orca_workflow_v1_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_orca_workflow_v1_workflow_proto_rawDesc), len(file_orca_workflow_v1_workflow_proto_rawDesc)))
	})
	return file_orca_workflow_v1_workflow_proto_rawDescData
}

var file_orca_workflow_v1_workflow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orca_workflow_v1_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orca_workflow_v1_workflow_proto_goTypes = []any{
	(Status)(0),                    // 0: orca.workflow.v1.Status
	(*Run)(nil),                    // 1: orca.workflow.v1.Run
	(*WorkflowRunDescription)(nil), // 2: orca.workflow.v1.WorkflowRunDescription
	(*WorkflowAction)(nil),         // 3: orca.workflow.v1.WorkflowAction
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*v1.TaskRunDescription)(nil),  // 5: orca.task.v1.TaskRunDescription
}
var file_orca_workflow_v1_workflow_proto_depIdxs = []int32{
	4,  // 0: orca.workflow.v1.Run.created_at:type_name -> google.protobuf.Timestamp
	4,  // 1: orca.workflow.v1.Run.scheduled_at:type_name -> google.protobuf.Timestamp
	4,  // 2: orca.workflow.v1.Run.started_at:type_name -> google.protobuf.Timestamp
	4,  // 3: orca.workflow.v1.Run.completed_at:type_name -> google.protobuf.Timestamp
	4,  // 4: orca.workflow.v1.Run.cancelled_at:type_name -> google.protobuf.Timestamp
	0,  // 5: orca.workflow.v1.Run.status:type_name -> orca.workflow.v1.Status
	1,  // 6: orca.workflow.v1.WorkflowRunDescription.run:type_name -> orca.workflow.v1.Run
	3,  // 7: orca.workflow.v1.WorkflowRunDescription.actions:type_name -> orca.workflow.v1.WorkflowAction
	2,  // 8: orca.workflow.v1.WorkflowAction.workflow_run:type_name -> orca.workflow.v1.WorkflowRunDescription
	5,  // 9: orca.workflow.v1.WorkflowAction.task_run:type_name -> orca.task.v1.TaskRunDescription
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_orca_workflow_v1_workflow_proto_init() }
func file_orca_workflow_v1_workflow_proto_init() {
	if File_orca_workflow_v1_workflow_proto != nil {
		return
	}
	file_orca_workflow_v1_workflow_proto_msgTypes[2].OneofWrappers = []any{
		(*WorkflowAction_WorkflowRun)(nil),
		(*WorkflowAction_TaskRun)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_orca_workflow_v1_workflow_proto_rawDesc), len(file_orca_workflow_v1_workflow_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orca_workflow_v1_workflow_proto_goTypes,
		DependencyIndexes: file_orca_workflow_v1_workflow_proto_depIdxs,
		EnumInfos:         file_orca_workflow_v1_workflow_proto_enumTypes,
		MessageInfos:      file_orca_workflow_v1_workflow_proto_msgTypes,
	}.Build()
	File_orca_workflow_v1_workflow_proto = out.File
	file_orca_workflow_v1_workflow_proto_goTypes = nil
	file_orca_workflow_v1_workflow_proto_depIdxs = nil
}
