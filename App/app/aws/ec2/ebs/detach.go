package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type DetachRequest struct {
	ID string `json:"id" validate:"required"`
}
type DetachResponse struct {
}
type DetachHandler struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewDetachHandler(client ports.VolumeClient, tp ports.Tracer) *DetachHandler {
	return &DetachHandler{client: client, tp: tp}
}

func (h *DetachHandler) Handle(ctx context.Context, req *DetachRequest) (*DetachResponse, error) {

	ctx, span := h.tp.Start(ctx, "Detach Elastic Block Device")
	defer span.End()

	_, err := h.client.DetachVolume(ctx, &ec2.DetachVolumeInput{
		VolumeId: &req.ID,
		Force:    aws.Bool(false),
	})
	if err != nil {
		return nil, err
	}

	return &DetachResponse{}, nil
}
