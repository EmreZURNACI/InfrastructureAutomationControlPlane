package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/ebs"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type ListService struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewListService(client ports.VolumeClient, tp ports.Tracer) *ListService {
	return &ListService{client: client, tp: tp}
}

func (s *ListService) Execute(ctx context.Context, req *ebs.ListRequest) (*ebs.ListResponse, error) {

	handler := ebs.NewListHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
