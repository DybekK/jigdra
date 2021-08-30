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
	workspaceFacadeHandler := InitializeWorkspaceFacadeHandler(postgresDatabase)

	//initialize middleware
	r.Use(authMiddleware.TokenAuthMiddleware())

	//initialize routing
	r.Handle("GET", "/v1/workspace_user/:id", workspaceFacadeHandler.GetUser)
	r.Handle("POST", "/v1/workspace_facade", workspaceFacadeHandler.CreateUserAndWorkspace)

	//catch errors
	log.Fatal(http.ListenAndServe(":8080", r))
}
