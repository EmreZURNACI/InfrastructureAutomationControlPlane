package network

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/network"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type ListSubnet struct {
	client ports.NetworkClient
}

func NewListSubnet(client ports.NetworkClient) *ListSubnet {
	return &ListSubnet{client: client}
}

func (s *ListSubnet) Execute(ctx context.Context, req *network.ListSubnetRequest) (*network.ListSubnetResponse, error) {

	handler := network.NewListSubnetHandler(s.client)
	return handler.Handle(ctx, req)
}
