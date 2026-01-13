package server

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/log"
	"github.com/gofiber/fiber/v2"
)

func gracefulShutdown(app *fiber.App) {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	log.Logger.Info("Shutting down server...")

	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		log.Logger.Error("Error during server shutdown")
	}

	log.Logger.Info("Server gracefully stopped")
}
