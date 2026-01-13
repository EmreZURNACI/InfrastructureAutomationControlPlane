package server

import (
	"net/http"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers/ec2/image"
	imgSrv "github.com/EmreZURNACI/InfrastructureAutomationControlPlane/service/ec2/image"
)

func (h *routesHandler) StartImages() {

	listInstanceService := imgSrv.NewListImageService(h.Dependencies.EC2Client)

	imageController := image.ImageService{
		ListService: listInstanceService,
	}

	route := h.App.Group("/image")

	route.Add(http.MethodGet, "/", imageController.List)

}
