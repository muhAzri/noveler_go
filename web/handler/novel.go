package handler

import (
	"net/http"
	"noveler_go/genre"
	"noveler_go/novel"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type novelHandler struct {
	novelService novel.Service
	genreService genre.Service
}

func NewNovelHandler(novelService novel.Service, genreService genre.Service) *novelHandler {
	return &novelHandler{novelService: novelService, genreService: genreService}
}

func (h *novelHandler) Index(c *gin.Context) {
	novels, err := h.novelService.GetAllNovel()

	if err != nil {
		c.HTML(http.StatusOK, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "novel_index.html", gin.H{
		"novels": novels,
	})
}

func (h *novelHandler) New(c *gin.Context) {
	genres, err := h.genreService.GetAllGenres()

	if err != nil {
		c.HTML(http.StatusOK, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "novel_new.html", gin.H{
		"genres": genres,
	})
}

func (h *novelHandler) Create(c *gin.Context) {
	var input novel.CreateNovelInput

	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	_, err = h.novelService.CreateNovel(input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/novel")

}

func (h *novelHandler) Detail(c *gin.Context) {
	var input novel.FindByIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	novel, err := h.novelService.GetNovelByID(input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	if novel.ID == uuid.Nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{"error": err.Error()})
		return
	}

	genres, err := h.genreService.GetGenreByIDS(novel.GenreIDs)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "novel_detail.html", gin.H{
		"novel":  novel,
		"genres": genres,
	})
}

func (h *novelHandler) Edit(c *gin.Context) {
	var input novel.FindByIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	genres, err := h.genreService.GetAllGenres()

	if err != nil {
		c.HTML(http.StatusOK, "500.html", gin.H{"error": err.Error()})
		return
	}

	novel, err := h.novelService.GetNovelByID(input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	if novel.ID == uuid.Nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "novel_edit.html", gin.H{
		"novel":  novel,
		"genres": genres,
	})
}

func (h *novelHandler) Update(c *gin.Context) {
	var input novel.CreateNovelInput
	var inputID novel.FindByIDInput

	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	err = c.ShouldBindUri(&inputID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	novel, err := h.novelService.GetNovelByID(inputID)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	if novel.ID == uuid.Nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{"error": err.Error()})
		return
	}

	_, err = h.novelService.UpdateNovel(inputID, input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/novel")

}
