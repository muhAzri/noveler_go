package handler

import (
	"net/http"
	"noveler_go/genre"
	"noveler_go/helper"

	"github.com/gin-gonic/gin"
)

type genreHandler struct {
	service genre.Service
}

func NewGenreHandler(genreService genre.Service) *genreHandler {
	return &genreHandler{service: genreService}
}

// / API Endpoint POST /api/v1/genres
func (h *genreHandler) CreateGenre(c *gin.Context) {
	var input genre.CreateGenreInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)

		response := helper.ApiResponse("Create Genre Error", http.StatusBadRequest, "error", nil, errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newGenre, err := h.service.CreateGenre(input)

	if err != nil {
		response := helper.ApiResponse("Create Genre Error", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Genre has been created", http.StatusOK, "success", newGenre, nil)
	c.JSON(http.StatusOK, response)

}

func (h *genreHandler) GetAllGenre(c *gin.Context) {
	genres, err := h.service.GetAllGenres()

	if err != nil {
		response := helper.ApiResponse("Get Genre Failed", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := genre.FormatGenres(genres)
	response := helper.ApiResponse("Genre success getted", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)
}
