package recover

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otwartytransport/otwartytransport-api/internal/errors"
)

func New() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				var ok bool
				if err, ok = r.(error); !ok {
					err = ctx.
						Status(fiber.StatusInternalServerError).
						JSON(errors.Unknown)
				}
			}
		}()

		return ctx.Next()
	}
}
