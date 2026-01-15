package server

import (
	"net/http"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers/hello"
)

func (h *routesHandler) StartHello() {

	h.App.Add(http.MethodGet, "/", hello.Hello)
}
