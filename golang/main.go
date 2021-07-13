package main

import (
	"fmt"
	"golang/model"
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("lets go")
	h := &handler{}
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(CORSMiddleware())
	middleware, jwt_err := jwt.New(model.Interface.GetMiddleWare())
	if jwt_err != nil {
		panic(jwt_err)
	}

	errInit := middleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("middleware.MiddlewareInit() error: " + errInit.Error())
	}

	r.GET("/v1", h.getUwa)
	r.POST("/v1/register", h.addUser)
	r.GET("/v1/login", middleware.LoginHandler)
	r.POST("/v1/login", middleware.LoginHandler)
	r.POST("/v1/logout", middleware.LogoutHandler)
	r.GET("/v1/refresh", middleware.RefreshHandler)
	r.GET("/v1/user/:id", h.getUserById)
	//Connect to database
	_, err := model.Interface.Initialize()
	if err != nil {
		panic(err)
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
