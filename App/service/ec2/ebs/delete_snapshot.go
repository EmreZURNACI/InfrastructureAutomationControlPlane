package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/ebs"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type DeleteSnapshotService struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewDeleteSnapshotService(client ports.VolumeClient, tp ports.Tracer) *DeleteSnapshotService {
	return &DeleteSnapshotService{client: client, tp: tp}
}

func (s *DeleteSnapshotService) Execute(ctx context.Context, req *ebs.DeleteSnapshotRequest) (*ebs.DeleteSnapshotResponse, error) {

	handler := ebs.NewDeleteSnapshotHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
