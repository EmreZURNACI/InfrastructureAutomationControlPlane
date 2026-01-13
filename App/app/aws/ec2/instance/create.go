package instance

import (
	"context"
	"encoding/json"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/domain"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/ptr"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type CreateRequest struct {
	InstanceName     string   `json:"instance_name" validate:"required"`
	InstanceType     string   `json:"instance_type" validate:"required"`
	ImageID          string   `json:"image_id" validate:"required"`
	Volume           []Volume `json:"volume" validate:"required"`
	Min              int32    `json:"min" validate:"required"`
	Max              int32    `json:"max" validate:"required"`
	KeyName          string   `json:"key_name" validate:"required"`
	SubnetID         string   `json:"subnet_id" validate:"required"`
	SecurityGroupIDS []string `json:"security_group_ids" validate:"required"`
}
type CreateResponse struct {
	States []State `json:"states"`
}

type CreateHandler struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewCreateInstanceHandler(client ports.InstanceClient, tp ports.Tracer) *CreateHandler {
	return &CreateHandler{client: client, tp: tp}
}

type Volume struct {
	Size int32 `json:"size"`
}
type State struct {
	ID   *string                 `json:"id"`
	Name types.InstanceStateName `json:"name"`
}

func (h *CreateHandler) Handle(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {

	ctx, span := h.tp.Start(ctx, "Create Instance")
	defer span.End()

	blockDevices := createBlockDevice(req.Volume)

	instanceInput := createInstanceInput(req, blockDevices)

	out, err := h.client.RunInstances(ctx, instanceInput)

	if err != nil {
		return nil, err
	}

	output := output(out.Instances)

	//dbInstanceInput := createDBInstanceInput(out.Instances, blockDevices)

	//postgres.CreateInstance(ctx, dbInstanceInput)

	return &CreateResponse{
		States: output,
	}, nil
}

func createBlockDevice(volumes []Volume) (blockDevices []types.BlockDeviceMapping) {

	number := 98 //b
	for _, value := range volumes {
		name := "xvd" + string(rune(number))
		blockDevices = append(blockDevices, types.BlockDeviceMapping{
			DeviceName: &name,
			Ebs: &types.EbsBlockDevice{
				VolumeSize:          &value.Size,
				VolumeType:          types.VolumeTypeGp2,
				DeleteOnTermination: ptr.Bool(true),
				Encrypted:           ptr.Bool(true),
			},
		})
		number++
	}
	return blockDevices
}

func createInstanceInput(req *CreateRequest, blockDevices []types.BlockDeviceMapping) *ec2.RunInstancesInput {
	return &ec2.RunInstancesInput{
		ImageId:             &req.ImageID,
		InstanceType:        types.InstanceType(req.InstanceType),
		MinCount:            &req.Min,
		MaxCount:            &req.Max,
		KeyName:             &req.KeyName,
		BlockDeviceMappings: blockDevices,
		Monitoring: &types.RunInstancesMonitoringEnabled{
			Enabled: ptr.Bool(true),
		},
		NetworkInterfaces: []types.InstanceNetworkInterfaceSpecification{
			{
				AssociatePublicIpAddress: ptr.Bool(true),
				Groups:                   req.SecurityGroupIDS,
				DeviceIndex:              ptr.Int32(0),
				SubnetId:                 &req.SubnetID,
			},
		},
	}
}

func output(instances []types.Instance) (states []State) {
	for _, value := range instances {
		states = append(states, State{
			ID:   value.InstanceId,
			Name: value.State.Name,
		})
	}
	return states
}

func createDBInstanceInput(instances []types.Instance, blockDevices []types.BlockDeviceMapping) (ec2s []domain.EC2) {
	for _, value := range instances {

		bs, _ := json.Marshal(blockDevices)
		block_devices := string(bs)

		ec2s = append(ec2s, domain.EC2{
			InstanceID:   value.InstanceId,
			ImageID:      value.ImageId,
			InstanceType: (*string)(&value.InstanceType),
			BlockDevices: &block_devices,
			IPv4:         value.PublicIpAddress,
			State:        (*string)(&value.State.Name),
		})
	}
	return
}

//Network Card attach ederken kullanılmaz.
//Eğer yeni oluşturacak olsaydım kullanılırdı.
// PrivateDnsNameOptions: &types.PrivateDnsNameOptionsRequest{
// 	HostnameType: types.HostnameType(req.InstanceName),
// },
