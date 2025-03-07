package postgres

//go:generate sqlc generate

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/MarioCarrion/videos/2025/three-things-i-dont-like/02-test-dependencies/without-dependencies/generated"
)

type User struct {
	ID        string
	Name      string
	BirthYear int
}

type UserRepository struct {
	q *generated.Queries
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		q: generated.New(conn),
	}
}

func (u *UserRepository) Create(ctx context.Context, name string, birthYear int) (*User, error) {
	params := generated.CreateUserParams{
		Name:      name,
		BirthYear: int16(birthYear),
	}

	id, err := u.q.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("Queries.CreateUser: %w", err)
	}

	return &User{
		ID:        id.String(),
		Name:      name,
		BirthYear: birthYear,
	}, nil
}

func (u *UserRepository) Get(ctx context.Context, id string) (*User, error) {
	pid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}

	user, err := u.q.GetUser(ctx, pid)
	if err != nil {
		return nil, fmt.Errorf("Queries.GetUser: %w", err)
	}

	return &User{
		ID:        id,
		Name:      user.Name,
		BirthYear: int(user.BirthYear),
	}, nil
}
