package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/config"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Start(dep Dependencies) {

	app := fiber.New(fiber.Config{
		AppName: "MyService",

		ServerHeader:          "EmreZURNACI - InfrastructureAutomationControlPlane",
		DisableStartupMessage: true,

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	})

	app.Use(race_limiter)
	app.Use(cors_handler)
	app.Use(otelfiber.Middleware())

	app.Add(http.MethodGet, "/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	app.Use(InternalOnly)

	routes := NewRoutesHandler(app, dep)
	routes.StartHello()
	routes.StartLiveness()
	routes.StartAdmin()
	routes.StartInstance()
	routes.StartVolumes()
	routes.StartNetworks()
	routes.StartImages()
	routes.StartKeys()

	log.Logger.Info(fmt.Sprintf("HTTP server started in %s:%d", config.AppConfig.ServerConfig.Host, config.AppConfig.ServerConfig.Port))
	if err := app.Listen(fmt.Sprintf("%s:%d", config.AppConfig.ServerConfig.Host, config.AppConfig.ServerConfig.Port)); err != nil {
		log.Logger.Error("HTTP server failed")
	}

	gracefulShutdown(app)

}
