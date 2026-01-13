package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/log"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type StopRequest struct {
	InstanceIDs []string `json:"instance_ids" validate:"required"`
}
type StopResponse struct {
	StoppedInstances []StoppedInstance `json:"stopped_instances"`
}
type StopHandler struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewStopHandler(client ports.InstanceClient, tp ports.Tracer) *StopHandler {
	return &StopHandler{client: client, tp: tp}
}

func (h *StopHandler) Handle(ctx context.Context, req *StopRequest) (*StopResponse, error) {

	ctx, span := h.tp.Start(ctx, "Stop Instances")
	defer span.End()

	out, err := h.client.StopInstances(ctx, &ec2.StopInstancesInput{
		InstanceIds: req.InstanceIDs,
	})
	if err != nil {
		log.Logger.Info(err.Error())
		return nil, err
	}

	var instances []StoppedInstance
	for _, value := range out.StoppingInstances {
		instances = append(instances, StoppedInstance{
			ID: value.InstanceId,
		})
	}
	return &StopResponse{
		StoppedInstances: instances,
	}, nil
}

type StoppedInstance struct {
	ID *string `json:"id"`
}
