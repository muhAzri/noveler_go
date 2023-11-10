package handler

import (
	"net/http"
	"noveler_go/chapter"
	"noveler_go/helper"

	"github.com/gin-gonic/gin"
)

type chapterHandler struct {
	chapterService chapter.Service
	baseUrl        string
}

func NewChapterHandler(chapterService chapter.Service, baseUrl string) *chapterHandler {
	return &chapterHandler{chapterService: chapterService, baseUrl: baseUrl}
}

func (h *chapterHandler) GetChapterListFromNovelID(c *gin.Context) {
	var inputID chapter.FindByIDInput

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Get Chapter List From Novel ID Failed", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	chapters, err := h.chapterService.GetChaptersByID(inputID.ID)

	if err != nil {
		response := helper.ApiResponse("Get Chapter List From Novel ID Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := chapter.FormatChapters(chapters)
	response := helper.ApiResponse("Get Chapter List From Novel ID Success", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)
}
func (h *chapterHandler) GetChapterDetail(c *gin.Context) {
	var inputID chapter.FindByIDInput

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Get Chapter Detail Failed", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get the current chapter
	currentChapter, err := h.chapterService.GetByID(inputID)

	if err != nil {
		response := helper.ApiResponse("Get Chapter Detail Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Get all chapters for the same novel
	chapters, err := h.chapterService.GetChaptersByID(currentChapter.NovelID.String())

	if err != nil {
		response := helper.ApiResponse("Get Chapter Detail Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Find the index of the current chapter in the list
	currentIndex := -1
	for i, chap := range chapters {
		if chap.ID == currentChapter.ID {
			currentIndex = i
			break
		}
	}

	// Calculate the indices of the previous and next chapters
	previousIndex := currentIndex - 1
	nextIndex := currentIndex + 1

	// Ensure indices are within bounds
	if previousIndex < 0 {
		previousIndex = -1 // Indicate no previous chapter
	}
	if nextIndex >= len(chapters) {
		nextIndex = -1 // Indicate no next chapter
	}

	// Retrieve the previous and next chapters or set them to empty chapters
	var previousChapter, nextChapter chapter.Chapter

	if previousIndex != -1 {
		previousChapter = chapters[previousIndex]
	}

	if nextIndex != -1 {
		nextChapter = chapters[nextIndex]
	}

	formatter := chapter.FormatChapterDetail(currentChapter, previousChapter, nextChapter, h.baseUrl)
	response := helper.ApiResponse("Get Chapter Detail Success", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)
}
