package instance

import (
	"context"
	"errors"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/ptr"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ListInstanceTypeRequest struct {
}
type ListInstanceTypeResponse struct {
	InstanceType []InstanceType `json:"instance_type"`
	Message      *string        `json:"message"`
}
type ListInstanceTypeHandler struct {
	client ports.InstanceClient
}

func NewListInstanceTypeHandler(client ports.InstanceClient) *ListInstanceTypeHandler {
	return &ListInstanceTypeHandler{
		client: client,
	}
}
func (h *ListInstanceTypeHandler) Handle(ctx context.Context, req *ListInstanceTypeRequest) (*ListInstanceTypeResponse, error) {
	out, err := h.client.ListInstanceTypes(ctx, &ec2.DescribeInstanceTypesInput{})
	if err != nil {
		return nil, err
	}

	if len(out.InstanceTypes) == 0 {
		return nil, errors.New("item not found")
	}

	var types []InstanceType
	for _, value := range out.InstanceTypes {
		types = append(types, InstanceType{
			Name:               value.InstanceType,
			CPU:                value.VCpuInfo,
			Memory:             value.MemoryInfo,
			FreeTierEligible:   value.FreeTierEligible,
			BareMetal:          value.BareMetal,
			Hypervisor:         value.Hypervisor,
			NetworkPerformance: value.NetworkInfo.NetworkPerformance,
		})
	}

	return &ListInstanceTypeResponse{
		Message:      ptr.String("instance types listed successfully"),
		InstanceType: types,
	}, nil
}

type InstanceType struct {
	Name               types.InstanceType           `json:"name"`
	CPU                *types.VCpuInfo              `json:"cpu"`
	Memory             *types.MemoryInfo            `json:"memory"`
	FreeTierEligible   *bool                        `json:"free_tier_eligible"`
	BareMetal          *bool                        `json:"bare_metal"`
	Hypervisor         types.InstanceTypeHypervisor `json:"hypervisor"`
	NetworkPerformance *string                      `json:"network_performance"`
}
