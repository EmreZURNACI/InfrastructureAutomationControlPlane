package server

import (
	"net/http"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"
)

func (h *routesHandler) StartLiveness() {

	h.App.Add(http.MethodGet, "/liveness", controllers.Liveness)
}
