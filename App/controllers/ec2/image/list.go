package image

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/app/aws/ec2/image"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"

	"github.com/gofiber/fiber/v2"
)

func (srv *ImageService) List(c *fiber.Ctx) error {

	res, err := srv.ListService.Execute(c.UserContext(), &image.ListRequest{})
	if err != nil {
		return c.JSON(controllers.FailureResponse(400, err.Error(), err.Error()))
	}

	return c.JSON(controllers.SuccessResponse(200, *res.Message, res.Images))
}
