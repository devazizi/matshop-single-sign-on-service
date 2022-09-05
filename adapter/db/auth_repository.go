package db

func (db DB) RegisterUser(user User) User {
	db.store.Create(&user)

	return user
}

func (db DB) GenerateToken(token OauthAccessToken) OauthAccessToken {
	db.store.Create(&token)

	return token
}
