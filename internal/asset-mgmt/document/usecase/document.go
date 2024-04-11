package usecase

import (
	"errors"

	"asset-management.com/internal/asset-mgmt/document"
	"asset-management.com/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DocumentUCParam struct {
	DocumentRepo document.Repository
}

func NewDocumentUsecase(param DocumentUCParam) document.Usecase {
	return &documentUC{
		documentRepo: param.DocumentRepo,
	}
}

type documentUC struct {
	documentRepo document.Repository
}

func (uc *documentUC) Index(c *fiber.Ctx) ([]model.Document, error) {
	documents, err := uc.documentRepo.GetIndex()
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (uc *documentUC) CreateDocument(c *fiber.Ctx, document *model.Document) error {
	return nil
}

func (uc *documentUC) GetDocument(c *fiber.Ctx, id uuid.UUID) (*model.Document, error) {
	document, err := uc.documentRepo.GetDocument(id)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (uc *documentUC) UpdateDocument(c *fiber.Ctx, id uuid.UUID) error {

	document, err := uc.documentRepo.GetDocument(id)
	if err != nil {
		return err
	}

	if document.ID == uuid.Nil {
		return errors.New("document not found")
	}

	var updateDocumentData model.UpdateDocument
	err = c.BodyParser(&updateDocumentData)
	if err != nil {
		return err
	}

	return nil
}
