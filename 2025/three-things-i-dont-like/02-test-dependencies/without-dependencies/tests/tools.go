//go:build tools

package postgres

// Use:
// go install -tags 'pgx5' github.com/golang-migrate/migrate/v4/cmd/migrate

import (
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)
