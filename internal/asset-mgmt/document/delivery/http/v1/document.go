package v1

import (
	"asset-management.com/internal/asset-mgmt/document"
	"asset-management.com/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type documentHandler struct {
	documentUC document.Usecase
}

func NewDocumentHandler(uc document.Usecase) document.Handlers {
	return &documentHandler{
		documentUC: uc,
	}
}

func (h *documentHandler) Index(ctx *fiber.Ctx) error {
	var documents []model.Document
	documents, err := h.documentUC.Index(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving documents", "data": err.Error()})
	}

	if len(documents) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No documents found", "data": nil})
	}

	return ctx.JSON(fiber.Map{"message": "Documents retrieved successfully", "data": documents})
}

func (h *documentHandler) CreateDocument(ctx *fiber.Ctx) error {
	document := new(model.Document)
	err := ctx.BodyParser(document)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing body", "data": err.Error()})
	}

	err = h.documentUC.CreateDocument(ctx, document)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error creating document", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Document created successfully", "data": document})
}

func (h *documentHandler) GetDocument(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID", "data": err.Error()})
	}

	document, err := h.documentUC.GetDocument(ctx, parsedId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving document", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Document retrieved successfully", "data": document})
}

func (h *documentHandler) UpdateDocument(ctx *fiber.Ctx) error {
	document := new(model.Document)
	err := ctx.BodyParser(document)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing body", "data": err.Error()})
	}

	id := ctx.Params("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID", "data": err.Error()})
	}

	err = h.documentUC.UpdateDocument(ctx, parsedId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error updating document", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Document updated successfully", "data": document})
}
