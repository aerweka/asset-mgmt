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
		mainAsset []*model.MainAsset
		err       error
	)

	mainAsset, err = h.mainAssetUC.Index(ctx)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving main assets", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Main assets retrieved successfully", "data": mainAsset})
}

func (h *mainAssetHandler) CreateMainAsset(ctx *fiber.Ctx) error {
	mainAsset := new(model.MainAssetRequest)
	err := ctx.BodyParser(mainAsset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing body", "data": err.Error()})
	}

	res, err := h.mainAssetUC.CreateMainAsset(ctx, mainAsset)
	if err != nil {
		httpCode := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			httpCode = fiber.StatusNotFound
		}
		return ctx.Status(httpCode).JSON(fiber.Map{"message": "Error creating main asset", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Main asset created successfully", "data": res})
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
		return ctx.Status(httpCode).JSON(fiber.Map{"message": "Error retrieving main asset", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Main asset retrieved successfully", "data": pbb})
}

func (h *mainAssetHandler) UpdateMainAsset(ctx *fiber.Ctx) error {
	mainAsset := new(model.MainAsset)
	err := ctx.BodyParser(mainAsset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing body", "data": err.Error()})
	}

	id := ctx.Params("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID", "data": err.Error()})
	}

	updatedMainAsset, err := h.mainAssetUC.UpdateMainAsset(ctx, mainAsset, parsedId)
	if err != nil {
		httpCode := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			httpCode = fiber.StatusNotFound
		}
		return ctx.Status(httpCode).JSON(fiber.Map{"message": "Error updating main asset", "data": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Main asset updated successfully", "data": updatedMainAsset})
}
