package workflow

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
	workflowsvcv1 "github.com/davidsbond/orca/internal/proto/orca/workflow/service/v1"
)

type (
	API struct {
		workflowsvcv1.UnimplementedWorkflowServiceServer

		service Service
	}

	Service interface {
		ScheduleWorkflow(ctx context.Context, name string, input json.RawMessage) (string, error)
		GetRun(ctx context.Context, runID string) (workflow.Run, error)
		DescribeRun(ctx context.Context, runID string) (Description, error)
		CancelRun(ctx context.Context, runID string) error
	}
)

func NewAPI(service Service) *API {
	return &API{service: service}
}

func (api *API) RegisterGRPC(r grpc.ServiceRegistrar) {
	workflowsvcv1.RegisterWorkflowServiceServer(r, api)
}

func (api *API) RegisterHTTP(ctx context.Context, r *runtime.ServeMux) {
	if err := workflowsvcv1.RegisterWorkflowServiceHandlerServer(ctx, r, api); err != nil {
		panic(err)
	}
}

func (api *API) Schedule(ctx context.Context, request *workflowsvcv1.ScheduleRequest) (*workflowsvcv1.ScheduleResponse, error) {
	runID, err := api.service.ScheduleWorkflow(ctx, request.GetWorkflowName(), request.GetInput())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "faied to schedule workflow %s: %v", request.GetWorkflowName(), err)
	}

	return &workflowsvcv1.ScheduleResponse{
		WorkflowRunId: runID,
	}, nil
}

func (api *API) GetRun(ctx context.Context, request *workflowsvcv1.GetRunRequest) (*workflowsvcv1.GetRunResponse, error) {
	run, err := api.service.GetRun(ctx, request.GetWorkflowRunId())
	switch {
	case errors.Is(err, ErrNotFound):
		return nil, status.Errorf(codes.NotFound, "workflow run %s does not exist", request.GetWorkflowRunId())
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to get workflow run %s: %v", request.GetWorkflowRunId(), err)
	default:
		return &workflowsvcv1.GetRunResponse{WorkflowRun: run.ToProto()}, nil
	}
}

func (api *API) DescribeRun(ctx context.Context, request *workflowsvcv1.DescribeRunRequest) (*workflowsvcv1.DescribeRunResponse, error) {
	description, err := api.service.DescribeRun(ctx, request.GetWorkflowRunId())
	switch {
	case errors.Is(err, ErrNotFound):
		return nil, status.Errorf(codes.NotFound, "workflow run %s does not exist", request.GetWorkflowRunId())
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to describe workflow run %s: %v", request.GetWorkflowRunId(), err)
	default:
		return &workflowsvcv1.DescribeRunResponse{Description: description.ToProto()}, nil
	}
}

func (api *API) CancelRun(ctx context.Context, request *workflowsvcv1.CancelRunRequest) (*workflowsvcv1.CancelRunResponse, error) {
	err := api.service.CancelRun(ctx, request.GetWorkflowRunId())
	switch {
	case errors.Is(err, ErrNotFound):
		return nil, status.Errorf(codes.NotFound, "workflow run %s does not exist", request.GetWorkflowRunId())
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to cancel workflow run %s: %v", request.GetWorkflowRunId(), err)
	default:
		return &workflowsvcv1.CancelRunResponse{}, nil
	}
}
