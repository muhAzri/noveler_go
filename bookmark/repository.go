package bookmark

import "gorm.io/gorm"

type Repository interface {
	Create(bookmark Bookmark) (Bookmark, error)
	Delete(bookmarkID string) error
	FindByUserID(userID string) ([]Bookmark, error)
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

func (r *repository) Delete(bookmarkID string) error {
	err := r.db.Where("id ? = ", bookmarkID).Delete(&bookmarkID).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByUserID(userID string) ([]Bookmark, error) {
	var bookmarks []Bookmark

	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&bookmarks).Error

	if err != nil {
		return bookmarks, err
	}

	return bookmarks, nil
}
