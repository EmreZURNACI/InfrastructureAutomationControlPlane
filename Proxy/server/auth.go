package server

import (
	"net/http"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/controllers/auth"
	authsrv "github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/service/auth"
)

func (h *routesHandler) StartAuth() {
	route := h.App.Group("/api").Group("/v1").Group("/auth")

	signInService := authsrv.NewSignInService(h.Dependencies.DB)
	authService := auth.AuthService{
		SignInService: signInService,
	}

	route.Add(http.MethodPost, "/signin", authService.SignIn)

}
