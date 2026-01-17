package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/instance"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type ListService struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewListService(client ports.InstanceClient, tp ports.Tracer) *ListService {
	return &ListService{client: client, tp: tp}
}

func (s *ListService) Execute(ctx context.Context, req *instance.ListRequest) (*instance.ListResponse, error) {

	handler := instance.NewListHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
