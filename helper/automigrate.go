package helper

import (
	"log"
	"noveler_go/bookmark"
	"noveler_go/chapter"
	"noveler_go/genre"
	"noveler_go/novel"
	"noveler_go/user"

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
	err = chapter.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Error migrating novel: %v", err)
	}

	err = user.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Error migrating novel: %v", err)
	}

	err = bookmark.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Error migrating novel: %v", err)
	}

}
