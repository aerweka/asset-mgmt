package repository

import (
	"asset-management.com/internal/asset-mgmt/document"
	"asset-management.com/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewDocumentRepository(db *gorm.DB) document.Repository {
	return &documentRepo{
		db: db,
	}
}

type documentRepo struct {
	db *gorm.DB
}

func (d *documentRepo) GetIndex(activaCode string) ([]*model.Document, error) {
	var documents []*model.Document
	var condition map[string]interface{}

	if activaCode != "" {
		condition = map[string]interface{}{
			"activa_code": activaCode,
		}
	}

	res := d.db.Where(condition).Find(&documents)
	if res.Error != nil {
		return nil, res.Error
	}

	return documents, nil
}

func (d *documentRepo) GetDocument(id uuid.UUID) (*model.Document, error) {
	var document model.Document

	res := d.db.First(&document, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &document, nil
}

func (d *documentRepo) CreateDocument(documents []*model.Document) error {
	res := d.db.Create(documents)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (d *documentRepo) UpdateDocument(id uuid.UUID, document *model.Document) error {
	res := d.db.Save(document)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
