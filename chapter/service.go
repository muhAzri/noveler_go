package chapter

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateChapter(inputId FindByIDInput, input CreateChapterInput) (Chapter, error)
	GetPaginatedChapters(offset, limit int, NovelID string) ([]Chapter, error)
	GetChaptersByID(inputID string) ([]Chapter, error)
	GetByID(inputID FindByIDInput) (Chapter, error)
	Delete(inputID FindByIDInput) error
	UpdateChapter(inputID FindByIDInput, input CreateChapterInput) (Chapter, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateChapter(inputID FindByIDInput, input CreateChapterInput) (Chapter, error) {

	chapter := Chapter{
		ID:        uuid.New(),
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now().In(time.UTC),
		UpdatedAt: time.Now().In(time.UTC),
	}

	novelID, err := uuid.Parse(inputID.ID)

	if err != nil {
		return chapter, err
	}

	chapter.NovelID = novelID

	newChapter, err := s.repository.Create(chapter)

	if err != nil {
		return newChapter, err
	}

	return newChapter, nil

}

func (s *service) GetPaginatedChapters(offset, limit int, NovelID string) ([]Chapter, error) {
	chapters, err := s.repository.FindByNovelID(offset, limit, NovelID)

	if err != nil {
		return chapters, err
	}

	return chapters, nil
}

func (s *service) GetByID(inputID FindByIDInput) (Chapter, error) {

	chapter, err := s.repository.GetById(inputID.ID)

	if err != nil {
		return chapter, err
	}

	return chapter, nil
}

func (s *service) Delete(inputID FindByIDInput) error {
	err := s.repository.Delete(inputID.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateChapter(inputID FindByIDInput, input CreateChapterInput) (Chapter, error) {
	chapter, err := s.repository.GetById(inputID.ID)
	if err != nil {
		return chapter, err
	}

	if chapter.ID == uuid.Nil {
		return chapter, errors.New("Chapter not found with that ID")
	}

	chapter.Title = input.Title
	chapter.Content = input.Content
	chapter.UpdatedAt = time.Now().UTC()

	updatedChapter, err := s.repository.Update(chapter)

	if err != nil {
		return updatedChapter, err
	}

	return updatedChapter, nil
}

func (s *service) GetChaptersByID(inputID string) ([]Chapter, error) {
	chapters, err := s.repository.FindsByID(inputID)

	if err != nil {
		return chapters, err
	}

	return chapters, nil
}
