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

func (d *documentRepo) GetIndex() ([]model.Document, error) {
	var documents []model.Document

	res := d.db.Find(&documents)
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
