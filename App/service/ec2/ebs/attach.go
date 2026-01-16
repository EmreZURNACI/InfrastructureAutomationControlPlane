package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/ebs"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type AttachService struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewAttachService(client ports.VolumeClient, tp ports.Tracer) *AttachService {
	return &AttachService{client: client, tp: tp}
}

func (s *AttachService) Execute(ctx context.Context, req *ebs.AttachRequest) (*ebs.AttachResponse, error) {

	handler := ebs.NewAttachHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
