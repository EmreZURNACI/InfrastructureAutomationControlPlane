package network

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/network"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
)

type ListVPC struct {
	client ports.NetworkClient
}

func NewListVPC(client ports.NetworkClient) *ListVPC {
	return &ListVPC{client: client}
}

func (s *ListVPC) Execute(ctx context.Context, req *network.ListVPCRequest) (*network.ListVPCResponse, error) {

	handler := network.NewListVPCHandler(s.client)
	return handler.Handle(ctx, req)
}
