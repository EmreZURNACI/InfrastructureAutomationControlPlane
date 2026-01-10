package postgres

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/log"
)

func (h *DB) createExtension() error {
	if err := h.db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Logger.Error("Extension kurulamadi")
		return err
	}
	return nil
}
