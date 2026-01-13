package server

import (
	"net/http"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/adaptor/tracer"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers/ec2/ebs"
	ebsSrv "github.com/EmreZURNACI/InfrastructureAutomationControlPlane/service/ec2/ebs"
)

func (h *routesHandler) StartVolumes() {

	tp := tracer.NewTracer(h.Dependencies.Tp)

	listService := ebsSrv.NewListService(h.Dependencies.EC2Client, tp)
	createService := ebsSrv.NewCreateService(h.Dependencies.EC2Client, tp)

	detailService := ebsSrv.NewDetailService(h.Dependencies.EC2Client, tp)
	deleteService := ebsSrv.NewDeleteService(h.Dependencies.EC2Client, tp)
	editService := ebsSrv.NewEditService(h.Dependencies.EC2Client, tp)

	attachService := ebsSrv.NewAttachService(h.Dependencies.EC2Client, tp)
	detachService := ebsSrv.NewDetachService(h.Dependencies.EC2Client, tp)

	createSnapshotService := ebsSrv.NewCreateSnapshotService(h.Dependencies.EC2Client, tp)
	deleteSnapshotService := ebsSrv.NewDeleteSnapshotService(h.Dependencies.EC2Client, tp)

	ebsController := ebs.EbsService{
		ListService:           listService,
		CreateService:         createService,
		DetailService:         detailService,
		DeleteService:         deleteService,
		EditService:           editService,
		AttachService:         attachService,
		DetachService:         detachService,
		CreateSnapshotService: createSnapshotService,
		DeleteSnapshotService: deleteSnapshotService,
	}

	route := h.App.Group("/ebs")

	route.Add(http.MethodGet, "/", ebsController.List)
	route.Add(http.MethodPost, "/", ebsController.Create)

	route.Add(http.MethodGet, "/:id", ebsController.Detail)
	route.Add(http.MethodDelete, "/:id", ebsController.Delete)
	route.Add(http.MethodPut, "/:id", ebsController.Edit)

	route.Add(http.MethodPatch, "/attach", ebsController.Attach)
	route.Add(http.MethodPatch, "/detach/:id", ebsController.Detach)

	route.Add(http.MethodPost, "/snapshot", ebsController.CreateSnapshot)
	route.Add(http.MethodDelete, "/snapshot/:id", ebsController.DeleteSnapshot)
}
