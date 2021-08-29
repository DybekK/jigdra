//+build wireinject

package main

import (
	"go-psql/database"
	"go-psql/database/repository"
	"go-psql/handler"
	"go-psql/service"

	"github.com/google/wire"
)

func InitializeWorkspaceUserHandler() handler.WorkspaceUserHandler {
	wire.Build(
		handler.NewWorkspaceUserHandler,
		service.NewWorkspaceUserService,
		repository.NewWorkspaceUserRepo,
		database.NewPostgresDatabase,
	)
	return handler.WorkspaceUserHandler{}
}
