package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-psql/workspace"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthMiddleware struct {
	workspaceUserService workspace.WorkspaceUserService
}

//factory

func NewAuthMiddleware(workspaceUserService workspace.WorkspaceUserService) AuthMiddleware {
	return AuthMiddleware{workspaceUserService: workspaceUserService}
}

//methods

func (auth *AuthMiddleware) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.tokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, bson.M{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (auth *AuthMiddleware) tokenValid(r *http.Request) error {
	token, err := auth.verifyToken(r)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return err
	}
	id := claims["identitykey"].(string)
	_, err = auth.workspaceUserService.GetUser(id)

	if errors.Is(err, pgx.ErrNoRows) {
		resp, err := http.Get("http://localhost:4201/v1/user/" + id)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return fmt.Errorf("user does not exist")
		}
		type respBody struct {
			Username string
		}
		var rB respBody
		err = json.NewDecoder(resp.Body).Decode(&rB)
		if err != nil {
			return err
		}
		err = auth.workspaceUserService.CreateUser(id, rB.Username)
		return err
	}

	return err
}

func (auth *AuthMiddleware) verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := auth.extractToken(r)
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

func (auth *AuthMiddleware) extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
