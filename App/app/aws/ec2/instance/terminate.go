package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/log"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type TerminateRequest struct {
	InstanceIDs []string `json:"instance_ids" validate:"required"`
}
type TerminateResponse struct {
	TerminatedInstances []TerminatedInstance `json:"terminated_instances"`
}
type TerminateHandler struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewTerminateHandler(client ports.InstanceClient, tp ports.Tracer) *TerminateHandler {
	return &TerminateHandler{client: client, tp: tp}
}

func (h *TerminateHandler) Handle(ctx context.Context, req *TerminateRequest) (*TerminateResponse, error) {

	ctx, span := h.tp.Start(ctx, "Terminate Instances")
	defer span.End()

	out, err := h.client.TerminateInstances(ctx, &ec2.TerminateInstancesInput{
		InstanceIds: req.InstanceIDs,
	})
	if err != nil {
		log.Logger.Info(err.Error())
		return nil, err
	}
	var instances []TerminatedInstance
	for _, value := range out.TerminatingInstances {
		instances = append(instances, TerminatedInstance{
			ID: value.InstanceId,
		})
	}
	return &TerminateResponse{
		TerminatedInstances: instances,
	}, nil
}

type TerminatedInstance struct {
	ID *string `json:"id"`
}
