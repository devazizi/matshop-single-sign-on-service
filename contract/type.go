package contract

import (
	"sso/adapter/db"
	"sso/dto"
)

type (
	ValidateUserRegisterUserRequest func(db db.DB, req dto.RegisterUserRequest) error
	ValidateUserLoginRequest        func(req dto.LoginUserRequest) error
)
