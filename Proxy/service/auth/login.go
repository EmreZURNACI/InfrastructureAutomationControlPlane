package auth

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/app/auth"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/ports"
)

type SignInService struct {
	repo auth.Repository
	ldap ports.LDAP
}

func NewSignInService(repo auth.Repository, ldap ports.LDAP) *SignInService {
	return &SignInService{repo: repo, ldap: ldap}
}

func (s *SignInService) Execute(ctx context.Context, req *auth.SignInRequest) (*auth.SignInResponse, error) {

	handler := auth.NewSignInHandler(s.repo, s.ldap)
	return handler.Handle(ctx, req)
}
