package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load("../.env")
	r := gin.Default()

	handler := InitializeWorkspaceUserHandler()

	r.Handle("GET", "/v1/workuser/:id", handler.GetUser)
	log.Fatal(http.ListenAndServe(":8080", r))
}
