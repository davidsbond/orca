// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: orca/workflow/service/v1/service.proto

package workflowsvcv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WorkflowService_Schedule_FullMethodName = "/orca.workflow.service.v1.WorkflowService/Schedule"
	WorkflowService_GetRun_FullMethodName   = "/orca.workflow.service.v1.WorkflowService/GetRun"
)

// WorkflowServiceClient is the client API for WorkflowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkflowServiceClient interface {
	// Schedule a workflow.
	Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
	// Get a workflow run.
	GetRun(ctx context.Context, in *GetRunRequest, opts ...grpc.CallOption) (*GetRunResponse, error)
}

type workflowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowServiceClient(cc grpc.ClientConnInterface) WorkflowServiceClient {
	return &workflowServiceClient{cc}
}

func (c *workflowServiceClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScheduleResponse)
	err := c.cc.Invoke(ctx, WorkflowService_Schedule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) GetRun(ctx context.Context, in *GetRunRequest, opts ...grpc.CallOption) (*GetRunResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRunResponse)
	err := c.cc.Invoke(ctx, WorkflowService_GetRun_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServiceServer is the server API for WorkflowService service.
// All implementations must embed UnimplementedWorkflowServiceServer
// for forward compatibility.
type WorkflowServiceServer interface {
	// Schedule a workflow.
	Schedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
	// Get a workflow run.
	GetRun(context.Context, *GetRunRequest) (*GetRunResponse, error)
	mustEmbedUnimplementedWorkflowServiceServer()
}

// UnimplementedWorkflowServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWorkflowServiceServer struct{}

func (UnimplementedWorkflowServiceServer) Schedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Schedule not implemented")
}
func (UnimplementedWorkflowServiceServer) GetRun(context.Context, *GetRunRequest) (*GetRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRun not implemented")
}
func (UnimplementedWorkflowServiceServer) mustEmbedUnimplementedWorkflowServiceServer() {}
func (UnimplementedWorkflowServiceServer) testEmbeddedByValue()                         {}

// UnsafeWorkflowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowServiceServer will
// result in compilation errors.
type UnsafeWorkflowServiceServer interface {
	mustEmbedUnimplementedWorkflowServiceServer()
}

func RegisterWorkflowServiceServer(s grpc.ServiceRegistrar, srv WorkflowServiceServer) {
	// If the following call pancis, it indicates UnimplementedWorkflowServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WorkflowService_ServiceDesc, srv)
}

func _WorkflowService_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).Schedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowService_Schedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).Schedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_GetRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).GetRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowService_GetRun_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).GetRun(ctx, req.(*GetRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkflowService_ServiceDesc is the grpc.ServiceDesc for WorkflowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkflowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orca.workflow.service.v1.WorkflowService",
	HandlerType: (*WorkflowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Schedule",
			Handler:    _WorkflowService_Schedule_Handler,
		},
		{
			MethodName: "GetRun",
			Handler:    _WorkflowService_GetRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orca/workflow/service/v1/service.proto",
}
