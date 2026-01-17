package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/instance"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type ListTypesService struct {
	client ports.InstanceClient
}

func NewListTypesService(client ports.InstanceClient) *ListTypesService {
	return &ListTypesService{client: client}
}

func (s *ListTypesService) Execute(ctx context.Context, req *instance.ListInstanceTypeRequest) (*instance.ListInstanceTypeResponse, error) {

	handler := instance.NewListInstanceTypeHandler(s.client)
	return handler.Handle(ctx, req)
}
