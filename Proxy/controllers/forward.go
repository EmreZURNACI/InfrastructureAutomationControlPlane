package controllers

import (
	"strings"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/config"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func Forward() fiber.Handler {

	return func(c *fiber.Ctx) error {
		userID := c.Locals("X-USER-ID").(string)
		perms := c.Locals("X-USER-PERMISSIONS").([]string)

		if c.Get("X-PROXIED-BY") == "gateway" {
			log.Logger.Info("LOOP DETECTED")
			return c.JSON(FailureResponse(fiber.ErrBadGateway.Code, fiber.ErrBadGateway.Message, fiber.ErrBadGateway.Message))

		}

		c.Request().Header.Set("X-PROXIED-BY", "gateway")
		c.Request().Header.Set("X-INTERNAL-SECRET", config.AppConfig.ServerConfig.ProxySecret)
		c.Request().Header.Set("X-USER-ID", userID)
		c.Request().Header.Set("X-USER-PERMISSIONS", strings.Join(perms, ","))

		base := "http://server:2963"

		path := strings.TrimPrefix(c.OriginalURL(), "/api/v1")
		url := base + path

		return proxy.Do(c, url)
	}
}
