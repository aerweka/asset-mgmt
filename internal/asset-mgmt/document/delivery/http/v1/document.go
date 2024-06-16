package v1

import (
	"reflect"
	"strings"

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
	queries := ctx.Queries()
	if len(queries) > 0 {
		documentStruct := reflect.TypeOf(model.Document{})
		found := false
		for i := 0; i < documentStruct.NumField(); i++ {
			field := documentStruct.Field(i)
			tag := field.Tag.Get("json")
			if _, ok := queries[tag]; ok {
				found = true
			}
		}
		if !found {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error retrieving documents", "data": "Field isn't there"})
		}
	}

	var (
		documents []*model.Document
		err       error
	)

	documents, err = h.documentUC.Index(ctx, queries)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving documents", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Documents retrieved successfully", "data": documents})
}

func (h *documentHandler) CreateDocument(ctx *fiber.Ctx) error {
	document := new(model.DocumentRequest)
	err := ctx.BodyParser(document)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing body", "data": err.Error()})
	}

	res, err := h.documentUC.CreateDocument(ctx, document)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error creating document", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Document created successfully", "data": res})
}

func (h *documentHandler) GetDocument(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID", "data": err.Error()})
	}

	document, err := h.documentUC.GetDocument(ctx, parsedId)
	if err != nil {
		httpCode := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			httpCode = fiber.StatusNotFound
		}
		return ctx.Status(httpCode).JSON(fiber.Map{"message": "Error retrieving document", "data": err.Error()})
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

	updatedDocument, err := h.documentUC.UpdateDocument(ctx, document, parsedId)
	if err != nil {
		httpCode := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			httpCode = fiber.StatusNotFound
		}
		return ctx.Status(httpCode).JSON(fiber.Map{"message": "Error updating document", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Document updated successfully", "data": updatedDocument})
}
