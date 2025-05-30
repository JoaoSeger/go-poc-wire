// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/meseger/portifolio/wire/internal/user/application"
	"github.com/meseger/portifolio/wire/internal/user/infrastructure"
)

// Injectors from wire.go:

func InitializeUserService() *application.UserService {
	inMemoryUserRepository := infrastructure.NewInMemoryUserRepository()
	userService := application.NewUserService(inMemoryUserRepository)
	return userService
}
