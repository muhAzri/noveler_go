package helper

import (
	"log"
	"noveler_go/genre"
	"noveler_go/novel"

	"gorm.io/gorm"
)

func AutomigrateDatabase(db *gorm.DB) {
	var err error

	err = genre.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Error migrating user: %v", err)
	}

	err = novel.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Error migrating novel: %v", err)
	}

}
