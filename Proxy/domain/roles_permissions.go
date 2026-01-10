package domain

import (
	"github.com/google/uuid"
)

// type RolesPermission struct {
// 	RoleID       uuid.UUID `gorm:"type:uuid;primaryKey;column:roles_id"`
// 	PermissionID uuid.UUID `gorm:"type:uuid;primaryKey;column:permissions_id"`

// 	Role       Role       `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASCADE"`
// 	Permission Permission `gorm:"foreignKey:PermissionID;references:ID;constraint:OnDelete:CASCADE"`
// }

type RolesPermission struct {
	RoleID       uuid.UUID `gorm:"type:uuid;primaryKey;column:roles_id"`
	PermissionID uuid.UUID `gorm:"type:uuid;primaryKey;column:permissions_id"`
}
