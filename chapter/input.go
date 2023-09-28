package chapter

type CreateChapterInput struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

type FindByIDInput struct {
	ID string `uri:"id" binding:"required"`
}
