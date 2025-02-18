package controller

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/davidsbond/orca/internal/daemon/controller/database/task"
	"github.com/davidsbond/orca/internal/daemon/controller/database/worker"
	"github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
	controllersvcv1 "github.com/davidsbond/orca/internal/proto/orca/private/controller/service/v1"
)

type (
	API struct {
		controllersvcv1.UnimplementedControllerServiceServer

		service Service
	}

	Service interface {
		RegisterWorker(ctx context.Context, w worker.Worker) error
		DeregisterWorker(ctx context.Context, id string) error
		ScheduleTask(ctx context.Context, params ScheduleTaskParams) (string, error)
		SetTaskRunStatus(ctx context.Context, id string, status task.Status, output json.RawMessage) error
		GetTaskRun(ctx context.Context, id string) (task.Run, error)
		ScheduleWorkflow(ctx context.Context, params ScheduleWorkflowParams) (string, error)
		SetWorkflowRunStatus(ctx context.Context, id string, status workflow.Status, output json.RawMessage) error
		GetWorkflowRun(ctx context.Context, id string) (workflow.Run, error)
	}
)

func NewAPI(service Service) *API {
	return &API{
		service: service,
	}
}

func (api *API) RegisterGRPC(r grpc.ServiceRegistrar) {
	controllersvcv1.RegisterControllerServiceServer(r, api)
}

func (api *API) RegisterHTTP(_ context.Context, _ *runtime.ServeMux) {
	return
}

func (api *API) RegisterWorker(ctx context.Context, request *controllersvcv1.RegisterWorkerRequest) (*controllersvcv1.RegisterWorkerResponse, error) {
	w := worker.Worker{
		ID:               request.GetWorker().GetId(),
		AdvertiseAddress: request.GetWorker().GetAdvertiseAddress(),
		Workflows:        request.GetWorker().GetWorkflows(),
		Tasks:            request.GetWorker().GetTasks(),
	}

	err := api.service.RegisterWorker(ctx, w)
	switch {
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to register worker: %v", err)
	default:
		return &controllersvcv1.RegisterWorkerResponse{}, nil
	}
}

func (api *API) DeregisterWorker(ctx context.Context, request *controllersvcv1.DeregisterWorkerRequest) (*controllersvcv1.DeregisterWorkerResponse, error) {
	err := api.service.DeregisterWorker(ctx, request.GetWorkerId())
	switch {
	case errors.Is(err, ErrWorkerNotFound):
		return nil, status.Errorf(codes.NotFound, "worker %s is not registered", request.GetWorkerId())
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to deregister worker: %v", err)
	default:
		return &controllersvcv1.DeregisterWorkerResponse{}, nil
	}
}

func (api *API) ScheduleTask(ctx context.Context, request *controllersvcv1.ScheduleTaskRequest) (*controllersvcv1.ScheduleTaskResponse, error) {
	runID, err := api.service.ScheduleTask(ctx, ScheduleTaskParams{
		WorkflowRunID: request.GetWorkflowRunId(),
		TaskName:      request.GetTaskName(),
		Input:         request.GetInput(),
		IdempotentKey: request.GetIdempotentKey(),
	})
	switch {
	case errors.Is(err, ErrWorkflowRunNotFound):
		return nil, status.Errorf(codes.NotFound, "workflow run %s does not exist", request.GetWorkflowRunId())
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to schedule task: %v", err)
	default:
		return &controllersvcv1.ScheduleTaskResponse{TaskRunId: runID}, nil
	}
}

func (api *API) SetTaskRunStatus(ctx context.Context, request *controllersvcv1.SetTaskRunStatusRequest) (*controllersvcv1.SetTaskRunStatusResponse, error) {
	err := api.service.SetTaskRunStatus(ctx, request.GetTaskRunId(), task.Status(request.GetStatus()), request.GetOutput())
	switch {
	case errors.Is(err, ErrWorkflowRunNotFound):
		return nil, status.Error(codes.NotFound, "workflow run does not exist")
	case errors.Is(err, ErrTaskRunNotFound):
		return nil, status.Errorf(codes.NotFound, "task run %s does not exist", request.GetTaskRunId())
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to set task status: %v", err)
	default:
		return &controllersvcv1.SetTaskRunStatusResponse{}, nil
	}
}

func (api *API) GetTaskRun(ctx context.Context, request *controllersvcv1.GetTaskRunRequest) (*controllersvcv1.GetTaskRunResponse, error) {
	run, err := api.service.GetTaskRun(ctx, request.GetTaskRunId())
	switch {
	case errors.Is(err, ErrTaskRunNotFound):
		return nil, status.Errorf(codes.NotFound, "task run %s does not exist", request.GetTaskRunId())
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to get task run: %v", err)
	}

	return &controllersvcv1.GetTaskRunResponse{TaskRun: run.ToProto()}, nil
}

func (api *API) ScheduleWorkflow(ctx context.Context, request *controllersvcv1.ScheduleWorkflowRequest) (*controllersvcv1.ScheduleWorkflowResponse, error) {
	runID, err := api.service.ScheduleWorkflow(ctx, ScheduleWorkflowParams{
		ParentWorkflowRunID: request.GetParentWorkflowRunId(),
		WorkflowName:        request.GetWorkflowName(),
		Input:               request.GetInput(),
		IdempotentKey:       request.GetIdempotentKey(),
	})
	switch {
	case errors.Is(err, ErrParentWorkflowNotFound):
		return nil, status.Errorf(codes.NotFound, "parent workflow run %s does not exist", request.GetParentWorkflowRunId())
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to schedule workflow: %v", err)
	default:
		return &controllersvcv1.ScheduleWorkflowResponse{WorkflowRunId: runID}, nil
	}
}

func (api *API) SetWorkflowRunStatus(ctx context.Context, request *controllersvcv1.SetWorkflowRunStatusRequest) (*controllersvcv1.SetWorkflowRunStatusResponse, error) {
	err := api.service.SetWorkflowRunStatus(ctx, request.GetWorkflowRunId(), workflow.Status(request.GetStatus()), request.GetOutput())
	switch {
	case errors.Is(err, ErrWorkflowRunNotFound):
		return nil, status.Errorf(codes.NotFound, "workflow run %s does not exist", request.GetWorkflowRunId())
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to set task status: %v", err)
	default:
		return &controllersvcv1.SetWorkflowRunStatusResponse{}, nil
	}
}

func (api *API) GetWorkflowRun(ctx context.Context, request *controllersvcv1.GetWorkflowRunRequest) (*controllersvcv1.GetWorkflowRunResponse, error) {
	run, err := api.service.GetWorkflowRun(ctx, request.GetWorkflowRunId())
	switch {
	case errors.Is(err, ErrWorkflowRunNotFound):
		return nil, status.Errorf(codes.NotFound, "workflow run %s does not exist", request.GetWorkflowRunId())
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to get workflow run: %v", err)
	}

	return &controllersvcv1.GetWorkflowRunResponse{WorkflowRun: run.ToProto()}, nil
}
