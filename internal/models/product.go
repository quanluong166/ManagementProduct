package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Reference  string     `gorm:"not null"`
	Name       string     `gorm:"not null"`
	AddedDate  time.Time  `gorm:"type:date"`
	Status     string     `gorm:"type:varchar(20);check:status IN ('Available', 'Out of Stock', 'On Order')"`
	CategoryID *uuid.UUID `gorm:"type:uuid"`
	Price      float64    `gorm:"default:0"`
	StockCity  string     `gorm:"type:varchar(50)"`
	SupplierID *uuid.UUID `gorm:"type:uuid"`
	Quantity   int        `gorm:"default:0"`
}
