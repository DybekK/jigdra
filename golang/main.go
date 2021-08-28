package main

import (
	"fmt"
	"golang/model"
	"golang/model/repository"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	userRepo        repository.UserRepository
	userService     model.UserService
	redirectRepo    repository.RedirectRepository
	redirectService model.RedirectService
)

func main() {
	fmt.Println("lets go")
	_ = godotenv.Load("../.env")
	userRepo = repository.NewUserRepository()
	userService = model.NewUserService(userRepo)
	redirectRepo = repository.NewRedirectRepository()
	redirectService = model.NewRedirectService(redirectRepo)
	h := &handler{}
	auth := &model.AuthHandler{}
	r := gin.Default()
	// returns 405 instead of 404 if you call a wrong method on an endpoint
	r.HandleMethodNotAllowed = true
	r.Use(gin.Logger())
	r.Use(CORSMiddleware())
	v1 := r.Group("/v1")
	{
		v1.GET("/", h.getUwa)
		v1.POST("/register", h.addUser)
		v1.GET("/login", h.login)
		v1.POST("/login", h.login)
		v1.GET("/user/:id", h.getUserById)
		//These endpoints require Authorization header with valid Bearer token
		v1.POST("/logout", auth.TokenAuthMiddleware(), h.logout)
		v1.POST("/refresh", auth.TokenAuthMiddleware(), h.refresh)
	}
	log.Fatal(r.Run(":4201"))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
