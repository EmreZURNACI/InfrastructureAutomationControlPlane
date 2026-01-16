package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/ebs"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type DetachService struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewDetachService(client ports.VolumeClient, tp ports.Tracer) *DetachService {
	return &DetachService{client: client, tp: tp}
}

func (s *DetachService) Execute(ctx context.Context, req *ebs.DetachRequest) (*ebs.DetachResponse, error) {

	handler := ebs.NewDetachHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
