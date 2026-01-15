package instance

import "github.com/EmreZURNACI/InfrastructureAutomationControlPlane/service/ec2/instance"

type InstanceService struct {
	CreateService    *instance.CreateService
	DetailService    *instance.DetailService
	EditService      *instance.EditService
	ListService      *instance.ListService
	ListTypesService *instance.ListTypesService
	RestartService   *instance.RestartService
	StartService     *instance.StartService
	StopService      *instance.StopService
	TerminateService *instance.TerminateService
}
