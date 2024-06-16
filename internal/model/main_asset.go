package model

import "github.com/google/uuid"

type MainAsset struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primarykey;default:gen_random_uuid()"`
	DocumentID   uuid.UUID `json:"document_id"`
	Document     Document
	Category     string `json:"category" gorm:"varchar(128)"`
	AssetTotal   int8   `json:"asset_total" gorm:"int8"`
	District     string `json:"district" gorm:"varchar(255)"`
	SubDistrict  string `json:"sub_district" gorm:"varchar(255)"`
	UrbanVillage string `json:"urban_village" gorm:"varchar(255)"`
}