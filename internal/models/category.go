package models

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"not null" json:"name"`
}

type categoryType string

type ProductCategory struct {
	Books       categoryType
	Electronics categoryType
	Clothing    categoryType
	Toys        categoryType
	Furniture   categoryType
}

var ProductCategories = ProductCategory{
	Books:       "Books",
	Electronics: "Electronics",
	Clothing:    "Clothing",
	Toys:        "Toys",
	Furniture:   "Furniture",
}

var CategoryMapping = map[string]categoryType{
	"Books":       ProductCategories.Books,
	"Electronics": ProductCategories.Electronics,
	"Clothing":    ProductCategories.Clothing,
	"Toys":        ProductCategories.Toys,
	"Furniture":   ProductCategories.Furniture,
}
