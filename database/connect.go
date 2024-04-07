package database

import (
	"fmt"

	"asset-management.com/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	var err error

	p := cfg.Database.Port

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, p, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return DB, nil
}
