package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/instance"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type EditService struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewEditService(client ports.InstanceClient, tp ports.Tracer) *EditService {
	return &EditService{client: client, tp: tp}
}

func (s *EditService) Execute(ctx context.Context, req *instance.EditRequest) (*instance.EditResponse, error) {

	handler := instance.NewEditHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
