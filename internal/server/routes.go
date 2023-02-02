package server

import (
	"github.com/otwartytransport/otwartytransport-api/internal/endpoint/status"
	"github.com/otwartytransport/otwartytransport-api/internal/endpoint/vehicles"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	status.SetupRoutes(v1)
	vehicles.SetupRoutes(v1)
}