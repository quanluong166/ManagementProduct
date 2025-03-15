package dbsvc

import (
	"fmt"
	"log"
	"os"
	"productManagement/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	GormDB *gorm.DB
}

var (
	DB Database
)

func loadEnvVariables() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func StartDB() error {
	// Load environment variables
	loadEnvVariables()

	// Get the values from the .env file
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	// Connect to the database using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	migrations := []interface{}{
		&models.Category{},
		&models.ProductCategory{},
		&models.Product{},
		&models.Supplier{},
	}

	db.AutoMigrate(migrations...)
	DB = Database{
		GormDB: db,
	}

	fmt.Println("DATABASE CONNTECTED")
	return nil
}

func GetDBConn() Database {
	return DB
}
