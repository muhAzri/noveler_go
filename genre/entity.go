package genre

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Genre struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	Name      string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Genre{})

	if err != nil {
		return err
	}

	return nil
}
