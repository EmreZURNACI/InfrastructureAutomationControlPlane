package ebs

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/ebs"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"
	"github.com/gofiber/fiber/v2"
)

func (srv *EbsService) Edit(c *fiber.Ctx) error {

	id := c.Params("id")

	if id == "" {
		return c.JSON(controllers.FailureResponse(400, "id required", "id required"))
	}

	var req ebs.EditRequest

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	req.ID = id

	if err := controllers.Validate(req); err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	res, err := srv.EditService.Execute(c.UserContext(), &req)
	if err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	return c.JSON(controllers.SuccessResponse(200, "block device edited successfully", res.State))
}
