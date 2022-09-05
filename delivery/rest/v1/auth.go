package v1

import (
	"net/http"
	"sso/adapter/db"
	"sso/contract"
	"sso/dto"
	"sso/interactor"

	"github.com/gin-gonic/gin"
)

func RegisterUser(db db.DB, validator contract.ValidateUserRegisterUserRequest) gin.HandlerFunc {

	return func(c *gin.Context) {
		requestBody := dto.RegisterUserRequest{}

		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, dto.BaseResponse{Message: err.Error()})
			return
		}

		if validationErr := validator(db, requestBody); validationErr != nil {
			c.JSON(http.StatusBadRequest, dto.BaseResponse{Message: validationErr.Error()})
			return
		}

		registerResponse, err := interactor.NewDB(db).RegisterClient(c.Request.Context(), requestBody)

		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.BaseResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, dto.BaseResponse{Status: true, Data: registerResponse})
	}
}

func LoginUser(db db.DB, validator contract.ValidateUserLoginRequest) gin.HandlerFunc {

	return func(c *gin.Context) {
		loginRequest := dto.LoginUserRequest{}

		if err := c.BindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, dto.BaseResponse{Message: err.Error()})
			return
		}

		if validationErr := validator(loginRequest); validationErr != nil {
			c.JSON(http.StatusUnprocessableEntity, dto.BaseResponse{Message: validationErr.Error()})
			return
		}

		loginResponse, err := interactor.NewDB(db).LoginClient(loginRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.BaseResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.BaseResponse{Data: loginResponse, Status: true})
	}
}
