package network

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/network"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"
	"github.com/gofiber/fiber/v2"
)

func (srv *NetworkService) SUBNET(c *fiber.Ctx) error {

	id := c.Params("id")

	if id == "" {
		return c.JSON(controllers.FailureResponse(fiber.StatusBadRequest, "id is required", "id is required"))
	}
	req := network.ListSubnetRequest{
		VpcID: id,
	}

	if err := controllers.Validate(req); err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	res, err := srv.ListSubnetService.Execute(c.UserContext(), &req)
	if err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	return c.JSON(controllers.SuccessResponse(200, *res.Message, res.Subnets))
}
