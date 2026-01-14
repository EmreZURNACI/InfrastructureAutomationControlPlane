package ebs

import "github.com/EmreZURNACI/InfrastructureAutomationControlPlane/service/ec2/ebs"

type EbsService struct {
	AttachService         *ebs.AttachService
	CreateSnapshotService *ebs.CreateSnapshotService
	CreateService         *ebs.CreateService
	DeleteSnapshotService *ebs.DeleteSnapshotService
	DeleteService         *ebs.DeleteService
	DetachService         *ebs.DetachService
	DetailService         *ebs.DetailService
	EditService           *ebs.EditService
	ListService           *ebs.ListService
}
