package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/ebs"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type CreateService struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewCreateService(client ports.VolumeClient, tp ports.Tracer) *CreateService {
	return &CreateService{client: client, tp: tp}
}

func (s *CreateService) Execute(ctx context.Context, req *ebs.CreateRequest) (*ebs.CreateResponse, error) {

	handler := ebs.NewCreateHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
