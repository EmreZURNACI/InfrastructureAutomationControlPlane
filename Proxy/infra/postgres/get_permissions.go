package postgres

import (
	"github.com/google/uuid"
)

func (h *DB) GetPermissions(roleID uuid.UUID) []string {
	var permissions []string
	err := h.db.
		Table("permissions AS p").
		Select("p.name").
		Joins("INNER JOIN roles_permissions rp ON rp.permissions_id = p.id").
		Where("rp.roles_id = ?", roleID).
		Scan(&permissions).Error

	if err != nil {
		return nil
	}

	return permissions

}

// SELECT name FROM roles_permissions rp
// 	INNER JOIN permissions p ON p.id=rp.permissions_id
// 		WHERE roles_id='d3e2103f-baf1-4e47-87a0-5dfabf5dadc8';
