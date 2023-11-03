package chapter

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(chapter Chapter) (Chapter, error)
	FindByNovelID(offset, limit int, NovelID string) ([]Chapter, error)
	FindsByID(novelID string) ([]Chapter, error)
	GetById(ID string) (Chapter, error)
	Delete(ID string) error
	Update(chapter Chapter) (Chapter, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(chapter Chapter) (Chapter, error) {
	err := r.db.Create(&chapter).Error

	if err != nil {
		return chapter, err
	}

	return chapter, nil
}

func (r *repository) FindByNovelID(offset, limit int, NovelID string) ([]Chapter, error) {
	var chapters []Chapter

	query := r.db.Where("novel_id = ?", NovelID).Order("created_at desc").Offset(offset).Limit(limit).Find(&chapters)

	if query.Error != nil {
		return chapters, query.Error
	}

	return chapters, nil
}

func (r *repository) GetById(ID string) (Chapter, error) {
	var chapter Chapter

	query := r.db.Where("id = ?", ID).Find(&chapter)

	if query.Error != nil {
		return chapter, query.Error
	}

	return chapter, nil
}

func (r *repository) Delete(ID string) error {
	err := r.db.Where("id = ?", ID).Delete(&Chapter{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(chapter Chapter) (Chapter, error) {
	err := r.db.Save(&chapter).Error

	if err != nil {
		return chapter, err
	}

	return chapter, nil
}

func (r *repository) FindsByID(novelID string) ([]Chapter, error) {
	var chapters []Chapter

	query := r.db.Where("novel_id = ?", novelID).Order("created_at desc").Find(&chapters)

	if query.Error != nil {
		return chapters, query.Error
	}

	return chapters, nil
}
