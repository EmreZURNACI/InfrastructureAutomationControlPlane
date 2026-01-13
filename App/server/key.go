package server

import (
	"net/http"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers/ec2/key"
	keySrv "github.com/EmreZURNACI/InfrastructureAutomationControlPlane/service/ec2/key"
)

func (h *routesHandler) StartKeys() {

	listService := keySrv.NewListService(h.Dependencies.EC2Client)

	keyController := key.KeyService{
		ListService: listService,
	}

	route := h.App.Group("/key")

	route.Add(http.MethodGet, "/", keyController.List)

}
