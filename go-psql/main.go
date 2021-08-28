package main

import (
	"go-psql/model"
	"go-psql/model/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	workspaceUserService model.WorkspaceUserService
)

func main() {
	_ = godotenv.Load("../.env")

	InitStuf()

	auth := &AuthHandler{}
	r := gin.Default()
	r.Use(auth.TokenAuthMiddleware())
	r.Handle("GET", "/v1/workuser/:id", getUser)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func InitStuf() {
	conn, err := repository.Initialize()
	if err != nil {
		panic(err)
	}
	workspaceUserRepo := repository.NewWorkspaceUserRepo(conn)
	workspaceUserService = model.NewWorkspaceUserService(workspaceUserRepo)

}
