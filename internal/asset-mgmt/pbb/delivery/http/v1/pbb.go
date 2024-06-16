package v1

import (
	"strings"

	"asset-management.com/internal/asset-mgmt/pbb"
	"asset-management.com/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type pbbHandler struct {
	pbbUC pbb.Usecase
}

func NewPbbHandler(uc pbb.Usecase) pbb.Handlers {
	return &pbbHandler{
		pbbUC: uc,
	}
}

func (h *pbbHandler) Index(ctx *fiber.Ctx) error {
	var (
		pbb []*model.PBB
		err error
	)

	pbb, err = h.pbbUC.Index(ctx)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving pbbs", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Documents retrieved successfully", "data": pbb})
}

func (h *pbbHandler) CreatePBB(ctx *fiber.Ctx) error {
	pbb := new(model.PBBRequest)
	err := ctx.BodyParser(pbb)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing body", "data": err.Error()})
	}

	res, err := h.pbbUC.CreatePBB(ctx, pbb)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error creating pbb", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "PBB created successfully", "data": res})
}

func (h *pbbHandler) GetPBB(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID", "data": err.Error()})
	}

	pbb, err := h.pbbUC.GetPBB(ctx, parsedId)
	if err != nil {
		httpCode := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			httpCode = fiber.StatusNotFound
		}
		return ctx.Status(httpCode).JSON(fiber.Map{"message": "Error retrieving pbb", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "PBB retrieved successfully", "data": pbb})
}

func (h *pbbHandler) UpdatePBB(ctx *fiber.Ctx) error {
	pbb := new(model.PBB)
	err := ctx.BodyParser(pbb)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing body", "data": err.Error()})
	}

	id := ctx.Params("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID", "data": err.Error()})
	}

	updatedPbb, err := h.pbbUC.UpdatePBB(ctx, pbb, parsedId)
	if err != nil {
		httpCode := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			httpCode = fiber.StatusNotFound
		}
		return ctx.Status(httpCode).JSON(fiber.Map{"message": "Error updating pbb", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "PBB updated successfully", "data": updatedPbb})
}
