package novel

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateNovel(input CreateNovelInput) (Novel, error)
	UpdateNovel(inputID FindByIDInput, input CreateNovelInput) (Novel, error)
	GetAllNovel() ([]Novel, error)
	GetNovelByID(inputID FindByIDInput) (Novel, error)
	GetNewestNovel() ([]Novel, error)
	GetNewlyUpdatedNovel() ([]Novel, error)
	GetSortByRateNovel() ([]Novel, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateNovel(input CreateNovelInput) (Novel, error) {
	novel := Novel{
		ID:          uuid.New(),
		Title:       input.Title,
		Description: input.Description,
		CoverImage:  input.CoverImage,
		Status:      input.Status,
		Author:      input.Author,
		GenreIDs:    input.GenreIDs,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	newNovel, err := s.repository.Create(novel)

	if err != nil {
		return newNovel, err
	}
	return newNovel, nil
}

func (s *service) UpdateNovel(inputID FindByIDInput, input CreateNovelInput) (Novel, error) {
	novel, err := s.repository.GetByID(inputID.ID)

	if err != nil {
		return novel, err
	}

	if novel.ID == uuid.Nil {
		return novel, errors.New("Novel not found")
	}

	novel.Title = input.Title
	novel.Description = input.Description
	novel.CoverImage = input.CoverImage
	novel.Status = input.Status
	novel.Author = input.Author
	novel.GenreIDs = input.GenreIDs
	novel.UpdatedAt = time.Now().In(time.UTC)

	updatedNovel, err := s.repository.Save(novel)

	if err != nil {
		return updatedNovel, err
	}

	return updatedNovel, nil
}

func (s *service) GetAllNovel() ([]Novel, error) {
	novels, err := s.repository.GetAll()

	if err != nil {
		return novels, err
	}

	return novels, nil
}

func (s *service) GetNovelByID(inputID FindByIDInput) (Novel, error) {
	novel, err := s.repository.GetByID(inputID.ID)

	if err != nil {
		return novel, err
	}

	if novel.ID == uuid.Nil {
		return novel, errors.New("Novel not found")
	}

	return novel, nil
}

func (s *service) GetNewestNovel() ([]Novel, error) {
	novels, err := s.repository.GetNewest()

	if err != nil {
		return novels, err
	}

	return novels, nil
}

func (s *service) GetNewlyUpdatedNovel() ([]Novel, error) {
	novels, err := s.repository.GetNewlyUpdated()

	if err != nil {
		return novels, err
	}

	return novels, nil
}

func (s *service) GetSortByRateNovel() ([]Novel, error) {
	novels, err := s.repository.GetSortByRate()

	if err != nil {
		return novels, err
	}

	return novels, nil
}
