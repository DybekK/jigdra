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
