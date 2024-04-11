package document

import (
	"asset-management.com/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetIndex() ([]model.Document, error)
	GetDocument(uuid.UUID) (*model.Document, error)
}
