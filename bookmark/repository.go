package bookmark

import "gorm.io/gorm"

type Repository interface {
	Create(bookmark Bookmark) (Bookmark, error)
	Delete(userID string, novelID string) error
	FindByUserID(userID string) ([]Bookmark, error)
	FindByUserAndNovelID(userID string, novelID string) (Bookmark, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(bookmark Bookmark) (Bookmark, error) {
	err := r.db.Create(&bookmark).Error

	if err != nil {
		return bookmark, err
	}

	return bookmark, nil
}

func (r *repository) Delete(userID string, novelID string) error {
	err := r.db.Where("user_id = ? AND novel_id = ?", userID, novelID).Delete(&Bookmark{}).Error
	return err
}
func (r *repository) FindByUserID(userID string) ([]Bookmark, error) {
	var bookmarks []Bookmark

	err := r.db.Where("user_id = ?", userID).Preload("Novel").Find(&bookmarks).Error

	if err != nil {
		return bookmarks, err
	}

	return bookmarks, nil
}

func (r *repository) FindByUserAndNovelID(userID string, novelID string) (Bookmark, error) {
	var bookmark Bookmark

	err := r.db.Where("user_id = ? AND novel_id = ?", userID, novelID).Find(&bookmark).Error

	if err != nil {
		return bookmark, err
	}

	return bookmark, nil
}
