package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DocumentRequest struct {
	ActiveCode string     `json:"activa_code"`
	Documents  []Document `json:"documents"`
}

type Document struct {
	ID               uuid.UUID      `json:"id" gorm:"type:uuid;primarykey;default:gen_random_uuid()"`
	Name             string         `json:"name" gorm:"type:varchar(255)"`
	ActivaCode       sql.NullString `json:"activa_code" gorm:"type:varchar(50)"`
	DueDate          time.Time      `json:"due_date" gorm:"type:timestamptz(0)"`
	DocumentURL      string         `json:"document_url" gorm:"type:varchar(255)"`
	DocumentType     string         `json:"document_type" gorm:"type:varchar(50)"`
	DocumentPublicId string         `json:"document_publid_id" gorm:"type:varchar(128)"`
	AcquisitionCost  float64        `json:"acquisition_cost" gorm:"type:numeric(10,2)"`
	AcquisitionDate  time.Time      `json:"acquisition_date" gorm:"type:timestamptz(0)"`
	Area             float64        `json:"area" gorm:"type:numeric(10,2)"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}

func (d Document) CheckIfSame(newData *Document) bool {
	if newData.Name == d.Name &&
		newData.ActivaCode.String == d.ActivaCode.String &&
		newData.DueDate == d.DueDate &&
		newData.DocumentURL == d.DocumentURL &&
		newData.DocumentPublicId == d.DocumentPublicId &&
		newData.DocumentType == d.DocumentType &&
		newData.AcquisitionCost == d.AcquisitionCost &&
		newData.AcquisitionDate == d.AcquisitionDate &&
		newData.Area == d.Area {
		return true
	}
	return false
}

type UpdateDocument struct {
	Name            string `json:"name"`
	DueDate         string `json:"due_date"`
	ActivaCode      string `json:"activa_code"`
	DocumentType    string `json:"document_type"`
	DocumentURL     string `json:"document_url"`
	AcquisitionCost int64  `json:"acquisition_cost"`
	AcquisitionDate string `json:"acquisition_date"`
}
