package postgresql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/MarioCarrion/videos/2023/transaction-in-context/internal"
)

type Role struct {
	conn *pgx.Conn
}

func NewRole(conn *pgx.Conn) *Role {
	return &Role{
		conn: conn,
	}
}

func (r *Role) Insert(ctx context.Context, name string, permissions []internal.Permission) (internal.Role, error) {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return internal.Role{}, fmt.Errorf("Begin %w", err)
	}

	const sql = `INSERT INTO roles(name) VALUES ($1) RETURNING id`

	var role internal.Role

	err = transaction(ctx, tx, func() error {
		row := tx.QueryRow(ctx, sql, &name)

		var id uuid.UUID

		if err = row.Scan(&id); err != nil {
			return fmt.Errorf("Insert %w", err)
		}

		for i, p := range permissions {
			permission, err := r.insertPermissionTx(ctx, tx, id, p.Type)
			if err != nil {
				return fmt.Errorf("insertPermissionTx %w", err)
			}

			permissions[i] = permission
		}

		role.ID = id
		role.Name = name
		role.Permissions = permissions

		return nil
	})
	if err != nil {
		return internal.Role{}, fmt.Errorf("transaction %w", err)
	}

	return role, nil
}

func (r *Role) InsertPermission(ctx context.Context, roleID uuid.UUID, ptype internal.PermissionType) (internal.Permission, error) {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return internal.Permission{}, fmt.Errorf("Begin %w", err)
	}

	var permission internal.Permission

	err = transaction(ctx, tx, func() error {
		permission, err = r.insertPermissionTx(ctx, tx, roleID, ptype)
		if err != nil {
			return fmt.Errorf("insertPermission %w", err)
		}

		return nil
	})
	if err != nil {
		return internal.Permission{}, fmt.Errorf("transaction %w", err)
	}

	return permission, nil
}

func (r *Role) insertPermissionTx(ctx context.Context, tx pgx.Tx, roleID uuid.UUID, ptype internal.PermissionType) (internal.Permission, error) {
	const sql = `INSERT INTO permissions(role_id, type) VALUES ($1, $2) RETURNING id`

	row := tx.QueryRow(ctx, sql, roleID, &ptype)

	var id uuid.UUID

	if err := row.Scan(&id); err != nil {
		return internal.Permission{}, fmt.Errorf("Insert %w", err)
	}

	return internal.Permission{
		ID:     id,
		RoleID: roleID,
		Type:   ptype,
	}, nil
}
