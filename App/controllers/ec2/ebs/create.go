package ebs

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/ebs"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"
	"github.com/gofiber/fiber/v2"
)

func (srv *EbsService) Create(c *fiber.Ctx) error {
	var req ebs.CreateRequest

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	if err := controllers.Validate(req); err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	res, err := srv.CreateService.Execute(c.UserContext(), &req)
	if err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	return c.JSON(controllers.SuccessResponse(201, "block device created successfully", *res.ID))

}
