package ebs

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/log"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type CreateRequest struct {
	AvailabilityZone string `json:"availability_zone,omitempty" validate:"required"`
	Size             int32  `json:"size" validate:"required"`
	Type             string `json:"type" validate:"required"`
	Iops             int32  `json:"iops" validate:"required"`
	Throughput       int32  `json:"throughput" validate:"required"`
}
type CreateResponse struct {
	ID *string `json:"id"`
}
type CreateHandler struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewCreateHandler(client ports.VolumeClient, tp ports.Tracer) *CreateHandler {
	return &CreateHandler{client: client, tp: tp}
}

func (h *CreateHandler) Handle(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {

	ctx, span := h.tp.Start(ctx, "Create Elastic Block Device")
	defer span.End()

	out, err := h.client.CreateVolume(ctx, &ec2.CreateVolumeInput{
		AvailabilityZone:   &req.AvailabilityZone,
		Size:               &req.Size,
		VolumeType:         types.VolumeType(req.Type),
		MultiAttachEnabled: aws.Bool(false),
		Encrypted:          aws.Bool(true),
		Throughput:         &req.Throughput,
		Iops:               &req.Iops,
	})
	if err != nil {
		log.Logger.Error("ebs did not create")
		return nil, err
	}

	return &CreateResponse{
		ID: out.VolumeId,
	}, nil
}
