package models

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"not null" json:"name"`
}
