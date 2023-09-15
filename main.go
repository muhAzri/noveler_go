package main

import (
	"log"
	"noveler_go/genre"
	"noveler_go/handler"
	"noveler_go/helper"

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

	// Routes
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/genre", genreHandler.CreateGenre)

	router.Run(":8080")
}
