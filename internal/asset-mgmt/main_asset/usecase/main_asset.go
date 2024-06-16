package usecase

import (
	"errors"

	"asset-management.com/internal/asset-mgmt/document"
	mainasset "asset-management.com/internal/asset-mgmt/main_asset"
	"asset-management.com/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MainAssetUCParam struct {
	MainAssetRepository mainasset.Repository
	DocumentRepo document.Repository
}

func NewMainAssetUsecase(param *MainAssetUCParam) mainasset.Usecase {
	return &mainAssetUC{
		mainAssetrepo: param.MainAssetRepository,
		documentRepo: param.DocumentRepo,
	}
}

type mainAssetUC struct {
	mainAssetrepo mainasset.Repository
	documentRepo document.Repository
}

func (uc *mainAssetUC) Index(c *fiber.Ctx) ([]*model.MainAsset, error) {
	mainAssets, err := uc.mainAssetrepo.GetIndex()
	if err != nil {
		return nil, err
	}

	return mainAssets, nil
}

func (uc *mainAssetUC) CreateMainAsset(c *fiber.Ctx, mainAsset *model.MainAssetRequest) (*model.MainAsset, error) {
	_, err := uc.documentRepo.GetDocument(mainAsset.DocumentID)
	if err != nil {
		return nil, err
	}

	newMainAsset := &model.MainAsset{
		ID:           uuid.New(),
		DocumentID:   mainAsset.DocumentID,
		Category:     mainAsset.Category,
		AssetTotal:   mainAsset.AssetTotal,
		District:     mainAsset.District,
		SubDistrict:  mainAsset.SubDistrict,
		UrbanVillage: mainAsset.UrbanVillage,
	}

	err = uc.mainAssetrepo.CreateMainAsset(newMainAsset)
	if err != nil {
		return nil, err
	}
	return newMainAsset, nil
}

func (uc *mainAssetUC) GetMainAsset(c *fiber.Ctx, id uuid.UUID) (*model.MainAsset, error) {
	mainAsset, err := uc.mainAssetrepo.GetMainAsset(id)
	if err != nil {
		return nil, err
	}

	return mainAsset, nil
}

func (uc *mainAssetUC) UpdateMainAsset(c *fiber.Ctx, newData *model.MainAsset, id uuid.UUID) (*model.MainAsset, error) {
	mainAsset, err := uc.mainAssetrepo.GetMainAsset(id)
	if err != nil {
		return nil, err
	}

	if mainAsset.ID == uuid.Nil {
		return nil, errors.New("pbb not found")
	}

	// check if newData is different from current data
	if mainAsset.CheckIfSame(newData) {
		return mainAsset, nil
	}

	// update document
	mainAsset.DocumentID = newData.DocumentID
	mainAsset.AssetTotal = newData.AssetTotal
	mainAsset.Category = newData.Category
	mainAsset.District = newData.District
	mainAsset.SubDistrict = newData.SubDistrict
	mainAsset.UrbanVillage = newData.UrbanVillage

	// var updateDocumentData model.UpdateDocument
	err = uc.mainAssetrepo.UpdateMainAsset(id, mainAsset)
	if err != nil {
		return nil, err
	}

	return mainAsset, nil
}
