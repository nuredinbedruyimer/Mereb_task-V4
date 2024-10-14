package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

// ConnectDB establishes a connection to the PostgreSQL database.
func ConnectDB() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Database connection established successfully!")

	// Set the global DB instance
	DB = DBInstance{
		Db: db,
	}
}

// GetDB returns the database instance.
func GetDB() *gorm.DB {
	return DB.Db
}
