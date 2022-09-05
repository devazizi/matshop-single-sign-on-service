package jwt_service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"os"
	"time"
)

type MyJWTClaims struct {
	*jwt.RegisteredClaims
	ClientId uint   `json:"client_id"`
	Audience string `json:"audience"`
}

func CreateToken(tokenId string, clientId uint, audience string, expiryDate time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	privateKey, _ := ioutil.ReadFile(os.Getenv("PRIVATE_KEY_PATH"))
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	token.Claims = &MyJWTClaims{
		&jwt.RegisteredClaims{
			Subject:   "access_token",
			ExpiresAt: jwt.NewNumericDate(expiryDate),
			ID:        tokenId,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		clientId,
		audience,
	}

	val, err := token.SignedString(signKey)

	if err != nil {
		return "", err
	}
	return val, nil
}

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		publicKey, _ := ioutil.ReadFile(os.Getenv("PUBLIC_KEY_PATH"))
		signKey, _ := jwt.ParseRSAPublicKeyFromPEM(publicKey)
		return signKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
