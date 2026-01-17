package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/instance"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type TerminateService struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewTerminateService(client ports.InstanceClient, tp ports.Tracer) *TerminateService {
	return &TerminateService{client: client, tp: tp}
}

func (s *TerminateService) Execute(ctx context.Context, req *instance.TerminateRequest) (*instance.TerminateResponse, error) {

	handler := instance.NewTerminateHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
