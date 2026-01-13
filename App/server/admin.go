package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *routesHandler) StartAdmin() {
	route := h.App.Group("/admin")
	route.Add(http.MethodGet, "/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Hello To Admin",
		})
	})

	// route.Add(http.MethodGet, "/users", controllers.GetUsers)
	// route.Add(http.MethodGet, "/computers", controllers.GetComputers)
}
