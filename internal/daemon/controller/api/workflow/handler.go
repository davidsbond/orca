package workflow

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"

	"github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
)

type (
	Handler struct {
		workflows Repository
	}

	Repository interface {
		Save(ctx context.Context, run workflow.Run) error
	}
)

func NewHandler(workflows Repository) *Handler {
	return &Handler{
		workflows: workflows,
	}
}

func (h *Handler) ScheduleWorkflow(ctx context.Context, name string, input json.RawMessage) (string, error) {
	run := workflow.Run{
		ID:           uuid.NewString(),
		WorkflowName: name,
		CreatedAt:    time.Now(),
		Status:       workflow.StatusPending,
	}

	if len(input) > 0 {
		run.Input.Valid = true
		run.Input.V = pgtype.JSONB{
			Bytes:  input,
			Status: pgtype.Present,
		}
	}

	if err := h.workflows.Save(ctx, run); err != nil {
		return "", err
	}

	return run.ID, nil
}
