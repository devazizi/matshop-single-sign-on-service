package contract

import "sso/adapter/db"

type Repository interface {
	RegisterUser(user db.User) db.User
	GenerateToken(token db.OauthAccessToken) db.OauthAccessToken
	CheckUserExists(email string) (db.User, error)
	CheckEmailIsUnique(email string) bool
	CheckCellNumberIsUnique(cellNumber string) bool
	CheckNationalCodeIsUnique(nationalCode string) bool
}
