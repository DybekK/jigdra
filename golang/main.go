package main

import (
	"fmt"
	"golang/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("lets go")
	_ = godotenv.Load("../.env")
	client := sql.InitMongoDatabase()
	h := InitializeHandler(client)
	auth := InitializeAuthMiddleware()
	r := gin.Default()
	// returns 405 instead of 404 if you call a wrong method on an endpoint
	r.HandleMethodNotAllowed = true
	r.Use(gin.Logger())
	r.Use(CORSMiddleware())
	v1 := r.Group("/v1")
	{
		v1.POST("/register", h.AddUser)
		v1.GET("/login", h.Redirect)
		v1.POST("/login", h.Login)
		v1.GET("/user/:id", h.GetUserById)
		//These endpoints require Authorization header with valid Bearer token
		v1.POST("/logout", auth.TokenAuthMiddleware(), h.Logout)
		v1.POST("/refresh", auth.TokenAuthMiddleware(), h.Refresh)
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
