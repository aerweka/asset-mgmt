package v1

import (
	"strings"

	mainasset "asset-management.com/internal/asset-mgmt/main_asset"
	"asset-management.com/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type mainAssetHandler struct {
	mainAssetUC mainasset.Usecase
}

func NewMainAssetHandler(uc mainasset.Usecase) mainasset.Handlers {
	return &mainAssetHandler{
		mainAssetUC: uc,
	}
}

func (h *mainAssetHandler) Index(ctx *fiber.Ctx) error {
	var (
		pbb []*model.PBB
		err error
	)

	pbb, err = h.mainAssetUC.Index(ctx)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving pbbs", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Documents retrieved successfully", "data": pbb})
}

func (h *mainAssetHandler) CreateMainAsset(ctx *fiber.Ctx) error {
	pbb := new(model.PBBRequest)
	err := ctx.BodyParser(pbb)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing body", "data": err.Error()})
	}

	res, err := h.mainAssetUC.CreateMainAsset(ctx, pbb)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error creating pbb", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "PBB created successfully", "data": res})
}

func (h *mainAssetHandler) GetMainAsset(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID", "data": err.Error()})
	}

	pbb, err := h.mainAssetUC.GetMainAsset(ctx, parsedId)
	if err != nil {
		httpCode := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			httpCode = fiber.StatusNotFound
		}
		return ctx.Status(httpCode).JSON(fiber.Map{"message": "Error retrieving pbb", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "PBB retrieved successfully", "data": pbb})
}

func (h *mainAssetHandler) UpdateMainAsset(ctx *fiber.Ctx) error {
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

	updatedPbb, err := h.mainAssetUC.UpdateMainAsset(ctx, pbb, parsedId)
	if err != nil {
		httpCode := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			httpCode = fiber.StatusNotFound
		}
		return ctx.Status(httpCode).JSON(fiber.Map{"message": "Error updating pbb", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "PBB updated successfully", "data": updatedPbb})
}
