package usecase

import (
	"errors"

	"asset-management.com/internal/asset-mgmt/pbb"
	"asset-management.com/internal/model"
	"asset-management.com/pkg/cloudinary"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PBBUCParam struct {
	PBBRepository pbb.Repository
	Cloudinary    *cloudinary.Cloudinary
}

func NewPbbUsecase(param *PBBUCParam) pbb.Usecase {
	return &pbbUC{
		pbbRepo:    param.PBBRepository,
		cloudinary: param.Cloudinary,
	}
}

type pbbUC struct {
	pbbRepo    pbb.Repository
	cloudinary *cloudinary.Cloudinary
}

func (uc *pbbUC) Index(c *fiber.Ctx) ([]*model.PBB, error) {
	documents, err := uc.pbbRepo.GetIndex()
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (uc *pbbUC) CreatePBB(c *fiber.Ctx, pbb *model.PBBRequest) (*model.PBB, error) {
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

	err := uc.pbbRepo.CreatePBB(newPbb)
	if err != nil {
		return nil, err
	}
	return newPbb, nil
}

func (uc *pbbUC) GetPBB(c *fiber.Ctx, id uuid.UUID) (*model.PBB, error) {
	pbb, err := uc.pbbRepo.GetPBB(id)
	if err != nil {
		return nil, err
	}

	return pbb, nil
}

func (uc *pbbUC) UpdatePBB(c *fiber.Ctx, newData *model.PBB, id uuid.UUID) (*model.PBB, error) {
	pbb, err := uc.pbbRepo.GetPBB(id)
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
	err = uc.pbbRepo.UpdatePBB(id, pbb)
	if err != nil {
		return nil, err
	}

	return pbb, nil
}
