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
	var (
		role internal.Role
		err  error
	)

	err = transaction(ctx, r.conn, func(tx pgx.Tx) error {
		uq := RoleQueries{conn: tx}

		role, err = uq.Insert(ctx, name)
		if err != nil {
			return fmt.Errorf("Insert %w", err)
		}

		for i, p := range permissions {
			permission, err := uq.InsertPermission(ctx, role.ID, p.Type)
			if err != nil {
				return fmt.Errorf("insertPermissionTx %w", err)
			}

			permissions[i] = permission
		}

		role.Permissions = permissions

		return nil
	})

	if err != nil {
		return internal.Role{}, fmt.Errorf("transaction %w", err)
	}

	return role, nil
}

func (r *Role) InsertPermission(ctx context.Context, roleID uuid.UUID, ptype internal.PermissionType) (internal.Permission, error) {
	var (
		permission internal.Permission
		err        error
	)

	err = transaction(ctx, r.conn, func(tx pgx.Tx) error {
		uq := RoleQueries{conn: tx}

		permission, err = uq.InsertPermission(ctx, roleID, ptype)
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

type RoleQueries struct {
	conn DBTX
}

func (r *RoleQueries) Insert(ctx context.Context, name string) (internal.Role, error) {
	const sql = `INSERT INTO roles(name) VALUES ($1) RETURNING id`

	row := r.conn.QueryRow(ctx, sql, &name)

	var id uuid.UUID

	if err := row.Scan(&id); err != nil {
		return internal.Role{}, fmt.Errorf("Insert %w", err)
	}

	return internal.Role{
		ID:   id,
		Name: name,
	}, nil
}

func (r *RoleQueries) InsertPermission(ctx context.Context, roleID uuid.UUID, ptype internal.PermissionType) (internal.Permission, error) {
	const sql = `INSERT INTO permissions(role_id, type) VALUES ($1, $2) RETURNING id`

	row := r.conn.QueryRow(ctx, sql, roleID, &ptype)

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
