package user

import (
	"go-users/internal/user/repository"
	"go-users/internal/user/service"
	"go-users/internal/user/controller"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	),
)
