package instance

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/instance"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"

	"github.com/gofiber/fiber/v2"
)

func (srv *InstanceService) Detail(c *fiber.Ctx) error {

	id := c.Params("id")

	req := instance.DetailRequest{ID: id}

	if err := controllers.Validate(req); err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	res, err := srv.DetailService.Execute(c.UserContext(), &req)
	if err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	return c.JSON(controllers.SuccessResponse(200, "instance detailed successfuly", res.Instance))
}
