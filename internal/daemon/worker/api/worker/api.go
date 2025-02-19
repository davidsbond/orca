package worker

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	workersvcv1 "github.com/davidsbond/orca/internal/proto/orca/private/worker/service/v1"
)

type (
	API struct {
		workersvcv1.UnimplementedWorkerServiceServer

		service Service
	}

	Service interface {
		ScheduleWorkflow(ctx context.Context, name string, runID string, input json.RawMessage) error
		ScheduleTask(ctx context.Context, name string, runID string, input json.RawMessage) error
		CancelWorkflowRun(ctx context.Context, runID string) error
	}
)

func NewAPI(service Service) *API {
	return &API{
		service: service,
	}
}

func (api *API) RegisterGRPC(r grpc.ServiceRegistrar) {
	workersvcv1.RegisterWorkerServiceServer(r, api)
}

func (api *API) RegisterHTTP(_ context.Context, _ *runtime.ServeMux) {
	return
}

func (api *API) RunWorkflow(ctx context.Context, request *workersvcv1.RunWorkflowRequest) (*workersvcv1.RunWorkflowResponse, error) {
	err := api.service.ScheduleWorkflow(ctx, request.GetWorkflowName(), request.GetWorkflowRunId(), request.GetInput())
	switch {
	case errors.Is(err, ErrNotFound):
		return nil, status.Errorf(codes.NotFound, "workflow %s is not registered on this worker", request.GetWorkflowName())
	case errors.Is(err, context.Canceled):
		return nil, status.Errorf(codes.Canceled, "recieved cancellation while scheduling workflow %s", request.GetWorkflowName())
	case errors.Is(err, context.DeadlineExceeded):
		return nil, status.Errorf(codes.DeadlineExceeded, "failed to schedule workflow %s within deadline", request.GetWorkflowName())
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to schedule workflow %s: %v", request.GetWorkflowName(), err)
	default:
		return &workersvcv1.RunWorkflowResponse{}, nil
	}
}

func (api *API) RunTask(ctx context.Context, request *workersvcv1.RunTaskRequest) (*workersvcv1.RunTaskResponse, error) {
	err := api.service.ScheduleTask(ctx, request.GetTaskName(), request.GetTaskRunId(), request.GetInput())
	switch {
	case errors.Is(err, ErrNotFound):
		return nil, status.Errorf(codes.NotFound, "task %s is not registered on this worker", request.GetTaskName())
	case errors.Is(err, context.Canceled):
		return nil, status.Errorf(codes.Canceled, "recieved cancellation while scheduling task %s", request.GetTaskName())
	case errors.Is(err, context.DeadlineExceeded):
		return nil, status.Errorf(codes.DeadlineExceeded, "failed to schedule task %s within deadline", request.GetTaskName())
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to schedule task %s: %v", request.GetTaskName(), err)
	default:
		return &workersvcv1.RunTaskResponse{}, nil
	}
}

func (api *API) CancelWorkflowRun(ctx context.Context, request *workersvcv1.CancelWorkflowRunRequest) (*workersvcv1.CancelWorkflowRunResponse, error) {
	if err := api.service.CancelWorkflowRun(ctx, request.GetWorkflowRunId()); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to cancel workflow run: %v", err)
	}

	return &workersvcv1.CancelWorkflowRunResponse{}, nil
}
