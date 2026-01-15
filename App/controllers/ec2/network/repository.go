package network

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/service/ec2/network"
)

type NetworkService struct {
	ListVPCService    *network.ListVPC
	ListSubnetService *network.ListSubnet
}
