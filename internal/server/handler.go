package server

import (
	documentHttp "asset-management.com/internal/asset-mgmt/document/delivery/http/v1"
	pbbHttp "asset-management.com/internal/asset-mgmt/pbb/delivery/http/v1"

	documentRepo "asset-management.com/internal/asset-mgmt/document/repository"
	pbbRepo "asset-management.com/internal/asset-mgmt/pbb/repository"

	documentUc "asset-management.com/internal/asset-mgmt/document/usecase"
	pbbUc "asset-management.com/internal/asset-mgmt/pbb/usecase"

	"asset-management.com/pkg/cloudinary"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func NewHandler(app *fiber.App, db *gorm.DB, cloudinary *cloudinary.Cloudinary) {
	api := app.Group("/api", logger.New())

	app.Get("health-check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Server is running"})
	})

	// Repo initialization -> START
	documentRepo := documentRepo.NewDocumentRepository(db)
	pbbRepo := pbbRepo.NewPbbRepository(db)
	// Repo initialization -> END

	// Usecase initialization -> START
	documentUsecase := documentUc.NewDocumentUsecase(&documentUc.DocumentUCParam{
		DocumentRepo: documentRepo,
		Cloudinary:   cloudinary,
	})
	pbbUsecase := pbbUc.NewPbbUsecase(&pbbUc.PBBUCParam{
		PBBRepository: pbbRepo,
		Cloudinary:    cloudinary,
	})

	// Usecase initialization -> END

	// Handler initialization -> START
	documentHandler := documentHttp.NewDocumentHandler(documentUsecase)
	pbbHandler := pbbHttp.NewDocumentHandler(pbbUsecase)
	// Handler initialization -> END

	// Route mapping
	documentHttp.NewDocumentRoutes(api, documentHandler)
	pbbHttp.NewDocumentRoutes(api, pbbHandler)
}
