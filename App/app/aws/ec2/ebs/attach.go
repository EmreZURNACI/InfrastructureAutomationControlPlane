package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type AttachRequest struct {
	InstanceID string `json:"instance_id" validate:"required"`
	VolumeID   string `json:"volume_id" validate:"required"`
	Device     string `json:"device" validate:"required"`
}
type AttachResponse struct {
}
type AttachHandler struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewAttachHandler(client ports.VolumeClient, tp ports.Tracer) *AttachHandler {
	return &AttachHandler{
		client: client,
		tp:     tp,
	}
}

func (h *AttachHandler) Handle(ctx context.Context, req *AttachRequest) (*AttachResponse, error) {

	ctx, span := h.tp.Start(ctx, "Attach Elastic Block Device")
	defer span.End()

	_, err := h.client.AttachVolume(ctx, &ec2.AttachVolumeInput{
		InstanceId: &req.InstanceID,
		VolumeId:   &req.VolumeID,
		Device:     &req.Device,
	})
	if err != nil {
		return nil, err
	}

	return &AttachResponse{}, nil
}
