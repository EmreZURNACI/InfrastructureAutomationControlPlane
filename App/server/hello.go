package server

import (
	"net/http"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"
)

func (h *routesHandler) StartHello() {

	h.App.Add(http.MethodGet, "/", controllers.Hello)
}
