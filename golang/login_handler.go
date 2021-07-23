package main

import (
	"fmt"
	"golang/model"
	"net/http"
	"net/url"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type handler struct{}

var auth = model.AuthHandler{}

func (h *handler) getUwa(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "success",
	})
}

var validate = validator.New()

func (h *handler) addUser(c *gin.Context) {
	var req_user model.User
	if err := c.BindJSON(&req_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user": "invalid request body"})
		return
	}

	validationError := validate.Struct(req_user)
	if validationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user": "invalid request body"})
		return
	}

	res, err := model.Interface.CreateUser(&req_user, c)

	if err != nil {
		if err.Error() == "409" {
			c.JSON(http.StatusConflict, gin.H{
				"email": "email already in use",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"username": err.Error(),
			})
		}
		return
	}
	q := url.Values{}
	q.Add("redirect", res)
	location := url.URL{Path: "/v1/login", RawQuery: q.Encode()}

	c.Redirect(http.StatusFound, location.RequestURI())

}

func (h *handler) getUserById(c *gin.Context) {

	id := c.Param("id")
	user, err := model.Interface.GetUserById(id, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"user": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, user)
	return

}

func (h *handler) refresh(c *gin.Context) {
	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	tokenReq := tokenReqBody{}
	if err := c.BindJSON(&tokenReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"token": "invalid request body"})
		return
	}
	token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("test"), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"alg": err.Error()})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		newTokenPair, err := auth.GenerateTokenPair(claims["identitykey"].(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"token": "signing error"})
			return
		}

		c.JSON(http.StatusOK, newTokenPair)
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"token": err.Error()})
}

func (h *handler) login(c *gin.Context) {
	if c.Request.Method == "GET" {
		q := c.Query("redirect")
		if q != "" {
			id, exists := model.Interface.VerifyRedirect(c, q)
			if exists != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"redirect": "redirect doesn't exist"})
				return
			}
			objid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"_id": "invalid hex"})
				return
			}
			res := model.Interface.GetCollection("users").FindOne(c, bson.M{"_id": objid})
			var user model.User
			decode_err := res.Decode(&user)
			if decode_err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"user": "failed to decode user"})
				return
			}

			newTokenPair, err := auth.GenerateTokenPair(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"token": "signing error"})
				return
			}

			c.JSON(http.StatusOK, newTokenPair)
			return
		}
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not found"})
	} else if c.Request.Method == "POST" {
		var req_login model.LoginUser
		if err := c.BindJSON(&req_login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"user": "invalid request body"})
			return
		}
		validation_error := validate.Struct(req_login)
		if validation_error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"user": "validation error"})
			return
		}
		user, user_err := model.Interface.GetUser(&req_login, c)
		if user_err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"login": "invalid email or password"})
			return
		}
		newTokenPair, err := auth.GenerateTokenPair(user.Id.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"token": "signing error"})
			return
		}

		c.JSON(http.StatusOK, newTokenPair)
		return
	}
}

func (h *handler) logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
