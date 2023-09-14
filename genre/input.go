package genre

type CreateGenreInput struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type FindByIDInput struct {
	ID string `uri:"id" binding:"required"`
}
