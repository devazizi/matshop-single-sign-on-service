package contract

import "sso/adapter/db"

type Repository interface {
	RegisterUser(user db.User) db.User
	GenerateToken(token db.OauthAccessToken) db.OauthAccessToken
}
