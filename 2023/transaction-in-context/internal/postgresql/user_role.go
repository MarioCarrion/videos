package postgresql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/MarioCarrion/videos/2023/transaction-in-context/internal"
)

type UserRole struct {
	conn *pgx.Conn
}

func NewUserRole(conn *pgx.Conn) *UserRole {
	return &UserRole{
		conn: conn,
	}
}

func (u *UserRole) Insert(ctx context.Context, id uuid.UUID, roleIDs ...uuid.UUID) error {
	err := transaction(ctx, u.conn, func(tx pgx.Tx) error {
		urq := userRoleQueries{conn: tx}

		for _, roleID := range roleIDs {
			if err := urq.Insert(ctx, id, roleID); err != nil {
				return fmt.Errorf("Exec %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("transaction %w", err)
	}

	return nil
}

func (u *UserRole) Select(ctx context.Context, id uuid.UUID) (internal.User, error) {
	urq := userRoleQueries{conn: u.conn}

	return urq.Select(ctx, id)
}

type userRoleQueries struct {
	conn DBTX
}

func (u *userRoleQueries) Insert(ctx context.Context, id uuid.UUID, roleID uuid.UUID) error {
	const sql = `INSERT INTO users_roles(user_id, role_id) VALUES ($1, $2)`

	_, err := u.conn.Exec(ctx, sql, &id, &roleID)
	if err != nil {
		return fmt.Errorf("Exec %w", err)
	}

	return nil
}

func (u *userRoleQueries) Select(ctx context.Context, id uuid.UUID) (internal.User, error) {
	const sql = `
SELECT
	U.id AS user_id,
	U.name AS user_name,
	R.id AS role_id,
	R.name AS role_name,
	P.id AS permission_name,
	P.type AS permission_type
FROM
	users U
	LEFT JOIN users_roles UR ON U.id = UR.user_id
	LEFT JOIN roles R ON R.id = UR.role_id
	LEFT JOIN permissions P ON R.id = P.role_id
WHERE
	U.id = $1
`

	rows, err := u.conn.Query(ctx, sql, id)
	if err != nil {
		return internal.User{}, fmt.Errorf("Select %w", err)
	}

	roles := make(map[uuid.UUID]internal.Role)

	var (
		userId         uuid.UUID
		userName       *string
		roleId         uuid.UUID
		roleName       *string
		permissionId   uuid.UUID
		permissionType string
	)

	_, err = pgx.ForEachRow(rows,
		[]any{&userId, &userName, &roleId, &roleName, &permissionId, &permissionType},
		func() error {
			role, ok := roles[roleId]
			if !ok {
				role = internal.Role{
					ID:   roleId,
					Name: *roleName,
				}
			}

			role.Permissions = append(role.Permissions,
				internal.Permission{
					ID:     permissionId,
					RoleID: role.ID,
					Type:   internal.PermissionType(permissionType),
				})

			roles[roleId] = role

			return nil
		})

	if userName == nil {
		return internal.User{}, fmt.Errorf("user not found")
	}

	user := internal.User{
		ID:   userId,
		Name: *userName,
	}

	for _, p := range roles {
		user.Roles = append(user.Roles, p)
	}

	return user, nil
}
