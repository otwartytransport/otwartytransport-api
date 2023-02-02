package status

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Running bool `json:"running"`
}

func SetupRoutes(router fiber.Router) {
	router.Get("/status", status)
}

// Status godoc
// @Summary Return API status
// @Produce json
// @Success 200 {object} Response
// @Router  /status [get]
func status(ctx *fiber.Ctx) error {
	return ctx.JSON(Response{
		Running: true,
	})
}
