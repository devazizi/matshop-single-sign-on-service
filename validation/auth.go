package validation

import (
	"sso/dto"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateUserRegisterUserRequest(req dto.RegisterUserRequest) error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.CellNumber, validation.Required),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.FirstName, validation.Required),
		validation.Field(&req.LastName, validation.Required),
		validation.Field(&req.CellNumber, validation.Required),
	)
}
