package ebs

import (
	"context"
	"errors"
	"time"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type DetailRequest struct {
	ID string `json:"id" validate:"required"`
}
type DetailResponse struct {
	Volume Volume `json:"volume"`
}
type DetailHandler struct {
	client ports.VolumeClient
	tp     ports.Tracer
}

func NewDetailHandler(client ports.VolumeClient, tp ports.Tracer) *DetailHandler {
	return &DetailHandler{client: client, tp: tp}
}

func (h *DetailHandler) Handle(ctx context.Context, req *DetailRequest) (*DetailResponse, error) {

	ctx, span := h.tp.Start(ctx, "Detail Elastic Block Device")
	defer span.End()

	out, err := h.client.DescribeVolumes(ctx, &ec2.DescribeVolumesInput{
		VolumeIds: []string{req.ID},
	})

	if err != nil {
		return nil, err
	}

	if len(out.Volumes) == 0 {
		return nil, errors.New("no volume were found with this ID.")
	}

	vol := Volume{
		VolumeId:         out.Volumes[0].VolumeId,
		AvailabilityZone: out.Volumes[0].AvailabilityZone,
		CreatedTime:      out.Volumes[0].CreateTime,
		Iops:             out.Volumes[0].Iops,
		Throughput:       out.Volumes[0].Throughput,
		Encrypted:        out.Volumes[0].Encrypted,
		Size:             out.Volumes[0].Size,
		VolumeType:       out.Volumes[0].VolumeType,
		State:            out.Volumes[0].State,
	}

	var attachments []VolumeAttachments
	for _, attachment := range out.Volumes[0].Attachments {
		attachments = append(attachments, VolumeAttachments{
			InstanceId:   attachment.InstanceId,
			Device:       attachment.Device,
			AttachedTime: attachment.AttachTime,
		})
	}
	vol.Attachments = attachments

	return &DetailResponse{
		Volume: vol,
	}, nil
}

type Volume struct {
	VolumeId         *string
	AvailabilityZone *string
	CreatedTime      *time.Time
	Iops             *int32
	Throughput       *int32
	Encrypted        *bool
	Size             *int32
	VolumeType       types.VolumeType
	State            types.VolumeState
	Attachments      []VolumeAttachments
}

type VolumeAttachments struct {
	InstanceId   *string
	Device       *string
	AttachedTime *time.Time
}
