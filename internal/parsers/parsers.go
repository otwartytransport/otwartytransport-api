package parsers

import (
	"github.com/gofiber/fiber/v2"
)

func BodyParser[K interface{}](ctx *fiber.Ctx, k *K) error {
	if err := ctx.BodyParser(k); err != nil {
		return err
	}

	return nil
}

func QueryParser[K interface{}](ctx *fiber.Ctx, k *K) error {
	if err := ctx.QueryParser(k); err != nil {
		return err
	}

	return nil
}
