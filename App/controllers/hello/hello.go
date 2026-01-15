package hello

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/controllers"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {

	return c.JSON(controllers.SuccessResponse(200, "Hello from InfrastructureAutomationControlPlane", nil))
}
