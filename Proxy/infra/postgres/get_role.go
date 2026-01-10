package postgres

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/domain"
	"github.com/google/uuid"
)

func (h *DB) GetRole(role string) uuid.UUID {
	var roleID string
	err := h.db.Model(&domain.Role{}).Select("id").Where("name = ?", role).Scan(&roleID).Error
	if err != nil {
		return uuid.UUID{}
	}
	return uuid.MustParse(roleID)
}
