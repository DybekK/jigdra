//+build wireinject

package main

import (
	"golang/handler"
	"golang/middleware"
	"golang/redirect"
	"golang/user"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeHandler(client *mongo.Client) handler.Handler {
	wire.Build(
		user.NewUserRepository,
		user.NewUserService,
		redirect.NewRedirectRepository,
		redirect.NewRedirectService,
		handler.NewHandler,
	)

	return handler.Handler{}
}

func InitializeAuthMiddleware() middleware.AuthMiddleware {
	wire.Build(
		middleware.NewAuthMiddleware,
	)

	return middleware.AuthMiddleware{}
}
