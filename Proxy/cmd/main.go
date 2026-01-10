package main

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/infra/postgres"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/log"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/server"
)

func main() {

	db, err := postgres.Connection()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	if err := db.PrepareDB(); err != nil {
		log.Logger.Error(err.Error())
		return
	}

	server.Start(server.Dependencies{
		DB: db,
	})

}
