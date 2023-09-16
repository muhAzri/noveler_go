package handler

import (
	"fmt"
	"net/http"
	"noveler_go/genre"

	"github.com/gin-gonic/gin"
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
	}

	c.HTML(http.StatusOK, "genre_index.html", gin.H{
		"genres": genres,
	})
}

func (h *genreHandler) NovelIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "novel_index.html", nil)
}
