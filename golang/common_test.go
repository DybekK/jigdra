package main

import (
	"golang/model"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func getMiddleWare() (*jwt.GinJWTMiddleware, error) {
	middleware, err := jwt.New(model.Interface.GetMiddleWare())
	if err != nil {
		return nil, err
	}
	return middleware, nil
}
