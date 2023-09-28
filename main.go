package main

import (
	"fmt"
	"log"
	"noveler_go/chapter"
	"noveler_go/genre"
	"noveler_go/handler"
	"noveler_go/helper"
	"noveler_go/novel"
	"path/filepath"

	webHandler "noveler_go/web/handler"

	"github.com/gin-contrib/multitemplate"
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
	genreHandler := handler.NewGenreHandler(genreService)

	// Novel
	novelRepository := novel.NewRepository(db)
	novelService := novel.NewService(novelRepository)

	//Chapter
	chapterRepository := chapter.NewRepository(db)
	chapterService := chapter.NewService(chapterRepository)

	//CMS Handler
	genreAdminHandler := webHandler.NewGenreHandler(genreService)
	novelAdminHandler := webHandler.NewNovelHandler(novelService, genreService)
	chapterAdminHandler := webHandler.NewChapterHandler(chapterService)
	router := gin.Default()

	// Load HTML & Static Assets
	router.LoadHTMLGlob("web/templates/**/*")
	router.HTMLRender = loadTemplates("./web/templates")
	router.Static("/assets", "./web/assets/")
	router.Static("/static", "./static/")

	//API Routes
	api := router.Group("/api/v1")
	api.POST("/genre", genreHandler.CreateGenre)

	// CMS routes
	router.GET("/", novelAdminHandler.Index)
	router.GET("/genre", genreAdminHandler.Index)
	router.GET("/genre/new", genreAdminHandler.New)
	router.POST("/genre/new", genreAdminHandler.Create)
	router.POST("/genre/:id/delete", genreAdminHandler.Delete)
	router.GET("/genre/:id/edit", genreAdminHandler.Edit)
	router.POST("/genre/:id/edit", genreAdminHandler.Update)
	router.GET("/novel", novelAdminHandler.Index)
	router.GET("/novel/new", novelAdminHandler.New)
	router.POST("/novel/create", novelAdminHandler.Create)
	router.GET("/novel/:id", novelAdminHandler.Detail)
	router.GET("/novel/:id/edit", novelAdminHandler.Edit)
	router.POST("/novel/:id/edit", novelAdminHandler.Update)
	router.GET("/chapter/:id/new", chapterAdminHandler.New)
	router.POST("/chapter/:id/new", chapterAdminHandler.Create)
	router.GET("/chapter/:id/edit", chapterAdminHandler.Edit)
	router.POST("/chapter/:id/edit", chapterAdminHandler.Update)
	router.DELETE("/chapter/:id/delete", chapterAdminHandler.Delete)

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
