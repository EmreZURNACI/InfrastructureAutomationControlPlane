package ebs

import (
	"context"
	"errors"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ListRequest struct {
}
type ListResponse struct {
	Volumes []ListedVolume `json:"volumes"`
}
type ListHandler struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewListHandler(client ports.VolumeClient, tp ports.Tracer) *ListHandler {
	return &ListHandler{client: client, tp: tp}
}

func (h *ListHandler) Handle(ctx context.Context, req *ListRequest) (*ListResponse, error) {

	ctx, span := h.tp.Start(ctx, "List Elastic Block Devices")
	defer span.End()

	out, err := h.client.DescribeVolumes(ctx, &ec2.DescribeVolumesInput{})

	if err != nil {
		return nil, err
	}

	if len(out.Volumes) == 0 {
		return nil, errors.New("no block device were found.")
	}

	var volumes []ListedVolume

	for _, volume := range out.Volumes {
		vol := ListedVolume{
			ID:    volume.VolumeId,
			Size:  volume.Size,
			State: volume.State,
		}
		volumes = append(volumes, vol)
	}

	return &ListResponse{
		Volumes: volumes,
	}, nil
}

type ListedVolume struct {
	ID    *string           `json:"id,omitempty"`
	Size  *int32            `json:"size,omitempty"`
	State types.VolumeState `json:"state,omitempty"`
}
