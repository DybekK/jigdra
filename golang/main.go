package main

import (
	"golang/middleware"
	"golang/nosql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load("../.env")
	r := gin.Default()

	//initialize mongo connection
	mongoDatabase := nosql.InitMongoDatabase()

	//initialize services
	handler := InitializeHandler(mongoDatabase)
	authMiddleware := InitializeAuthMiddleware()

	//initialize middleware
	// returns 405 instead of 404 if you call a wrong method on an endpoint
	r.HandleMethodNotAllowed = true
	r.Use(gin.Logger())
	r.Use(middleware.CORSMiddleware())

	//initialize routing
	v1 := r.Group("/v1")
	{
		v1.POST("/register", handler.AddUser)
		v1.GET("/login", handler.Redirect)
		v1.POST("/login", handler.Login)
		v1.GET("/user/:id", handler.GetUserById)
		//These endpoints require Authorization header with valid Bearer token
		v1.POST("/logout", authMiddleware.TokenAuthMiddleware(), handler.Logout)
		v1.POST("/refresh", authMiddleware.TokenAuthMiddleware(), handler.Refresh)
	}
	log.Fatal(r.Run(":4201"))
}
