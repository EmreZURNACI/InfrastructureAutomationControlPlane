package server

import (
	"time"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var cors_handler = cors.New(cors.Config{
	AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
	AllowOriginsFunc: func(origin string) bool {
		return true
	},
	AllowCredentials: true,
	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
})

func InternalOnly(c *fiber.Ctx) error {

	if c.Get("X-INTERNAL-SECRET") != config.AppConfig.ServerConfig.ProxySecret {
		return fiber.ErrForbidden
	}
	return c.Next()
}

var race_limiter = limiter.New(limiter.Config{
	Max:        20,
	Expiration: 30 * time.Second,
	LimitReached: func(c *fiber.Ctx) error {
		return c.SendFile("../html/toofast.html")
	},
})
