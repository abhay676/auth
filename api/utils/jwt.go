package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type SignedDetails struct {
	Email  string
	UserID string
	jwt.StandardClaims
}

var SecretKey = os.Getenv("JWT_SECRET")

func GenerateJWTToken(email, userId string) (string, error) {
	claims := &SignedDetails{
		Email:  email,
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(1)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateJWTToken(signedToken string) (claims *SignedDetails, msg error) {
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		msg = err
		return
	}
	claims, ok := token.Claims.(*SignedDetails)

	if !ok {
		msg = errors.New("token is expired")
		return
	}
	return claims, nil
}