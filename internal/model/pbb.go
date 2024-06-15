package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PBB struct {
	ID               uuid.UUID      `json:"id" gorm:"type:uuid;primarykey;default:gen_random_uuid()"`
	Nop              int64          `json:"nop" gorm:"type:int8"`
	LandArea         float64        `json:"land_area" gorm:"type:numeric(10,2)"`
	BuildingArea     float64        `json:"building_area" gorm:"type:numeric(10,2)"`
	Road             string         `json:"road" gorm:"varchar(255)"`
	Alley            string         `json:"alley" gorm:"varchar(255)"`
	Number           string         `json:"number" gorm:"varchar(100)"`
	AdditionalNumber string         `json:"additional_number" gorm:"varchar(100)"`
	Description      string         `json:"description" gorm:"varchar(255)"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}

func (d PBB) CheckIfSame(newData *PBB) bool {
	if d.Nop == newData.Nop &&
		d.LandArea == newData.LandArea &&
		d.BuildingArea == newData.BuildingArea &&
		d.Road == newData.Road &&
		d.Alley == newData.Alley &&
		d.Number == newData.Number &&
		d.AdditionalNumber == newData.AdditionalNumber &&
		d.Description == newData.Description {
		return true
	}
	return false
}

type PBBRequest struct {
	Nop              int64          `json:"nop"`
	LandArea         float64        `json:"land_area"`
	BuildingArea     float64        `json:"building_area"`
	Road             string         `json:"road"`
	Alley            sql.NullString `json:"alley"`
	Number           sql.NullString `json:"number"`
	AdditionalNumber sql.NullString `json:"additional_number"`
	Description      sql.NullString `json:"description"`
}
