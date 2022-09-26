package server

import (
	"github.com/gofiber/swagger"
	"github.com/otwartytransport/otwartytransport-api/internal/middleware/recover"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/otwartytransport/otwartytransport-api/docs"
)

// NewServer godoc
// @title         OtwartyTransport API
// @version       0.1
// @contact.email contact@otwartytransport.pl
// @BasePath      /v1
func NewServer() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: errorHandler,
	})
	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	registerRoutes(app)
	return app
}
