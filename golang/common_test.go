package main

import (
	"golang/model"
	"golang/model/repository"

	"github.com/gin-gonic/gin"
)

var (
	mockUserRepo     *repository.MockUserRepo
	mockRedirectRepo *repository.MockRedirectRepo
)

func getRouter() *gin.Engine {
	mockUserRepo = new(repository.MockUserRepo)
	userService = model.NewUserService(mockUserRepo)
	mockRedirectRepo = new(repository.MockRedirectRepo)
	redirectService = model.NewRedirectService(mockRedirectRepo)
	r := gin.Default()
	return r
}
