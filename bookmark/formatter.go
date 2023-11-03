package bookmark

import "noveler_go/novel"

type BookmarkNovelFormatter struct {
	BookmarkID string      `json:"id"`
	Novel      novel.NovelFormatter `json:"novel"`
}

func FormatBookmarkNovel(bookmark Bookmark) BookmarkNovelFormatter {
    novelFormatter := novel.FormatNovel(bookmark.Novel)

    formatter := BookmarkNovelFormatter{
        BookmarkID: string(bookmark.ID.String()),
        Novel: novelFormatter,
    }

    return formatter
}

func FormatBookmarkNovels(bookmarks []Bookmark) []BookmarkNovelFormatter {
    bookmarkFormatters := []BookmarkNovelFormatter{}

    for _, bookmark := range bookmarks {
        formatter := FormatBookmarkNovel(bookmark)
        bookmarkFormatters = append(bookmarkFormatters, formatter)
    }

    return bookmarkFormatters
}
