package ebs

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/ebs"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"
	"github.com/gofiber/fiber/v2"
)

func (srv *EbsService) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	req := ebs.DeleteRequest{
		ID: id,
	}

	if err := controllers.Validate(req); err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	_, err := srv.DeleteService.Execute(c.UserContext(), &req)
	if err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	return c.JSON(controllers.SuccessResponse(200, "block device deleted succesfully", nil))
}
