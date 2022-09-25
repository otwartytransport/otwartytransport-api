package server

import (
	"github.com/otwartytransport/otwartytransport-api/internal/endpoint/status"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	statusRoutes(v1)
}

func statusRoutes(router fiber.Router) {
	router.Get("/status", status.Status)
}
