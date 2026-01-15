package server

import (
	"net/http"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers/liveness"
)

func (h *routesHandler) StartLiveness() {

	h.App.Add(http.MethodGet, "/liveness", liveness.Liveness)
}
