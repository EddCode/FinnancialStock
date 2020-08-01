package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/eddcode/fintechApp/model"
)

var jwtSecretKey = []byte("1e5b70471caaf625b8821e18364c93dc6466e53af0203d4018d07ab69d0354ca5eeaf9497b4f2563778edee2aa17e8b7660a7bc384f050fb55205f0d3ad5161a")

func SingToken(email string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &model.Claim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(jwtSecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	claims := &model.Claim{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(a *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		},
	)

	if token != nil {
		return claims.Email, nil
	}

	return "", err

}
