package key

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/key"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type ListService struct {
	client ports.KeyClient
}

func NewListService(client ports.KeyClient) *ListService {
	return &ListService{client: client}
}

func (s *ListService) Execute(ctx context.Context, req *key.ListRequest) (*key.ListResponse, error) {

	handler := key.NewListHandler(s.client)
	return handler.Handle(ctx, req)
}
