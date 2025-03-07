package task

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/davidsbond/orca/internal/daemon/controller/database"
	taskv1 "github.com/davidsbond/orca/internal/proto/orca/task/v1"
)

type (
	Run struct {
		ID            string
		WorkflowRunID string
		TaskName      string
		CreatedAt     time.Time
		ScheduledAt   sql.Null[time.Time]
		StartedAt     sql.Null[time.Time]
		CompletedAt   sql.Null[time.Time]
		Status        Status
		Input         sql.Null[pgtype.JSONB]
		Output        sql.Null[pgtype.JSONB]
		WorkerID      sql.Null[string]
		IdempotentKey sql.Null[string]
	}

	Status int

	PostgresRepository struct {
		db *pgxpool.Pool
	}
)

const (
	StatusUnspecified Status = iota
	StatusPending
	StatusScheduled
	StatusRunning
	StatusComplete
	StatusFailed
	StatusSkipped
	StatusTimeout
)

var (
	ErrNotFound            = errors.New("not found")
	ErrWorkflowRunNotFound = errors.New("workflow run does not exist")
	ErrWorkerNotFound      = errors.New("worker does not exist")
)

func (r Run) ToProto() *taskv1.Run {
	run := &taskv1.Run{
		RunId:         r.ID,
		WorkflowRunId: r.WorkflowRunID,
		TaskName:      r.TaskName,
		CreatedAt:     timestamppb.New(r.CreatedAt),
		Status:        taskv1.Status(r.Status),
		Input:         r.Input.V.Bytes,
		Output:        r.Output.V.Bytes,
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

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (pr *PostgresRepository) Save(ctx context.Context, run Run) error {
	return database.Write(ctx, pr.db, func(ctx context.Context, tx pgx.Tx) error {
		const q = `
			INSERT INTO task_run (
			  id, workflow_run_id, task_name, 
			  created_at, scheduled_at, started_at,
			  completed_at, status, input, output,
			  worker_id, idempotent_key
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
			ON CONFLICT (id) DO UPDATE SET
				scheduled_at = EXCLUDED.scheduled_at,
				started_at = EXCLUDED.started_at,
				completed_at = EXCLUDED.completed_at,
				status = EXCLUDED.status,
				input = EXCLUDED.input,
				output = EXCLUDED.output,
				worker_id = EXCLUDED.worker_id
		`

		_, err := tx.Exec(ctx, q,
			run.ID,
			run.WorkflowRunID,
			run.TaskName,
			run.CreatedAt,
			run.ScheduledAt,
			run.StartedAt,
			run.CompletedAt,
			run.Status,
			run.Input,
			run.Output,
			run.WorkerID,
			run.IdempotentKey,
		)
		switch {
		case database.IsForeignKeyViolation(err, "workflow_run_id"):
			return ErrWorkflowRunNotFound
		case database.IsForeignKeyViolation(err, "worker_id"):
			return ErrWorkerNotFound
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
			SELECT 
			    id, workflow_run_id, task_name, 
			    created_at, scheduled_at, started_at, 
			    completed_at, status, input, output, 
			    worker_id, idempotent_key
			FROM task_run WHERE id = $1
		`

		var run Run
		err := tx.QueryRow(ctx, q, id).Scan(
			&run.ID,
			&run.WorkflowRunID,
			&run.TaskName,
			&run.CreatedAt,
			&run.ScheduledAt,
			&run.StartedAt,
			&run.CompletedAt,
			&run.Status,
			&run.Input,
			&run.Output,
			&run.WorkerID,
			&run.IdempotentKey,
		)
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return run, ErrNotFound
		case err != nil:
			return run, err
		default:
			return run, nil
		}
	})
}

func (pr *PostgresRepository) GetPendingTaskRuns(ctx context.Context) ([]Run, error) {
	return database.Read(ctx, pr.db, func(ctx context.Context, tx pgx.Tx) ([]Run, error) {
		const q = `
			SELECT 
			    id, workflow_run_id, task_name, 
			    created_at, scheduled_at, 
			    started_at, completed_at, 
			    status, input, output,
			    worker_id, idempotent_key
			FROM task_run WHERE (status = $1 AND worker_id IS NULL)
		`

		rows, err := tx.Query(ctx, q, StatusPending)
		if err != nil {
			return nil, err
		}

		return database.ScanAll[Run](ctx, rows, func(run *Run) []any {
			return []any{
				&run.ID,
				&run.WorkflowRunID,
				&run.TaskName,
				&run.CreatedAt,
				&run.ScheduledAt,
				&run.StartedAt,
				&run.CompletedAt,
				&run.Status,
				&run.Input,
				&run.Output,
				&run.WorkerID,
				&run.IdempotentKey,
			}
		})
	})
}

func (pr *PostgresRepository) ListForWorkflowRun(ctx context.Context, runID string) ([]Run, error) {
	return database.Read(ctx, pr.db, func(ctx context.Context, tx pgx.Tx) ([]Run, error) {
		const q = `
			SELECT 
			    id, workflow_run_id, task_name, 
			    created_at, scheduled_at, 
			    started_at, completed_at, 
			    status, input, output,
			    worker_id, idempotent_key
			FROM task_run WHERE workflow_run_id = $1
		`

		rows, err := tx.Query(ctx, q, runID)
		if err != nil {
			return nil, err
		}

		return database.ScanAll[Run](ctx, rows, func(run *Run) []any {
			return []any{
				&run.ID,
				&run.WorkflowRunID,
				&run.TaskName,
				&run.CreatedAt,
				&run.ScheduledAt,
				&run.StartedAt,
				&run.CompletedAt,
				&run.Status,
				&run.Input,
				&run.Output,
				&run.WorkerID,
				&run.IdempotentKey,
			}
		})
	})
}

func (pr *PostgresRepository) FindByIdempotentKey(ctx context.Context, name string, key string) (Run, error) {
	return database.Read(ctx, pr.db, func(ctx context.Context, tx pgx.Tx) (Run, error) {
		const q = `
			SELECT 
			    id, workflow_run_id, task_name, 
			    created_at, scheduled_at, started_at, 
			    completed_at, status, input, output, 
			    worker_id, idempotent_key
			FROM task_run 
			WHERE task_name = $1 AND idempotent_key = $2 AND status = $3
			LIMIT 1
		`

		var run Run
		err := tx.QueryRow(ctx, q, name, key, StatusComplete).Scan(
			&run.ID,
			&run.WorkflowRunID,
			&run.TaskName,
			&run.CreatedAt,
			&run.ScheduledAt,
			&run.StartedAt,
			&run.CompletedAt,
			&run.Status,
			&run.Input,
			&run.Output,
			&run.WorkerID,
			&run.IdempotentKey,
		)
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return run, ErrNotFound
		case err != nil:
			return run, err
		default:
			return run, nil
		}
	})
}
