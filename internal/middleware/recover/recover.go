package recover

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otwartytransport/otwartytransport-api/internal/errors"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				var ok bool
				if err, ok = r.(error); !ok {
					c.Status(500)
					err = c.JSON(errors.Unknown)
				}
			}
		}()

		return c.Next()
	}
}
