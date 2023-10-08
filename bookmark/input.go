package bookmark

type CreateBookmarkInput struct {
	NovelID string `json:"novel_id"`
}

type DeleteBookmarkInput struct {
	BookmarkID string `json:"bookmark_id"`
}
