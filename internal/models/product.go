package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Reference  string     `gorm:"not null" json:"reference"`
	Name       string     `gorm:"not null" json:"name"`
	AddedDate  time.Time  `gorm:"default:CURRENT_DATE" json:"added_date"`
	Status     string     `gorm:"type:varchar(20);check:status IN ('Available', 'Out of Stock', 'On Order')" json:"status"`
	CategoryID *uuid.UUID `gorm:"type:uuid" json:"category_id"`
	Price      float64    `gorm:"default:0" json:"price"`
	StockCity  string     `json:"stock_city"`
	SupplierID *uuid.UUID `gorm:"type:uuid" json:"supplier_id"`
	Quantity   int        `gorm:"default:0" json:"quantity"`
}
