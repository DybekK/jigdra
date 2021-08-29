package main

import (
	"go-psql/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load("../.env")
	r := gin.Default()

	postgresDatabase := database.InitPostgresDatabase()
	handler := InitializeWorkspaceUserHandler(postgresDatabase)

	r.Handle("GET", "/v1/workuser/:id", handler.GetUser)
	log.Fatal(http.ListenAndServe(":8080", r))
}
