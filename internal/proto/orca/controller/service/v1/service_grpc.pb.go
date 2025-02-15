// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: orca/controller/service/v1/service.proto

package controllersvcv1

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
	ControllerService_RegisterWorker_FullMethodName       = "/orca.controller.service.v1.ControllerService/RegisterWorker"
	ControllerService_DeregisterWorker_FullMethodName     = "/orca.controller.service.v1.ControllerService/DeregisterWorker"
	ControllerService_ScheduleTask_FullMethodName         = "/orca.controller.service.v1.ControllerService/ScheduleTask"
	ControllerService_SetTaskRunStatus_FullMethodName     = "/orca.controller.service.v1.ControllerService/SetTaskRunStatus"
	ControllerService_GetTaskRun_FullMethodName           = "/orca.controller.service.v1.ControllerService/GetTaskRun"
	ControllerService_ScheduleWorkflow_FullMethodName     = "/orca.controller.service.v1.ControllerService/ScheduleWorkflow"
	ControllerService_SetWorkflowRunStatus_FullMethodName = "/orca.controller.service.v1.ControllerService/SetWorkflowRunStatus"
	ControllerService_GetWorkflowRun_FullMethodName       = "/orca.controller.service.v1.ControllerService/GetWorkflowRun"
)

// ControllerServiceClient is the client API for ControllerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The ControllerService contains endpoints that are exposed on orca controller instances and relate specifically
// to controller operations. This service should only be used by workers.
type ControllerServiceClient interface {
	// Register a Worker
	RegisterWorker(ctx context.Context, in *RegisterWorkerRequest, opts ...grpc.CallOption) (*RegisterWorkerResponse, error)
	// Deregister a Worker
	DeregisterWorker(ctx context.Context, in *DeregisterWorkerRequest, opts ...grpc.CallOption) (*DeregisterWorkerResponse, error)
	// Schedule a Task
	ScheduleTask(ctx context.Context, in *ScheduleTaskRequest, opts ...grpc.CallOption) (*ScheduleTaskResponse, error)
	// Set the status of a Task run.
	SetTaskRunStatus(ctx context.Context, in *SetTaskRunStatusRequest, opts ...grpc.CallOption) (*SetTaskRunStatusResponse, error)
	// Get information on a task run.
	GetTaskRun(ctx context.Context, in *GetTaskRunRequest, opts ...grpc.CallOption) (*GetTaskRunResponse, error)
	// Schedule a Workflow.
	ScheduleWorkflow(ctx context.Context, in *ScheduleWorkflowRequest, opts ...grpc.CallOption) (*ScheduleWorkflowResponse, error)
	// Set the status of a Workflow run.
	SetWorkflowRunStatus(ctx context.Context, in *SetWorkflowRunStatusRequest, opts ...grpc.CallOption) (*SetWorkflowRunStatusResponse, error)
	// Get information on a workflow run.
	GetWorkflowRun(ctx context.Context, in *GetWorkflowRunRequest, opts ...grpc.CallOption) (*GetWorkflowRunResponse, error)
}

type controllerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerServiceClient(cc grpc.ClientConnInterface) ControllerServiceClient {
	return &controllerServiceClient{cc}
}

func (c *controllerServiceClient) RegisterWorker(ctx context.Context, in *RegisterWorkerRequest, opts ...grpc.CallOption) (*RegisterWorkerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterWorkerResponse)
	err := c.cc.Invoke(ctx, ControllerService_RegisterWorker_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) DeregisterWorker(ctx context.Context, in *DeregisterWorkerRequest, opts ...grpc.CallOption) (*DeregisterWorkerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeregisterWorkerResponse)
	err := c.cc.Invoke(ctx, ControllerService_DeregisterWorker_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) ScheduleTask(ctx context.Context, in *ScheduleTaskRequest, opts ...grpc.CallOption) (*ScheduleTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScheduleTaskResponse)
	err := c.cc.Invoke(ctx, ControllerService_ScheduleTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) SetTaskRunStatus(ctx context.Context, in *SetTaskRunStatusRequest, opts ...grpc.CallOption) (*SetTaskRunStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetTaskRunStatusResponse)
	err := c.cc.Invoke(ctx, ControllerService_SetTaskRunStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) GetTaskRun(ctx context.Context, in *GetTaskRunRequest, opts ...grpc.CallOption) (*GetTaskRunResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskRunResponse)
	err := c.cc.Invoke(ctx, ControllerService_GetTaskRun_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) ScheduleWorkflow(ctx context.Context, in *ScheduleWorkflowRequest, opts ...grpc.CallOption) (*ScheduleWorkflowResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScheduleWorkflowResponse)
	err := c.cc.Invoke(ctx, ControllerService_ScheduleWorkflow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) SetWorkflowRunStatus(ctx context.Context, in *SetWorkflowRunStatusRequest, opts ...grpc.CallOption) (*SetWorkflowRunStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetWorkflowRunStatusResponse)
	err := c.cc.Invoke(ctx, ControllerService_SetWorkflowRunStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) GetWorkflowRun(ctx context.Context, in *GetWorkflowRunRequest, opts ...grpc.CallOption) (*GetWorkflowRunResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWorkflowRunResponse)
	err := c.cc.Invoke(ctx, ControllerService_GetWorkflowRun_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControllerServiceServer is the server API for ControllerService service.
// All implementations must embed UnimplementedControllerServiceServer
// for forward compatibility.
//
// The ControllerService contains endpoints that are exposed on orca controller instances and relate specifically
// to controller operations. This service should only be used by workers.
type ControllerServiceServer interface {
	// Register a Worker
	RegisterWorker(context.Context, *RegisterWorkerRequest) (*RegisterWorkerResponse, error)
	// Deregister a Worker
	DeregisterWorker(context.Context, *DeregisterWorkerRequest) (*DeregisterWorkerResponse, error)
	// Schedule a Task
	ScheduleTask(context.Context, *ScheduleTaskRequest) (*ScheduleTaskResponse, error)
	// Set the status of a Task run.
	SetTaskRunStatus(context.Context, *SetTaskRunStatusRequest) (*SetTaskRunStatusResponse, error)
	// Get information on a task run.
	GetTaskRun(context.Context, *GetTaskRunRequest) (*GetTaskRunResponse, error)
	// Schedule a Workflow.
	ScheduleWorkflow(context.Context, *ScheduleWorkflowRequest) (*ScheduleWorkflowResponse, error)
	// Set the status of a Workflow run.
	SetWorkflowRunStatus(context.Context, *SetWorkflowRunStatusRequest) (*SetWorkflowRunStatusResponse, error)
	// Get information on a workflow run.
	GetWorkflowRun(context.Context, *GetWorkflowRunRequest) (*GetWorkflowRunResponse, error)
	mustEmbedUnimplementedControllerServiceServer()
}

// UnimplementedControllerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedControllerServiceServer struct{}

func (UnimplementedControllerServiceServer) RegisterWorker(context.Context, *RegisterWorkerRequest) (*RegisterWorkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWorker not implemented")
}
func (UnimplementedControllerServiceServer) DeregisterWorker(context.Context, *DeregisterWorkerRequest) (*DeregisterWorkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterWorker not implemented")
}
func (UnimplementedControllerServiceServer) ScheduleTask(context.Context, *ScheduleTaskRequest) (*ScheduleTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleTask not implemented")
}
func (UnimplementedControllerServiceServer) SetTaskRunStatus(context.Context, *SetTaskRunStatusRequest) (*SetTaskRunStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTaskRunStatus not implemented")
}
func (UnimplementedControllerServiceServer) GetTaskRun(context.Context, *GetTaskRunRequest) (*GetTaskRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskRun not implemented")
}
func (UnimplementedControllerServiceServer) ScheduleWorkflow(context.Context, *ScheduleWorkflowRequest) (*ScheduleWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleWorkflow not implemented")
}
func (UnimplementedControllerServiceServer) SetWorkflowRunStatus(context.Context, *SetWorkflowRunStatusRequest) (*SetWorkflowRunStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWorkflowRunStatus not implemented")
}
func (UnimplementedControllerServiceServer) GetWorkflowRun(context.Context, *GetWorkflowRunRequest) (*GetWorkflowRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowRun not implemented")
}
func (UnimplementedControllerServiceServer) mustEmbedUnimplementedControllerServiceServer() {}
func (UnimplementedControllerServiceServer) testEmbeddedByValue()                           {}

// UnsafeControllerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControllerServiceServer will
// result in compilation errors.
type UnsafeControllerServiceServer interface {
	mustEmbedUnimplementedControllerServiceServer()
}

func RegisterControllerServiceServer(s grpc.ServiceRegistrar, srv ControllerServiceServer) {
	// If the following call pancis, it indicates UnimplementedControllerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ControllerService_ServiceDesc, srv)
}

func _ControllerService_RegisterWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).RegisterWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_RegisterWorker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).RegisterWorker(ctx, req.(*RegisterWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_DeregisterWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).DeregisterWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_DeregisterWorker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).DeregisterWorker(ctx, req.(*DeregisterWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_ScheduleTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).ScheduleTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_ScheduleTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).ScheduleTask(ctx, req.(*ScheduleTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_SetTaskRunStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTaskRunStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).SetTaskRunStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_SetTaskRunStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).SetTaskRunStatus(ctx, req.(*SetTaskRunStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_GetTaskRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).GetTaskRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_GetTaskRun_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).GetTaskRun(ctx, req.(*GetTaskRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_ScheduleWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).ScheduleWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_ScheduleWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).ScheduleWorkflow(ctx, req.(*ScheduleWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_SetWorkflowRunStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetWorkflowRunStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).SetWorkflowRunStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_SetWorkflowRunStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).SetWorkflowRunStatus(ctx, req.(*SetWorkflowRunStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_GetWorkflowRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).GetWorkflowRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_GetWorkflowRun_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).GetWorkflowRun(ctx, req.(*GetWorkflowRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ControllerService_ServiceDesc is the grpc.ServiceDesc for ControllerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControllerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orca.controller.service.v1.ControllerService",
	HandlerType: (*ControllerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterWorker",
			Handler:    _ControllerService_RegisterWorker_Handler,
		},
		{
			MethodName: "DeregisterWorker",
			Handler:    _ControllerService_DeregisterWorker_Handler,
		},
		{
			MethodName: "ScheduleTask",
			Handler:    _ControllerService_ScheduleTask_Handler,
		},
		{
			MethodName: "SetTaskRunStatus",
			Handler:    _ControllerService_SetTaskRunStatus_Handler,
		},
		{
			MethodName: "GetTaskRun",
			Handler:    _ControllerService_GetTaskRun_Handler,
		},
		{
			MethodName: "ScheduleWorkflow",
			Handler:    _ControllerService_ScheduleWorkflow_Handler,
		},
		{
			MethodName: "SetWorkflowRunStatus",
			Handler:    _ControllerService_SetWorkflowRunStatus_Handler,
		},
		{
			MethodName: "GetWorkflowRun",
			Handler:    _ControllerService_GetWorkflowRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orca/controller/service/v1/service.proto",
}
