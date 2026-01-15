package instance

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/instance"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"
	"github.com/gofiber/fiber/v2"
)

func (srv *InstanceService) Stop(c *fiber.Ctx) error {

	var req instance.StopRequest

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	if err := controllers.Validate(req); err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))

	}

	res, err := srv.StopService.Execute(c.UserContext(), &req)
	if err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	return c.JSON(controllers.SuccessResponse(200, "machines stopped successfully", res.StoppedInstances))
}
