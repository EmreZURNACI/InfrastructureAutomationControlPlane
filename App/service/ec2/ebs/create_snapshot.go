package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/ebs"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type CreateSnapshotService struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewCreateSnapshotService(client ports.VolumeClient, tp ports.Tracer) *CreateSnapshotService {
	return &CreateSnapshotService{client: client, tp: tp}
}

func (s *CreateSnapshotService) Execute(ctx context.Context, req *ebs.CreateSnapshotRequest) (*ebs.CreateSnapshotResponse, error) {

	handler := ebs.NewCreateSnapshotEbsHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
