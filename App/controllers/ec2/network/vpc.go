package network

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/network"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"
	"github.com/gofiber/fiber/v2"
)

func (srv *NetworkService) VPC(c *fiber.Ctx) error {
	var req network.ListVPCRequest

	res, err := srv.ListVPCService.Execute(c.UserContext(), &req)
	if err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	return c.JSON(controllers.SuccessResponse(200, *res.Message, res.Vpcs))
}
