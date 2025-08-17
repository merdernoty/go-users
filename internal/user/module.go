package user

import (
	"go-users/internal/user/repository"
	"go-users/internal/user/service"
	"go-users/internal/user/controller"
	"go-users/internal/user/transport"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		transport.NewUserGRPCServer,
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	),
)
