package mainasset

import (
	"asset-management.com/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetIndex() ([]*model.MainAsset, error)
	GetMainAsset(uuid.UUID) (*model.MainAsset, error)
	CreateMainAsset(*model.MainAsset) error
	UpdateMainAsset(id uuid.UUID, document *model.MainAsset) error
}
