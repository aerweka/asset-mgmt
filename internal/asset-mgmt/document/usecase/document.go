package usecase

import (
	"database/sql"
	"errors"
	"fmt"

	"asset-management.com/internal/asset-mgmt/document"
	"asset-management.com/internal/model"
	"asset-management.com/pkg/cloudinary"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DocumentUCParam struct {
	DocumentRepo document.Repository
	Cloudinary   *cloudinary.Cloudinary
}

func NewDocumentUsecase(param *DocumentUCParam) document.Usecase {
	return &documentUC{
		documentRepo: param.DocumentRepo,
		cloudinary:   param.Cloudinary,
	}
}

type documentUC struct {
	documentRepo document.Repository
	cloudinary   *cloudinary.Cloudinary
}

func (uc *documentUC) Index(c *fiber.Ctx, queries map[string]string) ([]*model.Document, error) {
	documents, err := uc.documentRepo.GetIndex(queries)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (uc *documentUC) CreateDocument(c *fiber.Ctx, document *model.DocumentRequest) ([]*model.Document, error) {
	var NewDocument []*model.Document
	for _, v := range document.Documents {
		newDocumentUUID := uuid.New()
		documentUrl, err := uc.cloudinary.SendPdf(v.DocumentURL, fmt.Sprintf("%s-document", newDocumentUUID))
		if err != nil {
			return nil, err
		}

		document := &model.Document{
			ID:   newDocumentUUID,
			Name: v.Name,
			ActivaCode: sql.NullString{
				String: document.ActiveCode,
				Valid:  true,
			},
			DueDate:          v.DueDate,
			DocumentURL:      documentUrl.SecureURL,
			DocumentPublicId: documentUrl.PublicID,
			DocumentType:     v.DocumentType,
			AcquisitionCost:  v.AcquisitionCost,
			AcquisitionDate:  v.AcquisitionDate,
			Area:             v.Area,
		}

		NewDocument = append(NewDocument, document)
	}

	err := uc.documentRepo.CreateDocument(NewDocument)
	if err != nil {
		return nil, err
	}
	return NewDocument, nil
}

func (uc *documentUC) GetDocument(c *fiber.Ctx, id uuid.UUID) (*model.Document, error) {
	document, err := uc.documentRepo.GetDocument(id)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (uc *documentUC) UpdateDocument(c *fiber.Ctx, newData *model.Document, id uuid.UUID) (*model.Document, error) {
	document, err := uc.documentRepo.GetDocument(id)
	if err != nil {
		return nil, err
	}

	if document.ID == uuid.Nil {
		return nil, errors.New("document not found")
	}

	// check if newData is different from current data
	if document.CheckIfSame(newData) {
		return document, nil
	}

	// update document
	document.Name = newData.Name
	document.ActivaCode = sql.NullString{
		String: newData.ActivaCode.String,
		Valid:  true,
	}
	document.DueDate = newData.DueDate
	document.AcquisitionCost = newData.AcquisitionCost
	document.AcquisitionDate = newData.AcquisitionDate
	document.Area = newData.Area
	document.DocumentType = newData.DocumentType
	if document.DocumentURL != newData.DocumentURL {
		documentUrl, err := uc.cloudinary.SendPdf(newData.DocumentURL, fmt.Sprintf("%s-document", document.ID))
		if err != nil {
			return nil, err
		}
		document.DocumentURL = documentUrl.SecureURL
		document.DocumentPublicId = documentUrl.PublicID
	}

	err = uc.documentRepo.UpdateDocument(id, document)
	if err != nil {
		return nil, err
	}

	return document, nil
}
