package novel

type CreateNovelInput struct {
	Title       string   `json:"title" form:"title" binding:"required"`
	Description string   `json:"description" form:"description" binding:"required"`
	CoverImage  string   `json:"cover_image" form:"cover_image" binding:"required"`
	Status      string   `json:"status" form:"status" binding:"required"`
	Author      string   `json:"author" form:"author" binding:"required"`
	Rating      int      `json:"rating" form:"rating" binding:"required"`
	GenreIDs    []string `json:"genre_ids" form:"genre_ids" binding:"required"`
}

type FindByIDInput struct {
	ID string `uri:"id" binding:"required"`
}
