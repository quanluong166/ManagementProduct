package utils

import (
	"fmt"
	"os"
	"productManagement/internal/models"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func GenerateProductPDF(products []*models.Product) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "B", 16)
	pdf.AddPage()
	pdf.Cell(40, 10, "Product List")

	// Table headers
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(60, 10, "Reference", "1", 0, "C", false, 0, "")
	pdf.CellFormat(60, 10, "Name", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 10, "Price ($)", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Status", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 10, "Stock City", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Table content
	pdf.SetFont("Arial", "", 12)
	for _, product := range products {
		pdf.CellFormat(60, 10, product.Reference, "1", 0, "C", false, 0, "")
		pdf.CellFormat(60, 10, product.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(30, 10, fmt.Sprintf("%.2f", product.Price), "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 10, product.Status, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, product.StockCity, "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	// Ensure output directory exists
	os.MkdirAll("output", os.ModePerm)
	filePath := "output/products.pdf"

	// Save PDF file
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func ConvertToDateOnly(inputTime time.Time) string {
	return inputTime.Format("2006-01-02")
}
