//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jackc/pgx/v4"
	"go-psql/middleware"
	"go-psql/workspace"
)

func InitializeAuthMiddleware(postgresDatabase *pgx.Conn) middleware.AuthMiddleware {
	wire.Build(
		middleware.NewAuthMiddleware,
		workspace.NewWorkspaceUserService,
		workspace.NewWorkspaceUserRepo,
	)
	return middleware.AuthMiddleware{}
}

func InitializeWorkspaceUserHandler(postgresDatabase *pgx.Conn) workspace.WorkspaceUserHandler {
	wire.Build(
		workspace.NewWorkspaceUserHandler,
		workspace.NewWorkspaceUserService,
		workspace.NewWorkspaceUserRepo,
	)
	return workspace.WorkspaceUserHandler{}
}
