package auth

import (
	"github.com/google/uuid"
)

type Repository interface {
	GetPermissions(roleID uuid.UUID) []string
	GetRole(role string) uuid.UUID
}
