package task

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
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
		Input         json.RawMessage
		Output        json.RawMessage
	}

	Status int

	PostgresRepository struct {
		db *pgx.Conn
	}
)

const (
	StatusUnspecified Status = iota
	StatusPending
	StatusScheduled
	StatusRunning
	StatusComplete
	StatusFailed
)

var (
	ErrNotFound            = errors.New("not found")
	ErrWorkflowRunNotFound = errors.New("workflow run does not exist")
)

func (r Run) ToProto() *taskv1.Run {
	run := &taskv1.Run{
		RunId:         r.ID,
		WorkflowRunId: r.WorkflowRunID,
		TaskName:      r.TaskName,
		CreatedAt:     timestamppb.New(r.CreatedAt),
		Status:        taskv1.Status(r.Status),
		Input:         r.Input,
		Output:        r.Output,
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
			INSERT INTO task_run (id, workflow_run_id, task_name, created_at, scheduled_at, started_at, completed_at, status, input, output)
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
			run.WorkflowRunID,
			run.TaskName,
			run.CreatedAt,
			run.ScheduledAt,
			run.StartedAt,
			run.CompletedAt,
			run.Status,
			run.Input,
			run.Output,
		)
		switch {
		case database.IsForeignKeyViolation(err, "workflow_run_id"):
			return ErrWorkflowRunNotFound
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
			SELECT id, workflow_run_id, task_name, created_at, scheduled_at, started_at, completed_at, status, input, output 
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
