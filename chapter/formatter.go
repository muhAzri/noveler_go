package chapter

import (
	"time"

	"github.com/google/uuid"
)

type ChapterFormatter struct {
	ID        string    `json:"id"`
	NovelID   string    `json:"novel_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ChapterDetailFormatter struct {
	ID              string            `json:"id"`
	NovelID         string            `json:"novel_id"`
	Title           string            `json:"title"`
	Content         string            `json:"content"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	PreviousChapter *ChapterFormatter `json:"previous_chapter,omitempty"`
	NextChapter     *ChapterFormatter `json:"next_chapter,omitempty"`
}

func FormatChapter(chapter Chapter) ChapterFormatter {
	return ChapterFormatter{
		ID:        chapter.ID.String(),
		NovelID:   chapter.NovelID.String(),
		Title:     chapter.Title,
		CreatedAt: chapter.CreatedAt,
		UpdatedAt: chapter.UpdatedAt,
	}
}

func FormatChapters(chapters []Chapter) []ChapterFormatter {
	var chaptersFormatted []ChapterFormatter

	for _, chapter := range chapters {
		chaptersFormatted = append(chaptersFormatted, FormatChapter(chapter))
	}

	return chaptersFormatted
}

func FormatChapterDetail(chapter, previousChapter, nextChapter Chapter, baseUrl string) ChapterDetailFormatter {

	var formattedPrevious, formattedNext ChapterFormatter

	if previousChapter.ID != uuid.Nil {
		formattedPrevious = FormatChapter(previousChapter)
	}

	if nextChapter.ID != uuid.Nil {
		formattedNext = FormatChapter(nextChapter)
	}

	response := ChapterDetailFormatter{
		ID:        chapter.ID.String(),
		NovelID:   chapter.NovelID.String(),
		Title:     chapter.Title,
		Content:   baseUrl + chapter.Content,
		CreatedAt: chapter.CreatedAt,
		UpdatedAt: chapter.UpdatedAt,
	}

	// Include PreviousChapter only if it has a non-nil UUID

	if previousChapter.ID != uuid.Nil {
		response.PreviousChapter = &formattedPrevious
	} else {
		response.PreviousChapter = nil
	}

	// Include NextChapter only if it has a non-nil UUID
	if nextChapter.ID != uuid.Nil {
		response.NextChapter = &formattedNext
	} else {
		response.NextChapter = nil
	}

	return response
}
