package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "vault.db"
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Print("Database connected succesfully")
	return db
}
