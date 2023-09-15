package helper

import (
	"log"
	"noveler_go/genre"

	"gorm.io/gorm"
)

func AutomigrateDatabase(db *gorm.DB) {
	// var err error

	err := genre.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Error migrating user: %v", err)
	}

}
