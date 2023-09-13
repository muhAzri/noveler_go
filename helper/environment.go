package helper

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializeEnvironment() error {
	environment := os.Getenv("env")
	err := godotenv.Load(".env." + environment)
	if err != nil {
		return err
	}

	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)

	return nil
}
