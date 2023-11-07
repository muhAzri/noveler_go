package novel

import (
	"math/rand"

	"gorm.io/gorm"
)

type Repository interface {
	Create(novel Novel) (Novel, error)
	Save(novel Novel) (Novel, error)
	GetByID(ID string) (Novel, error)
	GetAll() ([]Novel, error)
	GetSortByRate() ([]Novel, error)
	GetNewest() ([]Novel, error)
	GetNewlyUpdated() ([]Novel, error)
	GetRandomNovel() ([]Novel, error)
	SearchNovels(params NovelSearchParametersInput) ([]Novel, error)
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

	err := r.db.Preload("Chapters").Where("id = ?", id).Find(&novel).Error

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

func (r *repository) GetSortByRate() ([]Novel, error) {
	var novels []Novel

	err := r.db.Order("rating ASC").Limit(10).Find(&novels).Error

	if err != nil {
		return novels, err
	}

	return novels, nil
}

func (r *repository) GetNewest() ([]Novel, error) {
	var novels []Novel

	err := r.db.Order("created_at DESC").Limit(10).Find(&novels).Error

	if err != nil {
		return novels, err
	}

	return novels, nil
}

func (r *repository) GetNewlyUpdated() ([]Novel, error) {
	var novels []Novel

	err := r.db.Order("updated_at DESC").Limit(10).Find(&novels).Error

	if err != nil {
		return novels, err
	}

	return novels, nil
}

func (r *repository) GetRandomNovel() ([]Novel, error) {
	var novels []Novel

	err := r.db.Find(&novels).Error

	if err != nil {
		return novels, err
	}

	randomNovels := make([]Novel, 0, 10)
	for i := 0; i < 10; i++ {
		randomNovels = append(randomNovels, novels[rand.Intn(len(novels))])
	}

	return randomNovels, nil
}

func (r *repository) SearchNovels(params NovelSearchParametersInput) ([]Novel, error) {
	var novels []Novel

	query := r.db.Model(&Novel{})

	// Apply filters based on the provided parameters
	if params.Title != "" {
		query = query.Where("title ILIKE ?", "%"+params.Title+"%")
	}
	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}
	if len(params.Genres) > 0 {
		for _, genreID := range params.Genres {
			query = query.Where("genre_ids @> ARRAY[?]", genreID)
		}
	}

	// Calculate offset for pagination
	offset := (params.Page - 1) * params.PageSize
	query = query.Offset(offset).Limit(params.PageSize)

	if err := query.Find(&novels).Error; err != nil {
		return nil, err
	}

	return novels, nil
}
