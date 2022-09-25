package server

import (
	"github.com/otwartytransport/otwartytransport-api/internal/endpoint/status"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	statusRoutes(app)
}

func statusRoutes(app *fiber.App) {
	app.Get("/status", status.Status)
}
