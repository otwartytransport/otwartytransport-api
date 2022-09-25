package server

import (
	"github.com/otwartytransport/otwartytransport-api/internal/middleware/recover"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewServer() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	registerRoutes(app)
	return app
}