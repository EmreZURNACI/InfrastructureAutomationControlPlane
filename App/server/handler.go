package server

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/adaptor/client"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/infra/postgres"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
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
	EC2Client *client.Ec2Client
	DB        *postgres.DB
	Tp        trace.Tracer
}
