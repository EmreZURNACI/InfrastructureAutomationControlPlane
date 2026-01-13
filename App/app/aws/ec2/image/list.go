package image

import (
	"context"
	"errors"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/ptr"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ListRequest struct {
}
type ListResponse struct {
	Message *string `json:"message"`
	Images  []Image `json:"images"`
}
type ListHandler struct {
	client ports.ImageClient
}

func NewListHandler(client ports.ImageClient) *ListHandler {
	return &ListHandler{
		client: client,
	}
}

func (h *ListHandler) Handle(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	out, err := h.client.DescribeImages(ctx, nil)
	if err != nil {
		return nil, err
	}

	if len(out.Images) == 0 {
		return nil, errors.New("no item found")
	}

	var images []Image
	for _, value := range out.Images {
		images = append(images, Image{
			ID:             value.ImageId,
			Name:           value.Name,
			Description:    value.Description,
			State:          value.State,
			Hypervisor:     value.Hypervisor,
			Architecture:   value.Architecture,
			Platform:       value.Platform,
			OwnerID:        value.OwnerId,
			Public:         value.Public,
			CreationDate:   value.CreationDate,
			RootDeviceType: value.RootDeviceName,
			Virtualization: value.VirtualizationType,
		})
	}

	return &ListResponse{
		Message: ptr.String("images listed successfully"),
		Images:  images,
	}, nil

}

type Image struct {
	ID             *string                  `json:"id"`
	Name           *string                  `json:"name"`
	Description    *string                  `json:"description,omitempty"`
	State          types.ImageState         `json:"state"`
	Architecture   types.ArchitectureValues `json:"architecture"`
	Hypervisor     types.HypervisorType     `json:"hypervisor"`
	Platform       types.PlatformValues     `json:"platform,omitempty"`
	OwnerID        *string                  `json:"owner_id"`
	Public         *bool                    `json:"public"`
	CreationDate   *string                  `json:"creation_date"`
	RootDeviceType *string                  `json:"root_device_type"`
	Virtualization types.VirtualizationType `json:"virtualization_type"`
}
