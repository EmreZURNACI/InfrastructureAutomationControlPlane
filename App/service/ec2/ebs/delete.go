package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/ebs"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type DeleteService struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewDeleteService(client ports.VolumeClient, tp ports.Tracer) *DeleteService {
	return &DeleteService{client: client, tp: tp}
}

func (s *DeleteService) Execute(ctx context.Context, req *ebs.DeleteRequest) (*ebs.DeleteResponse, error) {

	handler := ebs.NewDeleteHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
