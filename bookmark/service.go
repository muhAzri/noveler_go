package bookmark

import "github.com/google/uuid"

type Service interface {
	CreateBookmark(input CreateBookmarkInput, userID string) (Bookmark, error)
	DeleteBookmarkInput(input DeleteBookmarkInput) error
	FindBookmarksByUserID(userID string) ([]Bookmark, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateBookmark(input CreateBookmarkInput, userID string) (Bookmark, error) {
	newBookmark := Bookmark{
		ID:      uuid.New(),
		NovelID: uuid.MustParse(input.NovelID),
		UserID:  uuid.MustParse(userID),
	}

	createdBookmark, err := s.repository.Create(newBookmark)

	if err != nil {
		return createdBookmark, err
	}

	return createdBookmark, err
}

func (s *service) DeleteBookmark(input DeleteBookmarkInput) error {
	err := s.repository.Delete(input.BookmarkID)

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
