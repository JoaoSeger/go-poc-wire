//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/meseger/portifolio/wire/internal/user/application"
	"github.com/meseger/portifolio/wire/internal/user/domain"
	"github.com/meseger/portifolio/wire/internal/user/infrastructure"
)

func InitializeUserService() *application.UserService {
	wire.Build(
		infrastructure.NewInMemoryUserRepository,
		wire.Bind(new(domain.UserRepository), new(*infrastructure.InMemoryUserRepository)),
		application.NewUserService,
	)
	return nil
}
