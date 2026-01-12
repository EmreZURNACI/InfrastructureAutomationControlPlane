package auth

import "github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/service/auth"

type AuthService struct {
	SignInService *auth.SignInService
	SignUpService *auth.SignUpService
}
