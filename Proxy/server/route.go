package server

import (
	"fmt"
	"time"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/controllers"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/controllers/auth"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/config"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start(dep Dependencies) {

	proxy := fiber.New(fiber.Config{

		AppName:               config.AppConfig.ServerConfig.AppName,
		ServerHeader:          config.AppConfig.ServerConfig.Header,
		DisableStartupMessage: true,
		ReadTimeout:           10 * time.Second,
		WriteTimeout:          10 * time.Second,
		IdleTimeout:           15 * time.Second,
	})

	routeHandler := NewRoutesHandler(proxy, dep)

	proxy.Use(CorsHandler)

	routeHandler.StartAuth()

	proxy.Use(auth.Authentication)
	proxy.All("/api/v1/*", auth.DynamicAuthorization(), controllers.Forward())

	log.Logger.Info(fmt.Sprintf("HTTP server started in %s:%d", config.AppConfig.ServerConfig.Host, config.AppConfig.ServerConfig.Port))
	if err := proxy.Listen(fmt.Sprintf("%s:%d", config.AppConfig.ServerConfig.Host, config.AppConfig.ServerConfig.Port)); err != nil {
		log.Logger.Error("HTTP server failed")
	}

	gracefulShutdown(proxy)

}

var CorsHandler = cors.New(cors.Config{
	AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
	AllowOriginsFunc: func(origin string) bool {
		return true
	},
	AllowCredentials: true,
	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
})
