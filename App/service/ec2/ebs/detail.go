package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/ebs"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type DetailService struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewDetailService(client ports.VolumeClient, tp ports.Tracer) *DetailService {
	return &DetailService{client: client, tp: tp}
}

func (s *DetailService) Execute(ctx context.Context, req *ebs.DetailRequest) (*ebs.DetailResponse, error) {

	handler := ebs.NewDetailHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
