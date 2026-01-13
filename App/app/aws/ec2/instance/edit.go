package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EditRequest struct {
	ID             string    `json:"id" validate:"required"`
	SecurityGroups *[]string `json:"security_groups,omitempty" validate:"omitempty"`
}
type EditResponse struct {
}
type EditHandler struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewEditHandler(client ports.InstanceClient, tp ports.Tracer) *EditHandler {
	return &EditHandler{client: client, tp: tp}
}

func (h *EditHandler) Handle(ctx context.Context, req *EditRequest) (*EditResponse, error) {

	ctx, span := h.tp.Start(ctx, "Edit Instance")
	defer span.End()

	_, err := h.client.EditInstances(ctx, &ec2.ModifyInstanceAttributeInput{
		InstanceId: &req.ID,
		Groups:     *req.SecurityGroups,
	})
	if err != nil {
		return nil, err
	}

	return &EditResponse{}, nil
}
