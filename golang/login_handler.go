package main

import (
	"context"
	"golang/model"
	"net/http"
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

	result, err := model.Interface.CreateUser(&req_user, ctx)
	if err != nil {
		if err.Error() == "email already in use" {
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
	c.JSON(http.StatusOK, result)

}
