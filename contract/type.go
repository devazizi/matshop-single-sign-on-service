package contract

import "sso/dto"

type (
	ValidateUserRegisterUserRequest func(req dto.RegisterUserRequest) error
)
