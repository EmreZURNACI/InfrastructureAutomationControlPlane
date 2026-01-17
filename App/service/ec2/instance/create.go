package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/instance"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type CreateService struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewCreateService(client ports.InstanceClient, tp ports.Tracer) *CreateService {
	return &CreateService{client: client, tp: tp}
}

func (s *CreateService) Execute(ctx context.Context, req *instance.CreateRequest) (*instance.CreateResponse, error) {

	handler := instance.NewCreateInstanceHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
