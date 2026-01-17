package instance

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/instance"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type DetailService struct {
	client ports.InstanceClient
	tp     ports.Tracer
}

func NewDetailService(client ports.InstanceClient, tp ports.Tracer) *DetailService {
	return &DetailService{client: client, tp: tp}
}

func (s *DetailService) Execute(ctx context.Context, req *instance.DetailRequest) (*instance.DetailResponse, error) {

	handler := instance.NewDetailHandler(s.client, s.tp)
	return handler.Handle(ctx, req)
}
