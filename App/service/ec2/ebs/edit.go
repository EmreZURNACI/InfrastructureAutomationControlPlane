package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/ebs"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type EditService struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewEditService(client ports.VolumeClient, tp ports.Tracer) *EditService {
	return &EditService{client: client, tp: tp}
}

func (s *EditService) Execute(ctx context.Context, req *ebs.EditRequest) (*ebs.EditResponse, error) {

	handler := ebs.NewEditHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
