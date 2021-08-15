package main

import (
	"fmt"
	"golang/model"
	"golang/model/dto"
	"net/http"
	"net/url"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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
	var req_user dto.User
	if err := c.BindJSON(&req_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user": "invalid request body"})
		return
	}

	validationError := validate.Struct(req_user)
	if validationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user": "invalid request body"})
		return
	}

	res, err := userService.CreateUser(&req_user, c)
	hex, err := redirectRepo.SecureRedirect(c, res)
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
	q.Add("redirect", hex)
	location := url.URL{Path: "/v1/login", RawQuery: q.Encode()}

	c.Redirect(http.StatusFound, location.RequestURI())

}

func (h *handler) getUserById(c *gin.Context) {

	id := c.Param("id")
	user, err := userService.GetUserById(id, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"user": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, user)

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
		if q == "" {
			c.String(405, "405 Method not allowed")
			return
		}
		id, exists := redirectService.VerifyRedirect(c, q)
		if exists != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"redirect": "redirect doesn't exist"})
			return
		}
		objid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"_id": "invalid hex"})
			return
		}
		res, err := userService.GetUserById(objid.Hex(), c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"user": "user not found"})
		}

		newTokenPair, err := auth.GenerateTokenPair(res.Id.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"token": "signing error"})
			return
		}

		c.JSON(http.StatusOK, newTokenPair)

	} else if c.Request.Method == "POST" {
		var req_login dto.LoginUser
		if err := c.BindJSON(&req_login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"user": "invalid request body"})
			return
		}
		validation_error := validate.Struct(req_login)
		if validation_error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"user": "validation error"})
			return
		}
		user, user_err := userService.GetUser(&req_login, c)
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
