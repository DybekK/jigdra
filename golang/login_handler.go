package main

import (
	"context"
	"fmt"
	"golang/model"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type handler struct{}

func (h *handler) getUwa(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "success",
	})
}

var validate = validator.New()

func (h *handler) addUser(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

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
	defer cancel()

	res, err := model.Interface.CreateUser(&req_user, ctx)

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
	id := strings.Split(fmt.Sprintf("%v", res.InsertedID), "\"")[1]
	q := url.Values{}
	q.Add("redirect", strings.Split(id, ")")[0])
	location := url.URL{Path: "/v1/login", RawQuery: q.Encode()}

	c.Redirect(http.StatusFound, location.RequestURI())

}

func (h *handler) getUserById(c *gin.Context) {
	id := c.Param("id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	user, err := model.Interface.GetUserById(id, ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, user)

}
