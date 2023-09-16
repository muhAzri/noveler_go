package handler

import (
	"fmt"
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
		fmt.Println("HANDLE ERROR")
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
		// c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		fmt.Println("HANDLE ERROR")
		return
	}

	_, err = h.genreService.CreateGenre(input)

	if err != nil {
		// c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		fmt.Println("HANDLE ERROR")
		return
	}

	c.Redirect(http.StatusFound, "/genre")

}

func (h *genreHandler) Delete(c *gin.Context) {
	var input genre.FindByIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		fmt.Println("HANDLE ERROR")
		return
	}

	err = h.genreService.DeleteGenre(input)

	if err != nil {
		fmt.Println("HANDLE ERROR")
		return
	}

	c.Redirect(http.StatusFound, "/genre")
}

func (h *genreHandler) Edit(c *gin.Context) {
	var input genre.FindByIDInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		fmt.Println("HANDLE ERROR")
		return
	}

	genre, err := h.genreService.GetGenreByID(input)

	if err != nil {
		fmt.Println("HANDLE ERROR")
		return
	}

	if genre.ID == uuid.Nil {
		fmt.Println("HANDLE ERROR")
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
		// c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		fmt.Println("HANDLE ERROR")
		return
	}

	err = c.ShouldBind(&input)
	if err != nil {
		// c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		fmt.Println("HANDLE ERROR")
		return
	}

	_, err = h.genreService.UpdateGenre(inputId, input)

	if err != nil {
		// c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		fmt.Println("HANDLE ERROR")
		return
	}

	c.Redirect(http.StatusFound, "/genre")

}

func (h *genreHandler) NovelIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "novel_index.html", nil)
}
