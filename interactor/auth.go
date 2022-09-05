package interactor

import (
	"context"
	"sso/adapter/db"
	"sso/dto"
	"sso/service/helper"
	"sso/service/jwt_service"
	"time"
)

func (i Interactor) RegisterClient(ctx context.Context, req dto.RegisterUserRequest) (dto.RegisterUserResponse, error) {

	user := i.store.RegisterUser(db.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		NationalCode: req.NationalCode,
		CellNumber:   req.CellNumber,
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
