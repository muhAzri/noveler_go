package novel

import (
	"time"
)

type NovelFormatter struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	CoverImage string    `json:"cover_image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type SearchNovelFormatter struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	CoverImage string    `json:"cover_image"`
	Rating     float32   `json:"rating"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type DetailNovelFormatter struct {
	ID            string      `json:"id"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	ChaptersCount int         `json:"chapters_count"`
	CoverImage    string      `json:"cover_image"`
	Status        string      `json:"status"`
	Author        string      `json:"author"`
	Rating        float32     `json:"rating"`
	Genres        interface{} `json:"genres"`
	Bookmarked    bool        `json:"bookmarked"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

func FormatNovel(novel Novel) NovelFormatter {
	formatter := NovelFormatter{
		ID:         novel.ID.String(),
		Title:      novel.Title,
		CoverImage: novel.CoverImage,
		CreatedAt:  novel.CreatedAt,
		UpdatedAt:  novel.UpdatedAt,
	}

	return formatter
}

func FormatNovels(novels []Novel) []NovelFormatter {
	novelsFormatter := []NovelFormatter{}

	for _, novel := range novels {
		novelFormatter := FormatNovel(novel)
		novelsFormatter = append(novelsFormatter, novelFormatter)
	}

	return novelsFormatter
}

func FormatSearchNovel(novel Novel) SearchNovelFormatter {
	formatter := SearchNovelFormatter{
		ID:         novel.ID.String(),
		Title:      novel.Title,
		CoverImage: novel.CoverImage,
		Rating:     novel.Rating,
		CreatedAt:  novel.CreatedAt,
		UpdatedAt:  novel.UpdatedAt,
	}

	return formatter
}

func FormatSearchNovels(novels []Novel) []SearchNovelFormatter {
	novelsFormatter := []SearchNovelFormatter{}

	for _, novel := range novels {
		novelFormatter := FormatSearchNovel(novel)
		novelsFormatter = append(novelsFormatter, novelFormatter)
	}

	return novelsFormatter
}

func FormatDetailNovel(novel Novel, genres interface{}, ChaptersCount int, Bookmarked bool) DetailNovelFormatter {
	formatter := DetailNovelFormatter{
		ID:            novel.ID.String(),
		Title:         novel.Title,
		Description:   novel.Description,
		Bookmarked:    Bookmarked,
		ChaptersCount: ChaptersCount,
		CoverImage:    novel.CoverImage,
		Status:        novel.Status,
		Author:        novel.Author,
		Rating:        novel.Rating,
		Genres:        genres,
		CreatedAt:     novel.CreatedAt,
		UpdatedAt:     novel.UpdatedAt,
	}

	return formatter
}
