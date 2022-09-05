package validation

import (
	"sso/adapter/db"
	"sso/dto"
	"sso/validation/rule"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateUserRegisterUserRequest(db db.DB, req dto.RegisterUserRequest) error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Password, validation.Required, validation.Length(8, 40)),
		validation.Field(&req.CellNumber, validation.Required, validation.By(rule.CellNumberMustUnique(db))),
		validation.Field(&req.Email, validation.Required, is.Email, validation.By(rule.EmailMustUnique(db))),
		validation.Field(&req.FirstName, validation.Required),
		validation.Field(&req.LastName, validation.Required),
		validation.Field(&req.NationalCode, validation.Required, validation.By(rule.NationalCodeMustUnique(db))),
	)
}

func ValidateUserLoginRequest(req dto.LoginUserRequest) error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 40)),
	)
}
