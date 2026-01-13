package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type DeleteSnapshotRequest struct {
	ID string `json:"id" validate:"required"`
}
type DeleteSnapshotResponse struct {
}
type DeleteSnapshotHandler struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewDeleteSnapshotHandler(client ports.VolumeClient, tp ports.Tracer) *DeleteSnapshotHandler {
	return &DeleteSnapshotHandler{client: client, tp: tp}
}

func (h *DeleteSnapshotHandler) Handle(ctx context.Context, req *DeleteSnapshotRequest) (*DeleteSnapshotResponse, error) {

	ctx, span := h.tp.Start(ctx, "Delete Snapshot")
	defer span.End()

	_, err := h.client.DeleteSnapshot(ctx, &ec2.DeleteSnapshotInput{
		SnapshotId: &req.ID,
	})

	if err != nil {
		return nil, nil
	}

	return &DeleteSnapshotResponse{}, nil
}
