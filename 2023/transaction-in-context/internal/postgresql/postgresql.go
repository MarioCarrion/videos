// Package postgresql represents the repositories using the pgx third party library.
package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func transaction(ctx context.Context, tx pgx.Tx, f func() error) error {
	if err := f(); err != nil {
		_ = tx.Rollback(ctx)

		return fmt.Errorf("f %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("Commit %w", err)
	}

	return nil
}
