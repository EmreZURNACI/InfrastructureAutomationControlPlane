package auth

import (
	"slices"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/controllers"
	"github.com/gofiber/fiber/v2"
)

func Authorization(requiredPermission string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		perms := c.Locals("X-USER-PERMISSIONS").([]string)

		if !slices.Contains(perms, requiredPermission) {
			return c.JSON(controllers.FailureResponse(fiber.ErrForbidden.Code, "Forbidden", fiber.ErrForbidden.Message))

		}

		return c.Next()
	}
}
