package models

type Status string

type ProductStatus struct {
	Available  Status
	OnOrder    Status
	OutOfStock Status
}

var ProductStatuses = ProductStatus{
	Available:  "Available",
	OnOrder:    "On Order",
	OutOfStock: "Out Of Stock",
}

var StatusMapping = map[string]Status{
	"AVAILABLE":    ProductStatuses.Available,
	"ON_ORDER":     ProductStatuses.OnOrder,
	"OUT_OF_STOCK": ProductStatuses.OutOfStock,
}
