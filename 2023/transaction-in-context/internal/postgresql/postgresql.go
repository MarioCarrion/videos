// Package postgresql represents the repositories using the pgx third party library.
package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	Prepare(context.Context, string, string) (*pgconn.StatementDescription, error)
}

func transaction(ctx context.Context, conn *pgx.Conn, f func(tx pgx.Tx) error) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("Begin %w", err)
	}

	if err := f(tx); err != nil {
		_ = tx.Rollback(ctx)

		return fmt.Errorf("f %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("Commit %w", err)
	}

	return nil
}
