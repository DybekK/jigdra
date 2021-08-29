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

	//initialize postgres connection
	postgresDatabase := database.InitPostgresDatabase()

	//initialize services
	authMiddleware := InitializeAuthMiddleware(postgresDatabase)
	handler := InitializeWorkspaceUserHandler(postgresDatabase)

	//initialize middleware
	r.Use(authMiddleware.TokenAuthMiddleware())

	//initialize routing
	r.Handle("GET", "/v1/workuser/:id", handler.GetUser)

	//catch errors
	log.Fatal(http.ListenAndServe(":8080", r))
}
