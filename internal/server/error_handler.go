package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otwartytransport/otwartytransport-api/internal/errors"
)

func errorHandler(ctx *fiber.Ctx, err error) error {
	var (
		apiErr = errors.Unknown
		code   = fiber.StatusInternalServerError
	)

	// Custom status code for fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		apiErr = errors.Error{
			Type:    "FiberError",
			Message: e.Message,
		}
	}

	// Custom status code for our errors
	if e, ok := err.(errors.Error); ok {
		code = e.Code
		apiErr = e
	}

	return ctx.Status(code).JSON(apiErr)
}
