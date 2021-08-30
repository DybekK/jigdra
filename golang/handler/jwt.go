package handler

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokenPair(objectid string) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identitykey"] = objectid
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()
	t, err := token.SignedString([]byte("test"))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["exp"] = time.Now().Add(time.Hour * 10).Unix()
	rtClaims["identitykey"] = objectid

	rt, err := refreshToken.SignedString([]byte("test"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}
