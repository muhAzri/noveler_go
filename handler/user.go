package handler

import (
	"net/http"
	"noveler_go/auth"
	"noveler_go/helper"
	"noveler_go/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{service: userService, authService: authService}
}

// / API Endpoint POST /api/v1/register
func (h *userHandler) Register(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)

		response := helper.ApiResponse("Register User Failed", http.StatusBadRequest, "error", nil, errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.service.Register(input)

	if err != nil {
		response := helper.ApiResponse("Register User Failed", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	accessToken, refreshToken, err := h.authService.GenerateToken(newUser.ID.String())

	if err != nil {
		response := helper.ApiResponse("Register User Failed", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := user.FormatUser(newUser, accessToken, refreshToken)

	response := helper.ApiResponse("Register Success", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)

}

// API Endpoint POST /api/v1/sessions
func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		response := helper.ApiResponse("Login failed", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedinUser, err := h.service.Login(input)

	if err != nil {

		response := helper.ApiResponse("Login failed", http.StatusUnauthorized, "error", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	accessToken, refreshToken, err := h.authService.GenerateToken(loggedinUser.ID.String())

	if err != nil {
		response := helper.ApiResponse("Login failed", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, accessToken, refreshToken)

	response := helper.ApiResponse("Succesfully logged in", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Refresh(c *gin.Context) {
	var input user.RefreshInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.ApiResponse("Failed to refresh sessions", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	accessToken, err := h.authService.RefreshToken(input.RefreshToken)

	if err != nil {
		response := helper.ApiResponse("Failed to refresh sessions", http.StatusUnauthorized, "error", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	formatter := auth.FormatToken(accessToken)

	response := helper.ApiResponse("Session successfully refreshed", http.StatusOK, "success", formatter, nil)
	c.JSON(http.StatusOK, response)
}
