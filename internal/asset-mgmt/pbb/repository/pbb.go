package repository

import (
	"asset-management.com/internal/asset-mgmt/pbb"
	"asset-management.com/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewPbbRepository(db *gorm.DB) pbb.Repository {
	return &pbbRepo{
		db: db,
	}
}

type pbbRepo struct {
	db *gorm.DB
}

func (d *pbbRepo) GetIndex() ([]*model.PBB, error) {
	var pbb []*model.PBB
	var condition map[string]interface{}

	res := d.db.Where(condition).Find(&pbb)
	if res.Error != nil {
		return nil, res.Error
	}

	return pbb, nil
}

func (d *pbbRepo) GetPBB(id uuid.UUID) (*model.PBB, error) {
	var pbb model.PBB

	res := d.db.First(&pbb, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &pbb, nil
}

func (d *pbbRepo) CreatePBB(pbb *model.PBB) error {
	res := d.db.Create(pbb)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (d *pbbRepo) UpdatePBB(id uuid.UUID, pbb *model.PBB) error {
	res := d.db.Save(pbb)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
