package novel

import "time"

type NovelFormatter struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	CoverImage string    `json:"cover_image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
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
