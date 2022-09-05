package main

import (
	"os"
	"sso/adapter/db"
	"sso/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if _, err := os.Stat(".env"); err == nil {
		godotenv.Load(".env")
	}

	db := db.NewDB(os.Getenv("DATABASE_DSN"))

	ginEngine := gin.Default()

	router.RegisterV1Routes(ginEngine, db)

	ginEngine.Run(os.Getenv("GIN_ADDRESS"))
}
