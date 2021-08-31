package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang/redirect"
	"golang/user"
	"net/http"
	"net/url"
)

type Handler struct {
	userService     user.UserService
	redirectService redirect.RedirectService
}

//factory

func NewHandler(userService user.UserService, redirectService redirect.RedirectService) Handler {
	return Handler{userService: userService, redirectService: redirectService}
}

var validate = validator.New()

//methods

func (h *Handler) AddUser(c *gin.Context) {
	var reqUser user.User
	if err := c.BindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user": "invalid request body"})
		return
	}

	validationError := validate.Struct(reqUser)
	if validationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user": "invalid request body"})
		return
	}
	res, err := h.userService.CreateUser(&reqUser, c)
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
	hex, err := h.redirectService.SecureRedirect(c, res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	q := url.Values{}
	q.Add("redirect", hex)
	location := url.URL{Path: "/v1/login", RawQuery: q.Encode()}

	c.Redirect(http.StatusFound, location.RequestURI())

}

func (h *Handler) GetUserById(c *gin.Context) {

	id := c.Param("id")
	_user, err := h.userService.GetUserById(id, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"user": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, _user)

}

func (h *Handler) Refresh(c *gin.Context) {
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
		newTokenPair, err := user.GenerateTokenPair(claims["identityKey"].(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"token": "signing error"})
			return
		}

		c.JSON(http.StatusOK, newTokenPair)
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"token": err.Error()})
}

func (h *Handler) Redirect(c *gin.Context) {
	q := c.Query("redirect")
	if q == "" {
		c.String(405, "405 Method not allowed")
		return
	}
	id, exists := h.redirectService.VerifyRedirect(c, q)
	if exists != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"redirect": "redirect doesn't exist"})
		return
	}
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"_id": "invalid hex"})
		return
	}
	res, err := h.userService.GetUserById(objId.Hex(), c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user": "user not found"})
	}

	newTokenPair, err := user.GenerateTokenPair(res.Id.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"token": "signing error"})
		return
	}

	c.JSON(http.StatusOK, newTokenPair)
}

func (h *Handler) Login(c *gin.Context) {
	var reqLogin user.LoginUser
	if err := c.BindJSON(&reqLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user": "invalid request body"})
		return
	}
	validationError := validate.Struct(reqLogin)
	if validationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user": "validation error"})
		return
	}
	_user, userErr := h.userService.GetUser(&reqLogin, c)
	if userErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"login": "invalid email or password"})
		return
	}
	newTokenPair, err := user.GenerateTokenPair(_user.Id.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"token": "signing error"})
		return
	}

	c.JSON(http.StatusOK, newTokenPair)
}

func (h *Handler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
