package worker

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

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
		db *pgx.Conn
	}
)

var (
	ErrNotFound = errors.New("not found")
)

func NewPostgresRepository(db *pgx.Conn) *PostgresRepository {
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
