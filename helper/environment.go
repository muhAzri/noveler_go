package helper

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializeEnvironment() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)

	return nil
}
