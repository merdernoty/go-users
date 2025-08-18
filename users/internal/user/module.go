package user

import (
	"github.com/merdernoty/microservices-planner/users/internal/user/repository"
	"github.com/merdernoty/microservices-planner/users/internal/user/service"
	"github.com/merdernoty/microservices-planner/users/internal/user/controller"
	"github.com/merdernoty/microservices-planner/users/internal/user/transport"
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
