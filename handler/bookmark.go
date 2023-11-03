package handler

import (
	"net/http"
	"noveler_go/bookmark"
	"noveler_go/helper"
	"noveler_go/user"

	"github.com/gin-gonic/gin"
)

type bookmarkHandler struct {
	bookmarkService bookmark.Service
}

func NewBookmarkHandler(bookmarkService bookmark.Service) *bookmarkHandler {
	return &bookmarkHandler{bookmarkService: bookmarkService}
}

func (h *bookmarkHandler) AddOrRemoveBookmark(c *gin.Context) {
	var inputID bookmark.CreateBookmarkInput

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Add Or Remove Bookmark Failed", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	bookmarked, err := h.bookmarkService.FindByUserAndNovelID(currentUser.ID.String(), inputID.NovelID)

	if err != nil {
		response := helper.ApiResponse("Add Or Remove Bookmark Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if bookmarked {
		err = h.bookmarkService.DeleteBookmark(inputID, currentUser.ID.String())

		if err != nil {
			response := helper.ApiResponse(" Remove Bookmark Failed", http.StatusInternalServerError, "error", nil, err.Error())
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		response := helper.ApiResponse(" Remove Bookmark Success", http.StatusOK, "success", gin.H{
			"success": true,
		}, nil)
		c.JSON(http.StatusOK, response)
		return
	}

	err = h.bookmarkService.CreateBookmark(inputID, currentUser.ID.String())

	if err != nil {
		response := helper.ApiResponse("Add  Bookmark Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("Add Bookmark Success", http.StatusOK, "success", gin.H{
		"success": true,
	}, nil)
	c.JSON(http.StatusOK, response)
}

func (h *bookmarkHandler) GetUserBookmarks(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	bookmarks, err := h.bookmarkService.FindBookmarksByUserID(currentUser.ID.String())

	if err != nil {
		response := helper.ApiResponse("Get Bookmarks Failure", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := bookmark.FormatBookmarkNovels(bookmarks)
	response := helper.ApiResponse("Get Bookmarks Sucess", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)
}
