package db

import (
	"errors"
	"gorm.io/gorm"
)

func (db DB) RegisterUser(user User) User {
	db.store.Create(&user)

	return user
}

func (db DB) GenerateToken(token OauthAccessToken) OauthAccessToken {
	db.store.Create(&token)

	return token
}

func (db DB) CheckUserExists(email string) (User, error) {
	var user User
	result := db.store.Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("not match or credential")
	}

	return user, nil
}

func (db DB) CheckEmailIsUnique(email string) bool {
	var user User
	result := db.store.Select("id").Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}

func (db DB) CheckCellNumberIsUnique(cellNumber string) bool {
	var user User
	result := db.store.Select("id").Where("cell_number = ?", cellNumber).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}

func (db DB) CheckNationalCodeIsUnique(nationalCode string) bool {
	var user User
	result := db.store.Select("id").Where("national_code = ?", nationalCode).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}
