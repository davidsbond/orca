package workflow

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/davidsbond/orca/internal/daemon/controller/database"
	workflowv1 "github.com/davidsbond/orca/internal/proto/orca/workflow/v1"
)

type (
	Run struct {
		ID                  string
		ParentWorkflowRunID sql.Null[string]
		WorkflowName        string
		CreatedAt           time.Time
		ScheduledAt         sql.Null[time.Time]
		StartedAt           sql.Null[time.Time]
		CompletedAt         sql.Null[time.Time]
		Status              Status
		Input               json.RawMessage
		Output              json.RawMessage
	}

	Status int

	PostgresRepository struct {
		db *pgx.Conn
	}
)

var (
	ErrNotFound               = errors.New("not found")
	ErrParentWorkflowNotFound = errors.New("parent workflow not found")
)

const (
	StatusUnspecified Status = iota
	StatusPending
	StatusScheduled
	StatusRunning
	StatusComplete
	StatusFailed
)

func (r Run) ToProto() *workflowv1.Run {
	run := &workflowv1.Run{
		RunId:               r.ID,
		ParentWorkflowRunId: r.ParentWorkflowRunID.V,
		WorkflowName:        r.WorkflowName,
		CreatedAt:           timestamppb.New(r.CreatedAt),
		Status:              workflowv1.Status(r.Status),
		Input:               r.Input,
		Output:              r.Output,
	}

	if r.ScheduledAt.Valid {
		run.ScheduledAt = timestamppb.New(r.ScheduledAt.V)
	}

	if r.StartedAt.Valid {
		run.StartedAt = timestamppb.New(r.StartedAt.V)
	}

	if r.CompletedAt.Valid {
		run.CompletedAt = timestamppb.New(r.CompletedAt.V)
	}

	return run
}

func NewPostgresRepository(db *pgx.Conn) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (pr *PostgresRepository) Save(ctx context.Context, run Run) error {
	return database.Write(ctx, pr.db, func(ctx context.Context, tx pgx.Tx) error {
		const q = `
			INSERT INTO workflow_run (id, parent_workflow_run_id, workflow_name, created_at, scheduled_at, started_at, completed_at, status, input, output)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			ON CONFLICT (id) DO UPDATE SET
				scheduled_at = EXCLUDED.scheduled_at,
				started_at = EXCLUDED.started_at,
				completed_at = EXCLUDED.completed_at,
				status = EXCLUDED.status,
				input = EXCLUDED.input,
				output = EXCLUDED.output	
		`

		_, err := tx.Exec(ctx, q,
			run.ID,
			run.ParentWorkflowRunID,
			run.WorkflowName,
			run.CreatedAt,
			run.ScheduledAt,
			run.StartedAt,
			run.CompletedAt,
			run.Status,
			run.Input,
			run.Output,
		)
		switch {
		case database.IsForeignKeyViolation(err, "parent_workflow_run_id"):
			return ErrParentWorkflowNotFound
		case err != nil:
			return err
		default:
			return nil
		}
	})
}

func (pr *PostgresRepository) Get(ctx context.Context, id string) (Run, error) {
	return database.Read(ctx, pr.db, func(ctx context.Context, tx pgx.Tx) (Run, error) {
		const q = `
			SELECT id, parent_workflow_run_id, workflow_name, created_at, scheduled_at, started_at, completed_at, status, input, output 
			FROM workflow_run WHERE id = $1
		`

		var run Run
		err := tx.QueryRow(ctx, q, id).Scan(
			&run.ID,
			&run.ParentWorkflowRunID,
			&run.WorkflowName,
			&run.CreatedAt,
			&run.ScheduledAt,
			&run.StartedAt,
			&run.CompletedAt,
			&run.Status,
			&run.Input,
			&run.Output,
		)
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return run, ErrNotFound
		case err != nil:
			return run, err
		default:
			return run, nil
		}
	})
}
