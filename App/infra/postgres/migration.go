package postgres

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/domain"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/log"
)

func (h *DB) migrate() {
	tables := map[string]interface{}{
		"ec2": &domain.EC2{},
	}

	migrator := h.db.Migrator()

	for key, value := range tables {
		if err := migrator.AutoMigrate(value); err == nil {
			log.Logger.Info(key + " tablosu migrate edildi")
		}
	}

}
