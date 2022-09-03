package router

import "github.com/gin-gonic/gin"

func RegisterV1Routes(router *gin.Engine) {
	router.GET("/app", func(context *gin.Context) {
		context.JSON(200, map[string]string{"message": "ok"})
	})
}
