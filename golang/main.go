package main

import (
	"fmt"
	"golang/model"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("lets go")
	h := &handler{}
	auth := &model.AuthHandler{}
	r := gin.Default()
	// returns 405 instead of 404 if you call a wrong method on an endpoint
	r.HandleMethodNotAllowed = true
	r.Use(gin.Logger())
	r.Use(CORSMiddleware())
	r.GET("/v1", h.getUwa)
	r.POST("/v1/register", h.addUser)
	r.GET("/v1/login", h.login)
	r.POST("/v1/login", h.login)
	r.POST("/v1/refresh", h.refresh)
	r.GET("/v1/user/:id", h.getUserById)
	//These endpoints require Authorization header with valid Bearer token
	r.POST("/v1/logout", auth.TokenAuthMiddleware(), h.logout)
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
