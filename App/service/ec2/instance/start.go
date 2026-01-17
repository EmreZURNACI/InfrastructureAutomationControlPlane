package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/instance"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type StartService struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewStartService(client ports.InstanceClient, tp ports.Tracer) *StartService {
	return &StartService{client: client, tp: tp}
}

func (s *StartService) Execute(ctx context.Context, req *instance.StartRequest) (*instance.StartResponse, error) {

	handler := instance.NewStartHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
