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
	ClientId uint `json:"client_id"`
}

var secret = []byte("secret key for my application it is just a secret")

func CreateToken(tokenId string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	exp := time.Now().Add(time.Hour * 24 * 30)

	privateKey, _ := ioutil.ReadFile(os.Getenv("PRIVATE_KEY_PATH"))
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	token.Claims = &MyJWTClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			ID:        tokenId,
		},
		433,
	}

	fmt.Println(signKey)
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
