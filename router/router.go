package router

import (
	documentRouter "asset-management.com/internal/routes/document"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewRoute(app *fiber.App) {
	api := app.Group("/api", logger.New())

	app.Get("health-check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Server is running"})
	})

	documentRouter.NewDocumentRoutes(api)
}
