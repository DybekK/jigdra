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
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/v1", h.getUwa)
	r.POST("/v1/register", h.addUser)
	r.POST("/v1/login", h.login)
	//needs to be changed to "mongodb://mongodb:27017" if you want to run it in docker
	uri := "mongodb://localhost:27017"
	_, err := model.Interface.Initialize(uri)
	if err != nil {
		panic(err)
	}
	log.Fatal(r.Run(":4201"))
}
