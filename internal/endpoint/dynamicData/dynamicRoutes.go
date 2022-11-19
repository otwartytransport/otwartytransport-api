package dynamicData

import (
	"github.com/gofiber/fiber/v2"
)

var DynamicDataQuadtreeName = "dynamicData"

func SetupDynamicRoutes(router fiber.Router) {
	dynamicData := router.Group("/dynamicData")
	dynamicData.Get("/findNearest", FindNearestVehicle)
	dynamicData.Get("/findInBounds", FindVehicleInBounds)

	dynamicData.Post("/uploadData", uploadVehicle)
}

const (
	Tram string = "tram"
	Bus         = "bus"
)
