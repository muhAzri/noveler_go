package main

import (
	"fmt"
	"log"
	"noveler_go/genre"
	"noveler_go/handler"
	"noveler_go/helper"
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

	//CMS Handler
	genreAdminHandler := webHandler.NewGenreHandler(genreService)
	router := gin.Default()

	// Load HTML & Static Assets
	router.LoadHTMLGlob("web/templates/**/*")
	router.HTMLRender = loadTemplates("./web/templates")
	router.Static("/assets", "./web/assets/")

	//API Routes
	api := router.Group("/api/v1")
	api.POST("/genre", genreHandler.CreateGenre)

	// CMS routes
	router.GET("/", genreAdminHandler.Index)
	router.GET("/genre", genreAdminHandler.Index)
	router.GET("/novel", genreAdminHandler.NovelIndex)

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
