package chapter

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chapter struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	NovelID   uuid.UUID `gorm:"type:char(36);index"`
	Title     string    `gorm:"type:varchar(255)"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Chapter{})

	if err != nil {
		return err
	}

	return nil
}
