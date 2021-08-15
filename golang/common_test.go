package main

import (
	"golang/model"
	"golang/model/repository"

	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	userRepo = repository.NewUserRepository()
	userService = model.NewUserService(userRepo)
	redirectRepo = repository.NewRedirectRepository()
	redirectService = model.NewRedirectService(redirectRepo)
	r := gin.Default()
	return r
}
