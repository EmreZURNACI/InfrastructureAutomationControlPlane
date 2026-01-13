package server

import (
	"net/http"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers/ec2/network"
	networkSrv "github.com/EmreZURNACI/InfrastructureAutomationControlPlane/service/ec2/network"
)

func (h *routesHandler) StartNetworks() {

	listVPCService := networkSrv.NewListVPC(h.Dependencies.EC2Client)
	listSubnetService := networkSrv.NewListSubnet(h.Dependencies.EC2Client)

	networkController := network.NetworkService{
		ListVPCService:    listVPCService,
		ListSubnetService: listSubnetService,
	}

	route := h.App.Group("/network")

	route.Add(http.MethodGet, "/vpc", networkController.VPC)
	route.Add(http.MethodGet, "/subnet/:id", networkController.SUBNET)

}
