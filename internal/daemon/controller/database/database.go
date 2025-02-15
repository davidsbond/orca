package database

import (
	"context"
	"embed"
	_ "embed"
	"errors"
	"net/url"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func Open(ctx context.Context, addr string) (*pgx.Conn, error) {
	if err := MigrateUp(addr); err != nil {
		return nil, err
	}

	db, err := pgx.Connect(ctx, addr)
	if err != nil {
		return nil, err
	}

	return db, db.Ping(ctx)
}

type (
	WriteTransaction       func(ctx context.Context, tx pgx.Tx) error
	ReadTransaction[T any] func(ctx context.Context, tx pgx.Tx) (T, error)
)

func Write(ctx context.Context, db *pgx.Conn, fn WriteTransaction) error {
	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	if err = fn(ctx, tx); err != nil {
		if txErr := tx.Rollback(ctx); txErr != nil {
			return errors.Join(err, txErr)
		}

		return err
	}

	return tx.Commit(ctx)
}

func Read[T any](ctx context.Context, db *pgx.Conn, fn ReadTransaction[T]) (T, error) {
	var result T

	tx, err := db.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadOnly})
	if err != nil {
		return result, err
	}

	result, err = fn(ctx, tx)
	if err != nil {
		if txErr := tx.Rollback(ctx); txErr != nil {
			return result, errors.Join(err, txErr)
		}

		return result, err
	}

	return result, tx.Commit(ctx)
}

func IsForeignKeyViolation(err error, column string) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == pgerrcode.ForeignKeyViolation && strings.Contains(pgErr.ConstraintName, column)
	}

	return false
}

var (
	//go:embed migrations/*.sql
	migrations embed.FS
)

// MigrateUp updates the database schema to the latest migration script.
func MigrateUp(addr string) error {
	source, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	u, err := url.Parse(addr)
	if err != nil {
		return err
	}

	// Switch the url scheme to pgx5 as that is the destination driver for migrations.
	u.Scheme = "pgx5"
	migration, err := migrate.NewWithSourceInstance("iofs", source, u.String())
	if err != nil {
		return err
	}

	err = migration.Up()
	switch {
	case errors.Is(err, migrate.ErrNoChange):
		return errors.Join(migration.Close())
	case err != nil:
		// I'm lazy and this is very ugly.
		return errors.Join(err, errors.Join(migration.Close()))
	default:
		return errors.Join(migration.Close())
	}
}
