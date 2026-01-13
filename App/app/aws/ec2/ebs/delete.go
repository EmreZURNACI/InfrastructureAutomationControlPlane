package ebs

import (
	"context"
	"fmt"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/log"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type DeleteRequest struct {
	ID string `json:"id" validate:"required"`
}
type DeleteResponse struct {
}
type DeleteHandler struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewDeleteHandler(client ports.VolumeClient, tp ports.Tracer) *DeleteHandler {
	return &DeleteHandler{client: client, tp: tp}
}

func (h *DeleteHandler) Handle(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {

	ctx, span := h.tp.Start(ctx, "Delete Elastic Block Device")
	defer span.End()

	out, err := h.client.DeleteVolume(ctx, &ec2.DeleteVolumeInput{
		VolumeId: &req.ID,
	})
	if err != nil {
		log.Logger.Error("ebs did not delete")
		return nil, err
	}
	fmt.Println(out.ResultMetadata)
	return &DeleteResponse{}, nil
}
