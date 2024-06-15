package document

import (
	"asset-management.com/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetIndex(string) ([]*model.Document, error)
	GetDocument(uuid.UUID) (*model.Document, error)
	CreateDocument([]*model.Document) error
	UpdateDocument(id uuid.UUID, document *model.Document) error
}
