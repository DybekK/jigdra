//+build wireinject

package main

import (
	"go-psql/database/repository"
	"go-psql/handler"
	"go-psql/middleware"
	"go-psql/service"

	"github.com/google/wire"
	"github.com/jackc/pgx/v4/pgxpool"
)

func InitializeAuthMiddleware(postgresDatabase *pgxpool.Pool) middleware.AuthMiddleware {
	wire.Build(
		middleware.NewAuthMiddleware,
		service.NewWorkspaceUserService,
		repository.NewWorkspaceUserRepo,
	)
	return middleware.AuthMiddleware{}
}

func InitializeWorkspaceUserHandler(postgresDatabase *pgxpool.Pool) handler.WorkspaceUserHandler {
	wire.Build(
		handler.NewWorkspaceUserHandler,
		service.NewWorkspaceUserService,
		repository.NewWorkspaceUserRepo,
	)
	return handler.WorkspaceUserHandler{}
}
