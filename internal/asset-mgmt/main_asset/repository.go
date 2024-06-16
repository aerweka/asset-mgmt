package mainasset

import (
	"asset-management.com/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetIndex() ([]*model.PBB, error)
	GetMainAsset(uuid.UUID) (*model.PBB, error)
	CreateMainAsset(*model.PBB) error
	UpdateMainAsset(id uuid.UUID, document *model.PBB) error
}
