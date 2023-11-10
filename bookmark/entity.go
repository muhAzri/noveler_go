package bookmark

import (
	"noveler_go/novel"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bookmark struct {
	ID      uuid.UUID `gorm:"type:char(36);primary_key"`
	NovelID uuid.UUID `gorm:"type:char(36);index;ForeignKey:NovelID"`
	Novel   novel.Novel 
	UserID  uuid.UUID `gorm:"type:char(36);index"`
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Bookmark{})

	if err != nil {
		return err
	}

	return nil
}
