// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: orca/workflow/service/v1/service.proto

/*
Package workflowsvcv1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package workflowsvcv1

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = errors.New
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_WorkflowService_Schedule_0(ctx context.Context, marshaler runtime.Marshaler, client WorkflowServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq ScheduleRequest
		metadata runtime.ServerMetadata
		err      error
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	val, ok := pathParams["workflow_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_name")
	}
	protoReq.WorkflowName, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_name", err)
	}
	msg, err := client.Schedule(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_WorkflowService_Schedule_0(ctx context.Context, marshaler runtime.Marshaler, server WorkflowServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq ScheduleRequest
		metadata runtime.ServerMetadata
		err      error
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	val, ok := pathParams["workflow_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_name")
	}
	protoReq.WorkflowName, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_name", err)
	}
	msg, err := server.Schedule(ctx, &protoReq)
	return msg, metadata, err
}

func request_WorkflowService_GetRun_0(ctx context.Context, marshaler runtime.Marshaler, client WorkflowServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq GetRunRequest
		metadata runtime.ServerMetadata
		err      error
	)
	val, ok := pathParams["workflow_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_name")
	}
	protoReq.WorkflowName, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_name", err)
	}
	val, ok = pathParams["workflow_run_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_run_id")
	}
	protoReq.WorkflowRunId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_run_id", err)
	}
	msg, err := client.GetRun(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_WorkflowService_GetRun_0(ctx context.Context, marshaler runtime.Marshaler, server WorkflowServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq GetRunRequest
		metadata runtime.ServerMetadata
		err      error
	)
	val, ok := pathParams["workflow_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_name")
	}
	protoReq.WorkflowName, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_name", err)
	}
	val, ok = pathParams["workflow_run_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_run_id")
	}
	protoReq.WorkflowRunId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_run_id", err)
	}
	msg, err := server.GetRun(ctx, &protoReq)
	return msg, metadata, err
}

func request_WorkflowService_DescribeRun_0(ctx context.Context, marshaler runtime.Marshaler, client WorkflowServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq DescribeRunRequest
		metadata runtime.ServerMetadata
		err      error
	)
	val, ok := pathParams["workflow_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_name")
	}
	protoReq.WorkflowName, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_name", err)
	}
	val, ok = pathParams["workflow_run_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_run_id")
	}
	protoReq.WorkflowRunId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_run_id", err)
	}
	msg, err := client.DescribeRun(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_WorkflowService_DescribeRun_0(ctx context.Context, marshaler runtime.Marshaler, server WorkflowServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq DescribeRunRequest
		metadata runtime.ServerMetadata
		err      error
	)
	val, ok := pathParams["workflow_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_name")
	}
	protoReq.WorkflowName, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_name", err)
	}
	val, ok = pathParams["workflow_run_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_run_id")
	}
	protoReq.WorkflowRunId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_run_id", err)
	}
	msg, err := server.DescribeRun(ctx, &protoReq)
	return msg, metadata, err
}

func request_WorkflowService_CancelRun_0(ctx context.Context, marshaler runtime.Marshaler, client WorkflowServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq CancelRunRequest
		metadata runtime.ServerMetadata
		err      error
	)
	val, ok := pathParams["workflow_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_name")
	}
	protoReq.WorkflowName, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_name", err)
	}
	val, ok = pathParams["workflow_run_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_run_id")
	}
	protoReq.WorkflowRunId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_run_id", err)
	}
	msg, err := client.CancelRun(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_WorkflowService_CancelRun_0(ctx context.Context, marshaler runtime.Marshaler, server WorkflowServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq CancelRunRequest
		metadata runtime.ServerMetadata
		err      error
	)
	val, ok := pathParams["workflow_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_name")
	}
	protoReq.WorkflowName, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_name", err)
	}
	val, ok = pathParams["workflow_run_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_run_id")
	}
	protoReq.WorkflowRunId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_run_id", err)
	}
	msg, err := server.CancelRun(ctx, &protoReq)
	return msg, metadata, err
}

// RegisterWorkflowServiceHandlerServer registers the http handlers for service WorkflowService to "mux".
// UnaryRPC     :call WorkflowServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterWorkflowServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterWorkflowServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server WorkflowServiceServer) error {
	mux.Handle(http.MethodPost, pattern_WorkflowService_Schedule_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/orca.workflow.service.v1.WorkflowService/Schedule", runtime.WithHTTPPathPattern("/v1/workflow/{workflow_name}/run"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_WorkflowService_Schedule_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_WorkflowService_Schedule_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_WorkflowService_GetRun_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/orca.workflow.service.v1.WorkflowService/GetRun", runtime.WithHTTPPathPattern("/v1/workflow/{workflow_name}/run/{workflow_run_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_WorkflowService_GetRun_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_WorkflowService_GetRun_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_WorkflowService_DescribeRun_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/orca.workflow.service.v1.WorkflowService/DescribeRun", runtime.WithHTTPPathPattern("/v1/workflow/{workflow_name}/run/{workflow_run_id}/describe"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_WorkflowService_DescribeRun_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_WorkflowService_DescribeRun_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodDelete, pattern_WorkflowService_CancelRun_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/orca.workflow.service.v1.WorkflowService/CancelRun", runtime.WithHTTPPathPattern("/v1/workflow/{workflow_name}/run/{workflow_run_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_WorkflowService_CancelRun_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_WorkflowService_CancelRun_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

// RegisterWorkflowServiceHandlerFromEndpoint is same as RegisterWorkflowServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterWorkflowServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterWorkflowServiceHandler(ctx, mux, conn)
}

// RegisterWorkflowServiceHandler registers the http handlers for service WorkflowService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterWorkflowServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterWorkflowServiceHandlerClient(ctx, mux, NewWorkflowServiceClient(conn))
}

// RegisterWorkflowServiceHandlerClient registers the http handlers for service WorkflowService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "WorkflowServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "WorkflowServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "WorkflowServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterWorkflowServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client WorkflowServiceClient) error {
	mux.Handle(http.MethodPost, pattern_WorkflowService_Schedule_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/orca.workflow.service.v1.WorkflowService/Schedule", runtime.WithHTTPPathPattern("/v1/workflow/{workflow_name}/run"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_WorkflowService_Schedule_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_WorkflowService_Schedule_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_WorkflowService_GetRun_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/orca.workflow.service.v1.WorkflowService/GetRun", runtime.WithHTTPPathPattern("/v1/workflow/{workflow_name}/run/{workflow_run_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_WorkflowService_GetRun_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_WorkflowService_GetRun_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_WorkflowService_DescribeRun_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/orca.workflow.service.v1.WorkflowService/DescribeRun", runtime.WithHTTPPathPattern("/v1/workflow/{workflow_name}/run/{workflow_run_id}/describe"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_WorkflowService_DescribeRun_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_WorkflowService_DescribeRun_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodDelete, pattern_WorkflowService_CancelRun_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/orca.workflow.service.v1.WorkflowService/CancelRun", runtime.WithHTTPPathPattern("/v1/workflow/{workflow_name}/run/{workflow_run_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_WorkflowService_CancelRun_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_WorkflowService_CancelRun_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_WorkflowService_Schedule_0    = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3}, []string{"v1", "workflow", "workflow_name", "run"}, ""))
	pattern_WorkflowService_GetRun_0      = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"v1", "workflow", "workflow_name", "run", "workflow_run_id"}, ""))
	pattern_WorkflowService_DescribeRun_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3, 1, 0, 4, 1, 5, 4, 2, 5}, []string{"v1", "workflow", "workflow_name", "run", "workflow_run_id", "describe"}, ""))
	pattern_WorkflowService_CancelRun_0   = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"v1", "workflow", "workflow_name", "run", "workflow_run_id"}, ""))
)

var (
	forward_WorkflowService_Schedule_0    = runtime.ForwardResponseMessage
	forward_WorkflowService_GetRun_0      = runtime.ForwardResponseMessage
	forward_WorkflowService_DescribeRun_0 = runtime.ForwardResponseMessage
	forward_WorkflowService_CancelRun_0   = runtime.ForwardResponseMessage
)
