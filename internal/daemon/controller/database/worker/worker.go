package worker

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/davidsbond/orca/internal/daemon/controller/database"
)

type (
	Worker struct {
		ID               string
		AdvertiseAddress string
		Workflows        []string
		Tasks            []string
	}

	PostgresRepository struct {
		db *pgxpool.Pool
	}
)

var (
	ErrNotFound = errors.New("not found")
)

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (pr *PostgresRepository) Save(ctx context.Context, worker Worker) error {
	return database.Write(ctx, pr.db, func(ctx context.Context, tx pgx.Tx) error {
		const q = `
			INSERT INTO worker (id, advertise_address, workflows, tasks) 
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (id) DO UPDATE SET
				advertise_address = EXCLUDED.advertise_address,
				workflows = EXCLUDED.workflows,
				tasks = EXCLUDED.tasks
		`

		_, err := tx.Exec(ctx, q,
			worker.ID,
			worker.AdvertiseAddress,
			pgtype.FlatArray[string](worker.Workflows),
			pgtype.FlatArray[string](worker.Tasks),
		)

		return err
	})
}

func (pr *PostgresRepository) Delete(ctx context.Context, id string) error {
	return database.Write(ctx, pr.db, func(ctx context.Context, tx pgx.Tx) error {
		const q = `DELETE FROM worker WHERE id = $1`

		result, err := tx.Exec(ctx, q, id)
		if err != nil {
			return err
		}

		if result.RowsAffected() == 0 {
			return ErrNotFound
		}

		return nil
	})
}

func (pr *PostgresRepository) GetWorkerForTask(ctx context.Context, name string) (Worker, error) {
	return database.Read(ctx, pr.db, func(ctx context.Context, tx pgx.Tx) (Worker, error) {
		const q = `
			SELECT id, advertise_address, workflows, tasks 
			FROM worker WHERE $1 = ANY(tasks) 
			ORDER BY RANDOM() LIMIT 1
		`

		var (
			worker    Worker
			workflows pgtype.FlatArray[string]
			tasks     pgtype.FlatArray[string]
		)

		err := tx.QueryRow(ctx, q, name).Scan(
			&worker.ID,
			&worker.AdvertiseAddress,
			&workflows,
			&tasks,
		)
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return worker, ErrNotFound
		case err != nil:
			return worker, err
		default:
			worker.Workflows = workflows
			worker.Tasks = tasks

			return worker, nil
		}
	})
}

func (pr *PostgresRepository) GetWorkerForWorkflow(ctx context.Context, name string) (Worker, error) {
	return database.Read(ctx, pr.db, func(ctx context.Context, tx pgx.Tx) (Worker, error) {
		const q = `
			SELECT id, advertise_address, workflows, tasks 
			FROM worker WHERE $1 = ANY(workflows) 
			ORDER BY RANDOM() LIMIT 1
		`

		var (
			worker    Worker
			workflows pgtype.FlatArray[string]
			tasks     pgtype.FlatArray[string]
		)

		err := tx.QueryRow(ctx, q, name).Scan(
			&worker.ID,
			&worker.AdvertiseAddress,
			&workflows,
			&tasks,
		)
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return worker, ErrNotFound
		case err != nil:
			return worker, err
		default:
			worker.Workflows = workflows
			worker.Tasks = tasks

			return worker, nil
		}
	})
}

func (pr *PostgresRepository) Get(ctx context.Context, id string) (Worker, error) {
	return database.Read(ctx, pr.db, func(ctx context.Context, tx pgx.Tx) (Worker, error) {
		const q = `
			SELECT id, advertise_address, workflows, tasks 
			FROM worker WHERE id = $1
		`

		var (
			worker    Worker
			workflows pgtype.FlatArray[string]
			tasks     pgtype.FlatArray[string]
		)

		err := tx.QueryRow(ctx, q, id).Scan(
			&worker.ID,
			&worker.AdvertiseAddress,
			&workflows,
			&tasks,
		)
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return worker, ErrNotFound
		case err != nil:
			return worker, err
		default:
			worker.Workflows = workflows
			worker.Tasks = tasks

			return worker, nil
		}
	})
}
