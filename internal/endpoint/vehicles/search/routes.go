package search

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	search := router.Group("/search")
	// search.Get("/findNearest", FindNearestVehicle)
	// search.Get("/findKNearest", FindKNearestVehicle)
	// search.Get("/findInBounds", FindVehicleInBounds)

	// dynamicVehicles.Post("/uploadData", uploadVehicle)
}