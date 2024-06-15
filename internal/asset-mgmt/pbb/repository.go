package pbb

import (
	"asset-management.com/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetIndex() ([]*model.PBB, error)
	GetPBB(uuid.UUID) (*model.PBB, error)
	CreatePBB(*model.PBB) error
	UpdatePBB(id uuid.UUID, document *model.PBB) error
}
