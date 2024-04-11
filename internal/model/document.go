package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name            string    `json:"name" gorm:"type:varchar(255)"`
	ActivaCode      string    `json:"activa_code" gorm:"type:varchar(50)"`
	DueDate         string    `json:"due_date" gorm:"type:date"`
	DocumentURL     string    `json:"document_url" gorm:"type:varchar(255)"`
	DocumentType    string    `json:"document_type" gorm:"type:varchar(50)"`
	Amount          int16     `json:"amount" gorm:"type:numeric(10,2)"`
	AcquisitionCost int64     `json:"acquisition_cost" gorm:"type:numeric(10,2)"`
	AcquisitionDate string    `json:"acquisition_date" gorm:"type:date"`
}

type UpdateDocument struct {
	Name            string `json:"name"`
	DueDate         string `json:"due_date"`
	ActivaCode      string `json:"activa_code"`
	DocumentType    string `json:"document_type"`
	DocumentURL     string `json:"document_url" `
	AcquisitionCost int64  `json:"acquisition_cost" `
	AcquisitionDate string `json:"acquisition_date" `
}
