// Package internal defines the domain types used for this program.
package internal

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Name  string
	Roles []Role
}

type Role struct {
	ID          uuid.UUID
	Name        string
	Permissions []Permission
}

type PermissionType string

type Permission struct {
	ID     uuid.UUID
	RoleID uuid.UUID
	Type   PermissionType
}
