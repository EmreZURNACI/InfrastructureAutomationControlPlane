package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type CreateSnapshotRequest struct {
	ID          string  `json:"id" validate:"required"`
	Description *string `json:"description,omitempty" validate:"omitempty"`
}
type CreateSnapshotResponse struct {
	ID *string `json:"id"`
}
type CreateSnapshotHandler struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewCreateSnapshotEbsHandler(client ports.VolumeClient, tp ports.Tracer) *CreateSnapshotHandler {
	return &CreateSnapshotHandler{client: client, tp: tp}
}

func (h *CreateSnapshotHandler) Handle(ctx context.Context, req *CreateSnapshotRequest) (*CreateSnapshotResponse, error) {

	ctx, span := h.tp.Start(ctx, "Create Snapshot")
	defer span.End()

	out, err := h.client.CreateSnapshot(ctx, &ec2.CreateSnapshotInput{
		VolumeId:    &req.ID,
		Description: req.Description,
	})

	if err != nil {
		return nil, nil
	}

	return &CreateSnapshotResponse{
		ID: out.SnapshotId,
	}, nil
}
