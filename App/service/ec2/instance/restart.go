package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/instance"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type RestartService struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewRestartService(client ports.InstanceClient, tp ports.Tracer) *RestartService {
	return &RestartService{client: client, tp: tp}
}

func (s *RestartService) Execute(ctx context.Context, req *instance.RestartRequest) (*instance.RestartResponse, error) {

	handler := instance.NewRestartHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
