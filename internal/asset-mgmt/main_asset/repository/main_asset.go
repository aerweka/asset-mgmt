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

func (d *mainAssetRepo) GetIndex() ([]*model.MainAsset, error) {
	var mainAssets []*model.MainAsset
	var condition map[string]interface{}

	res := d.db.Where(condition).Find(&mainAssets)
	if res.Error != nil {
		return nil, res.Error
	}

	return mainAssets, nil
}

func (d *mainAssetRepo) GetMainAsset(id uuid.UUID) (*model.MainAsset, error) {
	var mainAsset model.MainAsset

	res := d.db.First(&mainAsset, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &mainAsset, nil
}

func (d *mainAssetRepo) CreateMainAsset(mainAsset *model.MainAsset) error {
	res := d.db.Create(mainAsset)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (d *mainAssetRepo) UpdateMainAsset(id uuid.UUID, mainAsset *model.MainAsset) error {
	res := d.db.Save(mainAsset)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
