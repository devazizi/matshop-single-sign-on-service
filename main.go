package main

import (
	"github.com/joho/godotenv"
	"os"
	"sso/router"

	"github.com/gin-gonic/gin"
)

func main() {

	if _, err := os.Stat(".env"); err == nil {
		godotenv.Load(".env")
	}

	ginEngine := gin.Default()

	router.RegisterV1Routes(ginEngine)

	ginEngine.Run(os.Getenv("GIN_ADDRESS"))
}
