package image

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/image"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type ListService struct {
	client ports.ImageClient
}

func NewListImageService(client ports.ImageClient) *ListService {
	return &ListService{client: client}
}

func (s *ListService) Execute(ctx context.Context, req *image.ListRequest) (*image.ListResponse, error) {

	handler := image.NewListHandler(s.client)
	return handler.Handle(ctx, req)
}
