package server

import (
	documentHttp "asset-management.com/internal/asset-mgmt/document/delivery/http/v1"
	"asset-management.com/internal/asset-mgmt/document/repository"
	"asset-management.com/internal/asset-mgmt/document/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func NewHandler(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api", logger.New())

	app.Get("health-check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Server is running"})
	})

	// Repo initialization -> START
	documentRepo := repository.NewDocumentRepository(db)
	// Repo initialization -> END

	// Usecase initialization -> START
	documentUsecase := usecase.NewDocumentUsecase(usecase.DocumentUCParam{
		DocumentRepo: documentRepo,
	})
	// Usecase initialization -> END

	// Handler initialization -> START
	documentHandler := documentHttp.NewDocumentHandler(documentUsecase)
	// Handler initialization -> END

	// Route mapping
	documentHttp.NewDocumentRoutes(api, documentHandler)
}
