package novel

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Novel struct {
	ID          uuid.UUID      `gorm:"type:char(36);primary_key"`
	Title       string         `gorm:"type:varchar(255)"`
	Description string         `gorm:"type:text"`
	CoverImage  string         `gorm:"type:text"`
	Status      string         `gorm:"type:varchar(255)"`
	Author      string         `gorm:"type:varchar(255)"`
	Rating      int            `gorm:"type:bigint"`
	GenreIDs    pq.StringArray `gorm:"type:text[]"`
	// Chapters    []chapter.Chapter
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Novel{})

	if err != nil {
		return err
	}

	return nil
}
