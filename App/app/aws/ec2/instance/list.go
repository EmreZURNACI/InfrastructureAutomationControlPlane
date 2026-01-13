package instance

import (
	"context"
	"errors"
	"time"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ListRequest struct {
}
type ListResponse struct {
	Instances []Instance `json:"instances"`
}
type ListHandler struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewListHandler(client ports.InstanceClient, tp ports.Tracer) *ListHandler {
	return &ListHandler{client: client, tp: tp}
}

func (h *ListHandler) Handle(ctx context.Context, req *ListRequest) (*ListResponse, error) {

	ctx, span := h.tp.Start(ctx, "List Instances")
	defer span.End()

	state_filter := "instance-state-name"
	out, err := h.client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name: &state_filter,
				Values: []string{
					"pending", "running", "shutting-down", "stopping", "stopped",
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if len(out.Reservations) == 0 {
		return nil, errors.New("no instances were found.")
	}

	var instances []Instance
	for _, value := range out.Reservations {
		instances = append(instances, Instance{
			InstanceID:          value.Instances[0].InstanceId,
			ImageID:             value.Instances[0].ImageId,
			BlockDeviceMappings: value.Instances[0].BlockDeviceMappings,
			CPU: &CPU{
				CoreCount:      value.Instances[0].CpuOptions.CoreCount,
				ThreadsPerCore: value.Instances[0].CpuOptions.ThreadsPerCore,
			},
			InstanceType: &value.Instances[0].InstanceType,
			PublicIpAddr: value.Instances[0].PublicIpAddress,
			LaunchTime:   value.Instances[0].LaunchTime,
			State:        &value.Instances[0].State.Name,
			RAM:          value.Instances[0].RamdiskId,
		})
	}

	return &ListResponse{
		Instances: instances,
	}, nil
}

type Instance struct {
	InstanceID          *string                            `json:"instance_id,omitempty"`
	ImageID             *string                            `json:"image_id,omitempty"`
	BlockDeviceMappings []types.InstanceBlockDeviceMapping `json:"block_device_mapping,omitempty"`
	CPU                 *CPU                               `json:"cpu,omitempty"`
	InstanceType        *types.InstanceType                `json:"instance_type,omitempty"`
	PublicIpAddr        *string                            `json:"public_ip_addr,omitempty"`
	LaunchTime          *time.Time                         `json:"launch_time,omitempty"`
	State               *types.InstanceStateName           `json:"state,omitempty"`
	RAM                 *string                            `json:"ram,omitempty"`
}

type CPU struct {
	CoreCount      *int32 `json:"core_count"`
	ThreadsPerCore *int32 `json:"threads_per_core"`
}
