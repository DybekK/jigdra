//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-psql/middleware"
	"go-psql/workspace"
)

func InitializeAuthMiddleware(postgresDatabase *pgxpool.Pool) middleware.AuthMiddleware {
	wire.Build(
		middleware.NewAuthMiddleware,
		workspace.NewWorkspaceUserService,
		workspace.NewWorkspaceUserRepo,
	)
	return middleware.AuthMiddleware{}
}

func InitializeWorkspaceUserHandler(postgresDatabase *pgxpool.Pool) handler.WorkspaceUserHandler {
	wire.Build(
		workspace.NewWorkspaceUserHandler,
		workspace.NewWorkspaceUserService,
		workspace.NewWorkspaceUserRepo,
	)
	return workspace.WorkspaceUserHandler{}
}
