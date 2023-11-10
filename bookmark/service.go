package bookmark

import (
	"github.com/google/uuid"
)

type Service interface {
	CreateBookmark(input CreateBookmarkInput, userID string) error
	DeleteBookmark(input CreateBookmarkInput, userID string) error
	FindBookmarksByUserID(userID string) ([]Bookmark, error)
	FindByUserAndNovelID(userID string, novelID string) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateBookmark(input CreateBookmarkInput, userID string) error {
	newBookmark := Bookmark{
		ID:      uuid.New(),
		NovelID: uuid.MustParse(input.NovelID),
		UserID:  uuid.MustParse(userID),
	}

	_, err := s.repository.Create(newBookmark)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteBookmark(input CreateBookmarkInput, userID string) error {
	err := s.repository.Delete(userID, input.NovelID)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) FindBookmarksByUserID(userID string) ([]Bookmark, error) {
	bookmarks, err := s.repository.FindByUserID(userID)

	if err != nil {
		return bookmarks, err
	}

	return bookmarks, nil
}

func (s *service) FindByUserAndNovelID(userID string, novelID string) (bool, error) {
	bookmarked, err := s.repository.FindByUserAndNovelID(userID, novelID)

	if err != nil {
		return false, err
	}

	if bookmarked.ID == uuid.Nil {
		return false, nil
	}

	return true, nil
}
