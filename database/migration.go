package database

import (
	"asset-management.com/internal/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Document{})
	if err != nil {
		return err
	}
	return nil
}
