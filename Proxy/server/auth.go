package server

import (
	"net/http"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/controllers/auth"
	authsrv "github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/service/auth"
)

func (h *routesHandler) StartAuth() {
	route := h.App.Group("/api").Group("/v1").Group("/auth")

	signInService := authsrv.NewSignInService(h.Dependencies.DB, h.Dependencies.Ldap)
	SignUpService := authsrv.NewSignUpService(h.Dependencies.Ldap)
	authService := auth.AuthService{
		SignInService: signInService,
		SignUpService: SignUpService,
	}

	route.Add(http.MethodPost, "/signin", authService.SignIn)
	route.Add(http.MethodPost, "/signup", authService.SignUp)

}
