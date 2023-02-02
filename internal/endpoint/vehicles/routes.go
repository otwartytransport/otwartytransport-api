package vehicles

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	vehicles := router.Group("/vehicles")
	SetupRoutes(vehicles)
}