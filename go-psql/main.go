package main

import (
	"go-psql/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load("../.env")
	r := gin.Default()

	//initialize postgres connection
	postgresDatabase := sql.InitPostgresDatabase()

	//initialize services
	authMiddleware := InitializeAuthMiddleware(postgresDatabase)
	workspaceUserHandler := InitializeWorkspaceUserHandler(postgresDatabase)

	//initialize middleware
	r.Use(authMiddleware.TokenAuthMiddleware())

	//initialize routing
	r.Handle("GET", "/v1/workuser/:id", workspaceUserHandler.GetUser)

	//catch errors
	log.Fatal(http.ListenAndServe(":8080", r))
}
