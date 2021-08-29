package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthMiddleware struct{}

func (a *AuthMiddleware) GenerateTokenPair(objectid string) (map[string]string, error) {
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

func (auth *AuthMiddleware) TokenValid(r *http.Request) error {
	token, err := auth.VerifyToken(r)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return err
	}
	id := claims["identitykey"].(string)
	fmt.Println(id)
	user := workspaceUserService.FindByUserId(id)
	if user != nil {
		return nil
	}
	resp, err := http.Get("http://localhost:4201/v1/user/" + id)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("user does not exist")
	}
	type respUser struct {
		Username string `json:"username"`
	}
	var ru respUser
	err = json.NewDecoder(resp.Body).Decode(&ru)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("username", ru.Username)

	_, err = workspaceUserService.SaveWorkspaceUser(id, ru.Username)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return err
}

func (auth *AuthMiddleware) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("test"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (auth *AuthMiddleware) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, bson.M{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
