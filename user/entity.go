package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:char(36);primary_key"`
	Username     string    `gorm:"type:varchar(255);unique"`
	Email        string    `gorm:"type:varchar(255);unique"`
	Role         string    `gorm:"type:varchar(255)"`
	PasswordHash string    `gorm:"type:varchar(255)"`
	Salt         string    `gorm:"type:varchar(36)"`
	CreatedAt    time.Time `gorm:"type:timestamp"`
	UpdatedAt    time.Time `gorm:"type:timestamp"`
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})

	if err != nil {
		return err
	}

	return nil
}
