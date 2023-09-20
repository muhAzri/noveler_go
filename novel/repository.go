package novel

import "gorm.io/gorm"

type Repository interface {
	Create(novel Novel) (Novel, error)
	Save(novel Novel) (Novel, error)
	GetByID(ID string) (Novel, error)
	GetAll() ([]Novel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(novel Novel) (Novel, error) {
	err := r.db.Create(&novel).Error

	if err != nil {
		return novel, err
	}

	return novel, nil
}

func (r *repository) Save(novel Novel) (Novel, error) {
	err := r.db.Save(&novel).Error

	if err != nil {
		return novel, err
	}

	return novel, nil
}

func (r *repository) GetByID(id string) (Novel, error) {
	var novel Novel

	err := r.db.Where("id = ?", id).Find(&novel).Preload("Genres").Error

	if err != nil {
		return novel, err
	}

	return novel, nil
}

func (r *repository) GetAll() ([]Novel, error) {
	var novels []Novel

	err := r.db.Order("updated_at DESC").Find(&novels).Error

	if err != nil {
		return novels, err
	}

	return novels, nil
}
