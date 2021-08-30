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
		workspace.NewWorkspaceUserRepository,
	)
	return middleware.AuthMiddleware{}
}

func InitializeWorkspaceFacadeHandler(postgresDatabase *pgxpool.Pool) workspace.WorkspaceFacadeHandler {
	wire.Build(
		workspace.NewWorkspaceFacadeHandler,
		workspace.NewWorkspaceFacade,
		workspace.NewWorkspaceUserService,
		workspace.NewWorkspaceService,
		workspace.NewWorkspaceRepository,
		workspace.NewWorkspaceUserRepository,
	)
	return workspace.WorkspaceFacadeHandler{}
}
