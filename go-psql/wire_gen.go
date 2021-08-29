// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/jackc/pgx/v4"
	"go-psql/database/repository"
	"go-psql/handler"
	"go-psql/middleware"
	"go-psql/service"
)

// Injectors from wire.go:

func InitializeAuthMiddleware(postgresDatabase *pgx.Conn) middleware.AuthMiddleware {
	workspaceUserRepository := repository.NewWorkspaceUserRepo(postgresDatabase)
	workspaceUserService := service.NewWorkspaceUserService(workspaceUserRepository)
	authMiddleware := middleware.NewAuthMiddleware(workspaceUserService)
	return authMiddleware
}

func InitializeWorkspaceUserHandler(postgresDatabase *pgx.Conn) handler.WorkspaceUserHandler {
	workspaceUserRepository := repository.NewWorkspaceUserRepo(postgresDatabase)
	workspaceUserService := service.NewWorkspaceUserService(workspaceUserRepository)
	workspaceUserHandler := handler.NewWorkspaceUserHandler(workspaceUserService)
	return workspaceUserHandler
}
