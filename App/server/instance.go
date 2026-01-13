package server

import (
	"net/http"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/adaptor/tracer"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers/ec2/instance"
	insSrv "github.com/EmreZURNACI/InfrastructureAutomationControlPlane/service/ec2/instance"
)

func (h *routesHandler) StartInstance() {

	tp := tracer.NewTracer(h.Dependencies.Tp)

	createInstanceService := insSrv.NewCreateService(h.Dependencies.EC2Client, tp)
	listInstanceService := insSrv.NewListService(h.Dependencies.EC2Client, tp)
	listInstanceTypesService := insSrv.NewListTypesService(h.Dependencies.EC2Client)

	startInstanceService := insSrv.NewStartService(h.Dependencies.EC2Client, tp)
	stopInstanceService := insSrv.NewStopService(h.Dependencies.EC2Client, tp)
	restartInstanceService := insSrv.NewRestartService(h.Dependencies.EC2Client, tp)
	terminateInstanceService := insSrv.NewTerminateService(h.Dependencies.EC2Client, tp)

	detailInstanceService := insSrv.NewDetailService(h.Dependencies.EC2Client, tp)
	editInstanceService := insSrv.NewEditService(h.Dependencies.EC2Client, tp)

	ec2Controller := instance.InstanceService{
		CreateService:    createInstanceService,
		ListService:      listInstanceService,
		StartService:     startInstanceService,
		StopService:      stopInstanceService,
		RestartService:   restartInstanceService,
		TerminateService: terminateInstanceService,
		DetailService:    detailInstanceService,
		EditService:      editInstanceService,
		ListTypesService: listInstanceTypesService,
	}

	route := h.App.Group("/instance")

	route.Add(http.MethodPost, "/", ec2Controller.Create)
	route.Add(http.MethodGet, "/", ec2Controller.List)

	route.Add(http.MethodPost, "/start", ec2Controller.Start)
	route.Add(http.MethodPost, "/stop", ec2Controller.Stop)
	route.Add(http.MethodPost, "/restart", ec2Controller.Restart)
	route.Add(http.MethodPost, "/terminate", ec2Controller.Terminate)
	route.Add(http.MethodGet, "/instance-types", ec2Controller.ListTypes)

	route.Add(http.MethodGet, "/:id", ec2Controller.Detail)
	route.Add(http.MethodPut, "/:id", ec2Controller.Edit)
}
