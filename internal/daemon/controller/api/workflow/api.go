package workflow

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	workflowsvcv1 "github.com/davidsbond/orca/internal/proto/orca/workflow/service/v1"
)

type (
	API struct {
		workflowsvcv1.UnimplementedWorkflowServiceServer

		service Service
	}

	Service interface {
		ScheduleWorkflow(ctx context.Context, name string, input json.RawMessage) (string, error)
	}
)

func NewAPI(service Service) *API {
	return &API{service: service}
}

func (api *API) Register(r grpc.ServiceRegistrar) {
	workflowsvcv1.RegisterWorkflowServiceServer(r, api)
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
	panic("implement me")
}
