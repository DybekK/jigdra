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
	taskHandler := InitializeTaskHandler(postgresDatabase, authMiddleware)

	//initialize middleware
	r.Use(authMiddleware.TokenAuthMiddleware())

	//initialize routing
	v1 := r.Group("/v1")
	{
		v1.Handle("GET", "/workspace_user/:id", workspaceFacadeHandler.GetUser)
		v1.Handle("POST", "/workspace_facade", workspaceFacadeHandler.CreateUserAndWorkspace)
		v1.Handle("POST", "/task", taskHandler.CreateTask)
		v1.Handle("GET", "/user_tasks/:uuid", taskHandler.GetUserTasks)
		v1.Handle("GET", "/task/:id", taskHandler.GetTask)
	}

	//catch errors
	log.Fatal(http.ListenAndServe(":8080", r))
}
