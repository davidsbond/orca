package api

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	workersvcv1 "github.com/davidsbond/orca/internal/proto/orca/worker/service/v1"
	"github.com/davidsbond/orca/pkg/task"
	"github.com/davidsbond/orca/pkg/workflow"
)

type (
	API struct {
		workersvcv1.UnimplementedWorkerServiceServer

		registry  Registry
		scheduler Scheduler
	}

	Registry interface {
		GetWorkflow(name string) (workflow.Workflow, bool)
		GetTask(name string) (task.Task, bool)
	}

	Scheduler interface {
		ScheduleWorkflow(ctx context.Context, runID string, workflow workflow.Workflow, input []byte) error
		ScheduleTask(ctx context.Context, runID string, task task.Task, input []byte) error
	}
)

func NewAPI(registry Registry, scheduler Scheduler) *API {
	return &API{
		registry:  registry,
		scheduler: scheduler,
	}
}

func (api *API) RunWorkflow(ctx context.Context, request *workersvcv1.RunWorkflowRequest) (*workersvcv1.RunWorkflowResponse, error) {
	wf, ok := api.registry.GetWorkflow(request.GetWorkflowName())
	if !ok {
		return nil, status.Errorf(codes.NotFound, "workflow %s is not registered on this worker", request.GetWorkflowName())
	}

	err := api.scheduler.ScheduleWorkflow(ctx, request.GetWorkflowRunId(), wf, request.GetParameters())
	switch {
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
	tsk, ok := api.registry.GetTask(request.GetTaskName())
	if !ok {
		return nil, status.Errorf(codes.NotFound, "task %s is not registered on this worker", request.GetTaskName())
	}

	err := api.scheduler.ScheduleTask(ctx, request.GetTaskRunId(), tsk, request.GetParameters())
	switch {
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
