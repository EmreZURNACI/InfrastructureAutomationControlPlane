package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type RestartRequest struct {
	InstanceIDs []string `json:"instance_ids" validate:"required"`
}
type RestartResponse struct {
}
type RestartHandler struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewRestartHandler(client ports.InstanceClient, tp ports.Tracer) *RestartHandler {
	return &RestartHandler{client: client, tp: tp}
}

func (h *RestartHandler) Handle(ctx context.Context, req *RestartRequest) (*RestartResponse, error) {

	ctx, span := h.tp.Start(ctx, "Restart Instances")
	defer span.End()

	_, err := h.client.RestartInstances(ctx, &ec2.RebootInstancesInput{
		InstanceIds: req.InstanceIDs,
	})
	if err != nil {
		return nil, err
	}

	return &RestartResponse{}, nil
}
