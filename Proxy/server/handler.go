package server

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/adaptor/ldap"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/infra/postgres"

	"github.com/gofiber/fiber/v2"
)

func NewRoutesHandler(app *fiber.App, dep Dependencies) *routesHandler {
	return &routesHandler{
		App:          app,
		Dependencies: dep,
	}
}

type routesHandler struct {
	Dependencies Dependencies
	App          *fiber.App
}
type Dependencies struct {
	DB   *postgres.DB
	Ldap *ldap.LdapConnection
}
