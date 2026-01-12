package main

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/adaptor/ldap"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/infra/postgres"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/log"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/server"
)

func main() {

	db, err := postgres.Connect()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	if err := db.Prepare(); err != nil {
		log.Logger.Error(err.Error())
		return
	}
	ldap, err := ldap.Connect()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	server.Start(server.Dependencies{
		DB:   db,
		Ldap: ldap,
	})
}
