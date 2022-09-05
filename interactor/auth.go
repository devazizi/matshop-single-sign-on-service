package interactor

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"sso/adapter/db"
	"sso/dto"
	"sso/service/helper"
	"sso/service/jwt_service"
	"time"
)

func (i Interactor) RegisterClient(ctx context.Context, req dto.RegisterUserRequest) (dto.RegisterUserResponse, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 1)

	if err != nil {
		return dto.RegisterUserResponse{}, err
	}

	user := i.store.RegisterUser(db.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		NationalCode: req.NationalCode,
		CellNumber:   req.CellNumber,
		Password:     string(password),
	})

	randomString := helper.RandomStringGenerator(100)
	audience := "client_authentication"
	accessTokenExpiryDate := time.Now().AddDate(0, 0, 30)
	token := i.store.GenerateToken(db.OauthAccessToken{
		Token:      randomString,
		UserId:     user.ID,
		Audience:   audience,
		ExpiryDate: accessTokenExpiryDate,
	})

	jwtToken, err := jwt_service.CreateToken(token.Token, user.ID, audience, accessTokenExpiryDate)
	if err != nil {
		return dto.RegisterUserResponse{}, err
	}

	return dto.RegisterUserResponse{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		CellNumber:   user.CellNumber,
		Email:        user.Email,
		NationalCode: user.NationalCode,
		AccessToken:  jwtToken,
	}, nil
}

func (i Interactor) LoginClient(req dto.LoginUserRequest) (dto.LoginUserResponse, error) {

	user, err := i.store.CheckUserExists(req.Email)
	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	if passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); passwordErr != nil {
		return dto.LoginUserResponse{}, errors.New("not match our credential")
	}

	randomString := helper.RandomStringGenerator(100)
	audience := "client_authentication"
	accessTokenExpiryDate := time.Now().AddDate(0, 0, 30)
	token := i.store.GenerateToken(db.OauthAccessToken{
		Token:      randomString,
		UserId:     user.ID,
		Audience:   audience,
		ExpiryDate: accessTokenExpiryDate,
	})

	jwtToken, err := jwt_service.CreateToken(token.Token, user.ID, audience, accessTokenExpiryDate)
	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	return dto.LoginUserResponse{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		CellNumber:   user.CellNumber,
		Email:        user.Email,
		NationalCode: user.NationalCode,
		AccessToken:  jwtToken,
	}, nil
}
