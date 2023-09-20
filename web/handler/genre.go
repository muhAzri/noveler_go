package handler

import (
	"net/http"
	"noveler_go/genre"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type genreHandler struct {
	genreService genre.Service
}

func NewGenreHandler(genreService genre.Service) *genreHandler {
	return &genreHandler{genreService: genreService}
}

func (h *genreHandler) Index(c *gin.Context) {

	genres, err := h.genreService.GetAllGenres()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "genre_index.html", gin.H{
		"genres": genres,
	})
}

func (h *genreHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "genre_new.html", nil)
}

func (h *genreHandler) Create(c *gin.Context) {
	var input genre.CreateGenreInput

	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	_, err = h.genreService.CreateGenre(input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/genre")

}

func (h *genreHandler) Delete(c *gin.Context) {
	var input genre.FindByIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	err = h.genreService.DeleteGenre(input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/genre")
}

func (h *genreHandler) Edit(c *gin.Context) {
	var input genre.FindByIDInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	genre, err := h.genreService.GetGenreByID(input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	if genre.ID == uuid.Nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "genre_edit.html", gin.H{
		"genre": genre,
	})
}

func (h *genreHandler) Update(c *gin.Context) {
	var input genre.CreateGenreInput
	var inputId genre.FindByIDInput

	err := c.ShouldBindUri(&inputId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	err = c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	_, err = h.genreService.UpdateGenre(inputId, input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/genre")

}
