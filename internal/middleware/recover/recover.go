package recover

import (
	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

func New() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				var ok bool
				if err, ok = r.(error); !ok {
					c.Status(500)
					err = c.JSON(Error{
						Type:    "UnknownError",
						Message: "Unknown error",
					})
				}
			}
		}()

		return c.Next()
	}
}
