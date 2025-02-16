package workflow

import (
	"context"
	"encoding/json"
	"errors"
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
		Get(ctx context.Context, id string) (workflow.Run, error)
		Save(ctx context.Context, run workflow.Run) error
	}
)

var (
	ErrNotFound = errors.New("not found")
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

func (h *Handler) GetRun(ctx context.Context, runID string) (workflow.Run, error) {
	run, err := h.workflows.Get(ctx, runID)
	switch {
	case errors.Is(err, workflow.ErrNotFound):
		return workflow.Run{}, ErrNotFound
	case err != nil:
		return workflow.Run{}, err
	default:
		return run, nil
	}
}
