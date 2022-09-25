package status

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Running bool `json:"running"`
}

func Status(ctx *fiber.Ctx) error {
	return ctx.JSON(Response{
		Running: true,
	})
}
