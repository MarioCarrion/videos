package postgresql

import (
	"context"
	"fmt"

	"github.com/MarioCarrion/videos/2023/transaction-in-context/internal"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserCloner struct {
	conn *pgx.Conn
}

func NewUserCloner(conn *pgx.Conn) *UserCloner {
	return &UserCloner{
		conn: conn,
	}
}

func (u *UserCloner) Clone(ctx context.Context, id uuid.UUID, name string) (internal.User, error) {
	var user internal.User

	transaction(ctx, u.conn, func(tx pgx.Tx) error {
		urq := userRoleQueries{conn: tx}

		userFound, err := urq.Select(ctx, id)
		if err != nil {
			return fmt.Errorf("urq.Select(1) %w", err)
		}

		uq := userQueries{conn: tx}

		userNew, err := uq.Insert(ctx, name)
		if err != nil {
			return fmt.Errorf("uq.Insert %w", err)
		}

		for _, role := range userFound.Roles {
			if err := urq.Insert(ctx, userNew.ID, role.ID); err != nil {
				return fmt.Errorf("urq.Insert %w", err)
			}
		}

		userFound, err = urq.Select(ctx, userNew.ID)
		if err != nil {
			return fmt.Errorf("urq.Select(2) %w", err)
		}

		user = userFound

		return nil
	})

	return user, nil
}
