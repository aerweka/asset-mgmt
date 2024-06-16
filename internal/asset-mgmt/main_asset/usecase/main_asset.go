package usecase

import (
	"errors"

	mainasset "asset-management.com/internal/asset-mgmt/main_asset"
	"asset-management.com/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MainAssetUCParam struct {
	MainAssetRepository mainasset.Repository
	// Cloudinary          *cloudinary.Cloudinary
}

func NewMainAssetUsecase(param *MainAssetUCParam) mainasset.Usecase {
	return &mainAssetUC{
		pbbRepo: param.MainAssetRepository,
	}
}

type mainAssetUC struct {
	pbbRepo mainasset.Repository
}

func (uc *mainAssetUC) Index(c *fiber.Ctx) ([]*model.PBB, error) {
	documents, err := uc.pbbRepo.GetIndex()
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (uc *mainAssetUC) CreateMainAsset(c *fiber.Ctx, pbb *model.PBBRequest) (*model.PBB, error) {
	newPbb := &model.PBB{
		ID:           uuid.New(),
		Nop:          pbb.Nop,
		LandArea:     pbb.LandArea,
		BuildingArea: pbb.BuildingArea,
		Road:         pbb.Road,
	}

	if pbb.Alley.Valid {
		newPbb.Alley = pbb.Alley.String
	}

	if pbb.Number.Valid {
		newPbb.Number = pbb.Number.String
	}

	if pbb.AdditionalNumber.Valid {
		newPbb.AdditionalNumber = pbb.AdditionalNumber.String
	}

	if pbb.Description.Valid {
		newPbb.Description = pbb.Description.String
	}

	err := uc.pbbRepo.CreateMainAsset(newPbb)
	if err != nil {
		return nil, err
	}
	return newPbb, nil
}

func (uc *mainAssetUC) GetMainAsset(c *fiber.Ctx, id uuid.UUID) (*model.PBB, error) {
	pbb, err := uc.pbbRepo.GetMainAsset(id)
	if err != nil {
		return nil, err
	}

	return pbb, nil
}

func (uc *mainAssetUC) UpdateMainAsset(c *fiber.Ctx, newData *model.PBB, id uuid.UUID) (*model.PBB, error) {
	pbb, err := uc.pbbRepo.GetMainAsset(id)
	if err != nil {
		return nil, err
	}

	if pbb.ID == uuid.Nil {
		return nil, errors.New("pbb not found")
	}

	// check if newData is different from current data
	if pbb.CheckIfSame(newData) {
		return pbb, nil
	}

	// update document
	pbb.Nop = newData.Nop
	pbb.LandArea = newData.LandArea
	pbb.BuildingArea = newData.BuildingArea
	pbb.Road = newData.Road
	pbb.Alley = newData.Alley
	pbb.AdditionalNumber = newData.AdditionalNumber
	pbb.Number = newData.Number
	pbb.Description = newData.Description

	// var updateDocumentData model.UpdateDocument
	err = uc.pbbRepo.UpdateMainAsset(id, pbb)
	if err != nil {
		return nil, err
	}

	return pbb, nil
}
