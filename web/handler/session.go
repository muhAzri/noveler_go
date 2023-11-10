package handler

import (
	"net/http"
	"noveler_go/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	userService user.Service
}

func NewSessionHandler(userService user.Service) *sessionHandler {
	return &sessionHandler{userService: userService}
}

func (h *sessionHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "session_new.html", nil)
}

func (h *sessionHandler) Create(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.HTML(http.StatusFound, "session_new.html", gin.H{
			"error": "Email or Password is incorrect",
		})
		return
	}

	user, err := h.userService.Login(input)
	if err != nil || user.Role != "admin" {
		c.HTML(http.StatusFound, "session_new.html", gin.H{
			"error": "Email or Password is incorrect",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID.String())
	session.Set("email", user.Email)
	session.Save()

	c.Redirect(http.StatusFound, "/novel")
}
