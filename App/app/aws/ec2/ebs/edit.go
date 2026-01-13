package ebs

import (
	"context"
	"errors"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type EditRequest struct {
	ID         string  `json:"id" validate:"required"`
	Type       *string `json:"type" validate:"omitempty"`
	Size       *int32  `json:"size" validate:"omitempty"`
	Throughput *int32  `json:"throughput" validate:"omitempty"`
	Iops       *int32  `json:"iops" validate:"omitempty"`
}
type EditResponse struct {
	State types.VolumeModificationState
}
type EditHandler struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewEditHandler(client ports.VolumeClient, tp ports.Tracer) *EditHandler {
	return &EditHandler{client: client, tp: tp}
}

func (h *EditHandler) Handle(ctx context.Context, req *EditRequest) (*EditResponse, error) {

	ctx, span := h.tp.Start(ctx, "Edit Elastic Block Device")
	defer span.End()

	vol, err := h.client.DescribeVolumes(ctx, &ec2.DescribeVolumesInput{
		VolumeIds: []string{req.ID},
	})

	if len(vol.Volumes) == 0 {
		return nil, errors.New("volume not found")
	}

	size := vol.Volumes[0].Size
	if req.Size != nil && *req.Size != 0 {
		size = req.Size
	}
	throughput := vol.Volumes[0].Throughput
	if req.Throughput != nil && *req.Throughput != 0 {
		throughput = req.Throughput
	}
	iops := vol.Volumes[0].Iops
	if req.Iops != nil && *req.Iops != 0 {
		iops = req.Iops
	}
	volType := vol.Volumes[0].VolumeType
	if req.Type != nil && *req.Type != "" {
		volType = types.VolumeType(*req.Type)
	}

	out, err := h.client.EditVolume(ctx, &ec2.ModifyVolumeInput{
		VolumeId:   &req.ID,
		Size:       size,
		Throughput: throughput,
		Iops:       iops,
		VolumeType: volType,
	})
	if err != nil {
		return nil, err
	}

	return &EditResponse{
		State: out.VolumeModification.ModificationState,
	}, nil
}
