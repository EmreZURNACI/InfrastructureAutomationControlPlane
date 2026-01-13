package instance

import (
	"context"
	"errors"
	"time"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/ptr"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type DetailRequest struct {
	ID string `json:"id" validate:"required"`
}
type DetailResponse struct {
	Instance InstanceDetail `json:"instance"`
}
type DetailHandler struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewDetailHandler(client ports.InstanceClient, tp ports.Tracer) *DetailHandler {
	return &DetailHandler{client: client, tp: tp}
}

func (h *DetailHandler) Handle(ctx context.Context, req *DetailRequest) (*DetailResponse, error) {

	ctx, span := h.tp.Start(ctx, "Detail Instance")
	defer span.End()

	out, err := h.client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name: ptr.String("instance-id"),
				Values: []string{
					req.ID,
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if len(out.Reservations) == 0 {
		return nil, errors.New("no instance were found with this ID.")
	}

	instance := createDetailInput(out.Reservations[0].Instances[0])
	return &DetailResponse{
		Instance: *instance,
	}, nil

}

type InstanceDetail struct {
	InstanceID            *string                            `json:"instance_id,omitempty"`
	ImageID               *string                            `json:"image_id,omitempty"`
	HyperVisor            *types.HypervisorType              `json:"hypervisor,omitempty"`
	Architecture          *types.ArchitectureValues          `json:"architecture,omitempty"`
	BootMode              *types.InstanceBootModeValues      `json:"boot_mode,omitempty"`
	BlockDeviceMappings   []types.InstanceBlockDeviceMapping `json:"block_device_mapping,omitempty"`
	CPU                   CPU                                `json:"cpu,omitempty"`
	InstanceType          *types.InstanceType                `json:"instance_type,omitempty"`
	NetworkInterfaces     []types.InstanceNetworkInterface   `json:"network_interfaces,omitempty"`
	PublicIpAddr          *string                            `json:"public_ip_addr,omitempty"`
	PublicDNSName         *string                            `json:"public_dns_name,omitempty"`
	PrivateDNSName        *string                            `json:"private_dns_name,omitempty"`
	PrivateDNSNameOptions *types.HostnameType                `json:"private_dns_name_options,omitempty"`
	PrivateIpAddress      *string                            `json:"private_ip_addr,omitempty"`
	Ipv6                  *string                            `json:"ipv6,omitempty"`
	LaunchTime            *time.Time                         `json:"launch_time,omitempty"`
	SecurityGroups        []types.GroupIdentifier            `json:"security_groups,omitempty"`
	State                 *types.InstanceStateName           `json:"state,omitempty"`
	VpcID                 *string                            `json:"vpc_id,omitempty"`
	RAM                   *string                            `json:"ram,omitempty"`
	VirtualizationType    *types.VirtualizationType          `json:"virtualization_type,omitempty"`
	Monitoring            *types.MonitoringState             `json:"monitoring,omitempty"`
	AmiLaunchIndex        *int32                             `json:"ami_launch_index,omitempty"`
}

func createDetailInput(instance types.Instance) *InstanceDetail {
	return &InstanceDetail{
		ImageID:             instance.ImageId,
		HyperVisor:          &instance.Hypervisor,
		Architecture:        &instance.Architecture,
		BootMode:            &instance.CurrentInstanceBootMode,
		BlockDeviceMappings: instance.BlockDeviceMappings,
		CPU: CPU{
			CoreCount:      instance.CpuOptions.CoreCount,
			ThreadsPerCore: instance.CpuOptions.ThreadsPerCore,
		},
		InstanceType:          &instance.InstanceType,
		NetworkInterfaces:     instance.NetworkInterfaces,
		PublicIpAddr:          instance.PublicIpAddress,
		PublicDNSName:         instance.PublicDnsName,
		PrivateDNSName:        instance.PrivateDnsName,
		PrivateDNSNameOptions: &instance.PrivateDnsNameOptions.HostnameType,
		PrivateIpAddress:      instance.PrivateIpAddress,
		Ipv6:                  instance.Ipv6Address,
		LaunchTime:            instance.LaunchTime,
		SecurityGroups:        instance.SecurityGroups,
		State:                 &instance.State.Name,
		VpcID:                 instance.VpcId,
		RAM:                   instance.RamdiskId,
		VirtualizationType:    &instance.VirtualizationType,
		Monitoring:            &instance.Monitoring.State,
		AmiLaunchIndex:        instance.AmiLaunchIndex,
	}
}
