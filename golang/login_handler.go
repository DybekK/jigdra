package main

import (
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
	if c.Request.Method == "POST" {
		var req_user model.User
		if err := c.BindJSON(&req_user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationError := validate.Struct(req_user)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
			return
		}

		res, err := model.Interface.CreateUser(&req_user, c)

		if err != nil {
			if err.Error() == "409" {
				c.JSON(http.StatusConflict, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
			return
		}
		q := url.Values{}
		q.Add("redirect", res)
		location := url.URL{Path: "/v1/login", RawQuery: q.Encode()}

		c.Redirect(http.StatusFound, location.RequestURI())
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported http method"})

}

func (h *handler) getUserById(c *gin.Context) {
	if c.Request.Method == "GET" {
		id := c.Param("id")
		user, err := model.Interface.GetUserById(id, c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported http method"})
}

func (h *handler) refresh(c *gin.Context) {
	if c.Request.Method == "POST" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported http method"})
	}
	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	tokenReq := tokenReqBody{}

	if err := c.BindJSON(&tokenReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	token, err := auth.VerifyToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		newTokenPair, err := auth.GenerateTokenPair(claims["identitykey"].(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newTokenPair)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func (h *handler) login(c *gin.Context) {
	if c.Request.Method == "GET" {
		q := c.Query("redirect")
		if q != "" {
			id, exists := model.Interface.VerifyRedirect(c, q)
			if exists != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "failed redirect"})
				return
			}
			objid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			res := model.Interface.GetCollection("users").FindOne(c, bson.M{"_id": objid})
			var user model.User
			decode_err := res.Decode(&user)
			if decode_err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": decode_err.Error()})
				return
			}

			newTokenPair, err := auth.GenerateTokenPair(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, newTokenPair)
			return
		}
	} else if c.Request.Method == "POST" {
		var req_login model.LoginUser
		if err := c.BindJSON(&req_login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validation_error := validate.Struct(req_login)
		if validation_error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validation_error.Error()})
			return
		}
		user, user_err := model.Interface.GetUser(&req_login, c)
		if user_err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": user_err.Error()})
			return
		}
		newTokenPair, err := auth.GenerateTokenPair(user.Id.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newTokenPair)
		return
	}
	c.JSON(http.StatusBadRequest, "unsupported http method")
}

func (h *handler) logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
