package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/instance"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type StopService struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewStopService(client ports.InstanceClient, tp ports.Tracer) *StopService {
	return &StopService{client: client, tp: tp}
}

func (s *StopService) Execute(ctx context.Context, req *instance.StopRequest) (*instance.StopResponse, error) {

	handler := instance.NewStopHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
