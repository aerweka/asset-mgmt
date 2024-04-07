package router

import (
	"asset-management.com/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func NewDocumentRoutes(api fiber.Router) {
	document := api.Group("/documents")

	document.Get("/", handler.Index)
	document.Get("/:id", handler.GetDocument)
	document.Post("/", handler.CreateDocument)
	document.Put("/:id", handler.UpdateDocument)
}
