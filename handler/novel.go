package handler

import (
	"net/http"
	"noveler_go/helper"
	"noveler_go/novel"

	"github.com/gin-gonic/gin"
)

type novelHandler struct {
	service novel.Service
}

func NewNovelHandler(novelService novel.Service) *novelHandler {
	return &novelHandler{service: novelService}
}

// API Endpoint GET /api/v1/novels/newest
func (h *novelHandler) NewestNovel(c *gin.Context) {
	novels, err := h.service.GetNewestNovel()

	if err != nil {
		response := helper.ApiResponse("Newest Novel Fetch Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := novel.FormatNovels(novels)
	response := helper.ApiResponse("Newest Novel Fetched Successfully", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)
}

// API Endpoint GET /api/v1/novels/updated
func (h *novelHandler) UpdatedNovel(c *gin.Context) {
	novels, err := h.service.GetNewlyUpdatedNovel()

	if err != nil {
		response := helper.ApiResponse("Newest Updated Novel Fetch Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := novel.FormatNovels(novels)
	response := helper.ApiResponse("Newest Updated Novel Fetched Successfully", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)
}

// API Endpoint GET /api/v1/novels/best
func (h *novelHandler) BestNovel(c *gin.Context) {
	novels, err := h.service.GetSortByRateNovel()

	if err != nil {
		response := helper.ApiResponse("Best Novel Fetch Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := novel.FormatNovels(novels)
	response := helper.ApiResponse("Best Novel Fetched Successfully", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)
}
