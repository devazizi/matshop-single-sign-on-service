package router

import (
	"sso/adapter/db"
	v1 "sso/delivery/rest/v1"
	"sso/validation"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.Engine, db db.DB) {

	restV1 := router.Group("/api/v1")
	{
		authRest := restV1.Group("/auth")
		{
			authRest.POST("/register", v1.RegisterUser(db, validation.ValidateUserRegisterUserRequest))
			authRest.POST("/login", v1.LoginUser(db, validation.ValidateUserLoginRequest))
		}
	}
}
