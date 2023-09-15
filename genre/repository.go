package genre

import "gorm.io/gorm"

type Repository interface {
	Create(genre Genre) (Genre, error)
	Update(genre Genre) (Genre, error)
	GetByID(ID string) (Genre, error)
	DeleteByID(ID string) error
	GetAll() ([]Genre, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(genre Genre) (Genre, error) {
	err := r.db.Create(&genre).Error

	if err != nil {
		return genre, err
	}

	return genre, nil
}

func (r *repository) Update(genre Genre) (Genre, error) {
	err := r.db.Save(&genre).Error

	if err != nil {
		return genre, err
	}

	return genre, nil
}

func (r *repository) GetByID(ID string) (Genre, error) {
	var genre Genre

	err := r.db.Where("id = ?", ID).Find(&genre).Error

	if err != nil {
		return genre, err
	}

	return genre, nil
}

func (r *repository) DeleteByID(ID string) error {
	err := r.db.Where("id = ?", ID).Delete(&Genre{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAll() ([]Genre, error) {
	var genres []Genre

	err := r.db.Find(&genres).Error

	if err != nil {
		return genres, err
	}

	return genres, nil
}
