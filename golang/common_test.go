package main

import (
	"golang/model"
	"golang/model/repository"

	"github.com/gin-gonic/gin"
)


func getRouter() *gin.Engine {
	userRepo = new(repository.MockUserRepo)
	userService = model.NewUserService(userRepo)
	redirectRepo = new(repository.MockRedirectRepo)
	redirectService = model.NewRedirectService(redirectRepo)
	r := gin.Default()
	return r
}
