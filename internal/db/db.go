package db

import (
	"log"
	messageserver "pod/internal/messageServer"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to DataBase: %v", err)
	}

	if err := db.AutoMigrate(&messageserver.Message{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	return db, nil
}
