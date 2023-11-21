package postgresql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/MarioCarrion/videos/2023/transaction-in-context/internal"
)

type User struct {
	conn *pgx.Conn
}

func NewUser(conn *pgx.Conn) *User {
	return &User{
		conn: conn,
	}
}

func (u *User) Insert(ctx context.Context, name string) (internal.User, error) {
	uq := userQueries{conn: u.conn}

	return uq.Insert(ctx, name)
}

type userQueries struct {
	conn DBTX
}

func (u *userQueries) Insert(ctx context.Context, name string) (internal.User, error) {
	const sql = `INSERT INTO users(name) VALUES ($1) RETURNING id`

	row := u.conn.QueryRow(ctx, sql, &name)

	var id uuid.UUID

	if err := row.Scan(&id); err != nil {
		return internal.User{}, fmt.Errorf("Insert %w", err)
	}

	return internal.User{
		ID:   id,
		Name: name,
	}, nil
}
