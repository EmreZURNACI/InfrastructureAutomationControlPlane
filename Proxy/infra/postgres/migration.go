package postgres

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/domain"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/log"
)

func (h *DB) migrate() {
	tables := map[string]interface{}{
		"roles":             &domain.Role{},
		"permissions":       &domain.Permission{},
		"users":             &domain.User{},
		"roles_permissions": &domain.RolesPermission{},
	}

	migrator := h.db.Migrator()

	for key, value := range tables {
		if err := migrator.AutoMigrate(value); err == nil {
			log.Logger.Info(key + " tablosu migrate edildi")
		}
	}

}
