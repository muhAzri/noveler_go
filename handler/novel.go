package handler

import (
	"net/http"
	"noveler_go/bookmark"
	"noveler_go/chapter"
	"noveler_go/genre"
	"noveler_go/helper"
	"noveler_go/novel"
	"noveler_go/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type novelHandler struct {
	novelService    novel.Service
	bookmarkService bookmark.Service
	chapterService  chapter.Service
	genreService    genre.Service
}

func NewNovelHandler(novelService novel.Service, bookmarkService bookmark.Service, chapterService chapter.Service, genreService genre.Service) *novelHandler {
	return &novelHandler{novelService: novelService, bookmarkService: bookmarkService, chapterService: chapterService, genreService: genreService}
}

// API Endpoint GET /api/v1/novels/newest
func (h *novelHandler) NewestNovel(c *gin.Context) {
	novels, err := h.novelService.GetNewestNovel()

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
	novels, err := h.novelService.GetNewlyUpdatedNovel()

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
	novels, err := h.novelService.GetSortByRateNovel()

	if err != nil {
		response := helper.ApiResponse("Best Novel Fetch Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := novel.FormatNovels(novels)
	response := helper.ApiResponse("Best Novel Fetched Successfully", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)
}

func (h *novelHandler) RecommendedNovel(c *gin.Context) {
	recommendedNovel, err := h.novelService.GetRandomNovel()

	if err != nil {
		response := helper.ApiResponse("Get Detail Novel Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	recommendedFormatter := novel.FormatNovels(recommendedNovel)

	response := helper.ApiResponse("Recommended Novel Fetched SuccessFully", http.StatusOK, "success", recommendedFormatter, nil)
	c.JSON(http.StatusOK, response)

}

func (h *novelHandler) DetailNovel(c *gin.Context) {
	var inputID novel.FindByIDInput

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Get Detail Novel Failed", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	if err != nil {
		response := helper.ApiResponse("Get Detail Novel Failed", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	gettedNovel, err := h.novelService.GetNovelByID(inputID)

	if err != nil {
		response := helper.ApiResponse("Get Detail Novel Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	genreNamed, err := h.genreService.GetGenreByIDS(gettedNovel.GenreIDs)

	if err != nil {
		response := helper.ApiResponse("Get Detail Novel Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if err != nil {
		response := helper.ApiResponse("Get Detail Novel Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	chapters, err := h.chapterService.GetChaptersByID(inputID.ID)

	if err != nil {
		response := helper.ApiResponse("Get Detail Novel Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	bookmarked, err := h.bookmarkService.FindByUserAndNovelID(currentUser.ID.String(), inputID.ID)

	if err != nil {
		response := helper.ApiResponse("Get Detail Novel Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	genreFormatter := genre.FormatGenres(genreNamed)
	formatter := novel.FormatDetailNovel(gettedNovel, genreFormatter, len(chapters), bookmarked)
	response := helper.ApiResponse("Detail Novel Fetched SuccessFully", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)
}

// API Endpoint GET /api/v1/novels/search
func (h *novelHandler) SearchNovels(c *gin.Context) {
	var searchInput novel.NovelSearchParametersInput

	// Extract search parameters from query parameters
	searchInput.Title = c.DefaultQuery("title", "")
	searchInput.Status = c.DefaultQuery("status", "")

	genreQueryParam := c.Query("genres")
	if genreQueryParam != "" {
		searchInput.Genres = []string{genreQueryParam}
	} else {
		searchInput.Genres = nil
	}

	// Extract pagination parameters from query string with default values
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	searchInput.Page = page
	searchInput.PageSize = pageSize

	novels, err := h.novelService.SearchNovels(searchInput)

	if err != nil {
		response := helper.ApiResponse("Search Novels Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := novel.FormatSearchNovels(novels)
	response := helper.ApiResponse("Novels Search Successful", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)
}
