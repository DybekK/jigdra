package main

import (
	"fmt"
	"golang/model"
	"log"
	"os"

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
	r.POST("/v1/login", middleware.LoginHandler)
	r.POST("/v1/logout", middleware.LogoutHandler)
	r.GET("/v1/refresh", middleware.RefreshHandler)
	r.GET("/v1/user/:id", h.getUserById)
	db_user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	db_passwd := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	uri := fmt.Sprintf("mongodb://%s:%s@mongodb:27017", db_user, db_passwd)
	_, err := model.Interface.Initialize(uri)
	if err != nil {
		panic(err)
	}
	log.Fatal(r.Run(":4201"))
}
