package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name         string    `json:"name" gorm:"type:varchar(255)"`
	ActivaCode   string    `json:"activa_code" gorm:"type:varchar(50)"`
	DueDate      string    `json:"due_date" gorm:"type:date"`
	DocumentURL  string    `json:"document_url" gorm:"type:varchar(255)"`
	DocumentType string    `json:"document_type" gorm:"type:varchar(50)"`
}
