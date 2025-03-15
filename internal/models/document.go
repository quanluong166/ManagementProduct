package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Filename   string    `json:"filename"`
	Path       string    `json:"path"`
	ProductID  uuid.UUID `gorm:"type:uuid" json:"product_id"`
	UploadedAt time.Time `json:"uploaded_at"`
}
