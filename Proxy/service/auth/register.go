package auth

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/app/auth"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/ports"
)

type SignUpService struct {
	ldap ports.LDAP
}

func NewSignUpService(ldap ports.LDAP) *SignUpService {
	return &SignUpService{ldap: ldap}
}

func (s *SignUpService) Execute(ctx context.Context, req *auth.SignUpRequest) (*auth.SignUpResponse, error) {

	handler := auth.NewSignUpHandler(s.ldap)
	return handler.Handle(ctx, req)
}
