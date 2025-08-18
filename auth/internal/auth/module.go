package auth

import (
	"github.com/merdernoty/microservices-planner/auth/internal/auth/controller"
	"github.com/merdernoty/microservices-planner/auth/internal/auth/service"
	"github.com/merdernoty/microservices-planner/auth/internal/auth/transport"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		transport.NewAuthGRPCServer,
		service.NewAuthService,
		controller.NewAuthController,
	),
)
