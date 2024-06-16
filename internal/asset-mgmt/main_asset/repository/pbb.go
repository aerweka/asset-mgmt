package repository

import (
	mainasset "asset-management.com/internal/asset-mgmt/main_asset"
	"asset-management.com/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewMainAssetRepository(db *gorm.DB) mainasset.Repository {
	return &mainAssetRepo{
		db: db,
	}
}

type mainAssetRepo struct {
	db *gorm.DB
}

func (d *mainAssetRepo) GetIndex() ([]*model.PBB, error) {
	var pbb []*model.PBB
	var condition map[string]interface{}

	res := d.db.Where(condition).Find(&pbb)
	if res.Error != nil {
		return nil, res.Error
	}

	return pbb, nil
}

func (d *mainAssetRepo) GetMainAsset(id uuid.UUID) (*model.PBB, error) {
	var pbb model.PBB

	res := d.db.First(&pbb, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &pbb, nil
}

func (d *mainAssetRepo) CreateMainAsset(pbb *model.PBB) error {
	res := d.db.Create(pbb)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (d *mainAssetRepo) UpdateMainAsset(id uuid.UUID, pbb *model.PBB) error {
	res := d.db.Save(pbb)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
