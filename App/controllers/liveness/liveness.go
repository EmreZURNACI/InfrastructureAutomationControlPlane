package liveness

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"
	"github.com/gofiber/fiber/v2"
)

func Liveness(c *fiber.Ctx) error {

	return c.JSON(controllers.SuccessResponse(200, "OK", nil))
}
