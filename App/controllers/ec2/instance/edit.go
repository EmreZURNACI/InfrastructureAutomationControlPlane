package instance

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/instance"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"
	"github.com/gofiber/fiber/v2"
)

func (srv *InstanceService) Edit(c *fiber.Ctx) error {

	id := c.Params("id")

	if id == "" {
		return c.JSON(controllers.FailureResponse(fiber.StatusBadRequest, "id is required", "id is required"))
	}

	var req instance.EditRequest

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(controllers.FailureResponse(fiber.StatusBadRequest, err.Error(), err.Error()))
	}

	req.ID = id

	if err := controllers.Validate(req); err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	_, err := srv.EditService.Execute(c.UserContext(), &req)
	if err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	return c.JSON(controllers.SuccessResponse(200, "instance edited successfully", nil))
}
