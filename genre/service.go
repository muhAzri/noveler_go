package genre

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateGenre(input CreateGenreInput) (Genre, error)
	UpdateGenre(inputID FindByIDInput, input CreateGenreInput) (Genre, error)
	DeleteGenre(inputID FindByIDInput) error
	GetAllGenres() ([]Genre, error)
	GetGenreByID(inputID FindByIDInput) (Genre, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateGenre(input CreateGenreInput) (Genre, error) {
	genre := Genre{
		ID:        uuid.New(),
		Name:      input.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newGenre, err := s.repository.Create(genre)

	if err != nil {
		return newGenre, err
	}

	return newGenre, nil
}

func (s *service) UpdateGenre(inputID FindByIDInput, input CreateGenreInput) (Genre, error) {
	genre, err := s.repository.GetByID(inputID.ID)
	if err != nil {
		return genre, err
	}

	if genre.ID == uuid.Nil {
		return genre, errors.New("Genre not found with that ID")
	}

	genre.Name = input.Name
	genre.UpdatedAt = time.Now()

	updatedGenre, err := s.repository.Update(genre)

	if err != nil {
		return updatedGenre, err
	}

	return updatedGenre, nil
}

func (s *service) DeleteGenre(inputID FindByIDInput) error {
	err := s.repository.DeleteByID(inputID.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAllGenres() ([]Genre, error) {
	genres, err := s.repository.GetAll()

	if err != nil {
		return genres, err
	}

	return genres, nil
}

func (s *service) GetGenreByID(inputID FindByIDInput) (Genre, error) {
	genre, err := s.repository.GetByID(inputID.ID)

	if err != nil {
		return genre, err
	}

	return genre, nil
}
