package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/log"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type StartRequest struct {
	InstanceIDs []string `json:"instance_ids" validate:"required"`
}
type StartResponse struct {
	StartedInstances []StartedInstance `json:"started_instances"`
}
type StartHandler struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewStartHandler(client ports.InstanceClient, tp ports.Tracer) *StartHandler {
	return &StartHandler{client: client, tp: tp}
}

func (h *StartHandler) Handle(ctx context.Context, req *StartRequest) (*StartResponse, error) {

	ctx, span := h.tp.Start(ctx, "Start Instances")
	defer span.End()

	out, err := h.client.StartInstances(ctx, &ec2.StartInstancesInput{
		InstanceIds: req.InstanceIDs,
	})
	if err != nil {
		log.Logger.Info(err.Error())
		return nil, err
	}
	var instances []StartedInstance
	for _, value := range out.StartingInstances {
		instances = append(instances, StartedInstance{
			ID: value.InstanceId,
		})
	}
	return &StartResponse{
		StartedInstances: instances,
	}, nil
}

type StartedInstance struct {
	ID *string `json:"id"`
}
