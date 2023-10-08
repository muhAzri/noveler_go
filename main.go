package main

import (
	"fmt"
	"log"
	"noveler_go/auth"
	"noveler_go/chapter"
	"noveler_go/genre"
	"noveler_go/handler"
	"noveler_go/helper"
	"noveler_go/middleware"
	"noveler_go/novel"
	"noveler_go/user"
	"os"
	"path/filepath"

	webHandler "noveler_go/web/handler"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	err := helper.InitializeEnvironment()

	if err != nil {
		log.Fatal(err)
	}

	db, err := helper.InitializeDatabase()

	if err != nil {
		log.Fatal(err)
	}

	helper.AutomigrateDatabase(db)

	genreRepository := genre.NewRepository(db)
	genreService := genre.NewService(genreRepository)

	// Novel
	novelRepository := novel.NewRepository(db)
	novelService := novel.NewService(novelRepository)

	//Chapter
	chapterRepository := chapter.NewRepository(db)
	chapterService := chapter.NewService(chapterRepository)

	//User
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	//JWT SERVICE
	authService := auth.NewService()

	//API HANDLER
	genreHandler := handler.NewGenreHandler(genreService)
	userHandler := handler.NewUserHandler(userService, authService)

	//CMS Handler
	genreAdminHandler := webHandler.NewGenreHandler(genreService)
	novelAdminHandler := webHandler.NewNovelHandler(novelService, genreService)
	chapterAdminHandler := webHandler.NewChapterHandler(chapterService)
	sessionHandler := webHandler.NewSessionHandler(userService)

	// ROUTES
	router := gin.Default()
	cookieStore := cookie.NewStore([]byte(os.Getenv("SECRET_KEY")))
	router.Use(sessions.Sessions("session", cookieStore))

	// Load HTML & Static Assets
	router.LoadHTMLGlob("web/templates/**/*")
	router.HTMLRender = loadTemplates("./web/templates")
	router.Static("/assets", "./web/assets/")
	router.Static("/static", "./static/")

	//API Routes
	api := router.Group("/api/v1")
	api.POST("/genre", genreHandler.CreateGenre)

	//User
	api.POST("/register", userHandler.Register)
	api.POST("/sessions", userHandler.Login)
	api.POST("/sessions/refresh", userHandler.Refresh)

	// CMS routes
	router.GET("/", middleware.AuthAdminMiddleware(), novelAdminHandler.Index)

	router.GET("/genre", middleware.AuthAdminMiddleware(), genreAdminHandler.Index)
	router.GET("/genre/new", middleware.AuthAdminMiddleware(), genreAdminHandler.New)
	router.POST("/genre/new", middleware.AuthAdminMiddleware(), genreAdminHandler.Create)
	router.POST("/genre/:id/delete", middleware.AuthAdminMiddleware(), genreAdminHandler.Delete)
	router.GET("/genre/:id/edit", middleware.AuthAdminMiddleware(), genreAdminHandler.Edit)
	router.POST("/genre/:id/edit", middleware.AuthAdminMiddleware(), genreAdminHandler.Update)

	router.GET("/novel", middleware.AuthAdminMiddleware(), novelAdminHandler.Index)
	router.GET("/novel/new", middleware.AuthAdminMiddleware(), novelAdminHandler.New)
	router.POST("/novel/create", middleware.AuthAdminMiddleware(), novelAdminHandler.Create)
	router.GET("/novel/:id", middleware.AuthAdminMiddleware(), novelAdminHandler.Detail)
	router.GET("/novel/:id/edit", middleware.AuthAdminMiddleware(), novelAdminHandler.Edit)
	router.POST("/novel/:id/edit", middleware.AuthAdminMiddleware(), novelAdminHandler.Update)

	router.GET("/chapter/:id/new", middleware.AuthAdminMiddleware(), chapterAdminHandler.New)
	router.POST("/chapter/:id/new", middleware.AuthAdminMiddleware(), chapterAdminHandler.Create)
	router.GET("/chapter/:id/edit", middleware.AuthAdminMiddleware(), chapterAdminHandler.Edit)
	router.POST("/chapter/:id/edit", middleware.AuthAdminMiddleware(), chapterAdminHandler.Update)
	router.DELETE("/chapter/:id/delete", middleware.AuthAdminMiddleware(), chapterAdminHandler.Delete)

	router.GET("/login", sessionHandler.New)
	router.POST("/session", sessionHandler.Create)

	router.Run(":8080")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		templateName := filepath.Base(include)
		r.AddFromFiles(filepath.Base(include), files...)
		fmt.Println("Loaded template:", templateName)
	}

	return r
}
