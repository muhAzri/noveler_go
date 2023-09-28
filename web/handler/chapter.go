package handler

import (
	"fmt"
	"net/http"
	"noveler_go/chapter"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type chapterHandler struct {
	chapterService chapter.Service
}

func NewChapterHandler(chapterService chapter.Service) *chapterHandler {
	return &chapterHandler{chapterService: chapterService}
}

func (h *chapterHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "chapter_new.html", nil)
}

func (h *chapterHandler) Create(c *gin.Context) {
	var input chapter.CreateChapterInput
	var inputID chapter.FindByIDInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	err = c.ShouldBindUri(&inputID)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{"error": err.Error()})
		return
	}

	title := strings.ReplaceAll(strings.ToLower(input.Title), " ", "-")

	// Generate a unique value (timestamp or random string)
	uniqueValue := strconv.FormatInt(time.Now().UnixNano(), 10)

	// Concatenate the unique value to the file name
	fileName := title + "-" + inputID.ID + "-" + uniqueValue + ".html"
	filePath := "static/chapters/" + fileName

	file, err := os.Create(filePath)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	_, err = file.WriteString(input.Content)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	input.Content = filePath

	_, err = h.chapterService.CreateChapter(inputID, input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/novel/"+inputID.ID)
}

func (h *chapterHandler) Edit(c *gin.Context) {
	var inputID chapter.FindByIDInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	chapter, err := h.chapterService.GetByID(inputID)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	htmlBytes, err := os.ReadFile(chapter.Content)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	htmlContent := string(htmlBytes)

	c.HTML(http.StatusOK, "chapter_edit.html", gin.H{
		"chapter": chapter,
		"content": htmlContent,
	})
}

func (h *chapterHandler) Update(c *gin.Context) {
	var inputID chapter.FindByIDInput
	var input chapter.CreateChapterInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{"error": err.Error()})
		return
	}

	err = c.ShouldBind(&input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	chapter, err := h.chapterService.GetByID(inputID)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	file, err := os.Create(chapter.Content)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	_, err = file.WriteString(input.Content)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	input.Content = chapter.Content

	_, err = h.chapterService.UpdateChapter(inputID, input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/novel/"+chapter.NovelID.String())

}

func (h *chapterHandler) Delete(c *gin.Context) {
	var input chapter.FindByIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	chapter, err := h.chapterService.GetByID(input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	if chapter.ID == uuid.Nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	_ = os.Remove(chapter.Content)

	err = h.chapterService.Delete(input)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

}
