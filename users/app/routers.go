package app

import (
    "github.com/merdernoty/microservices-planner/users/internal/user/controller"
    "go.uber.org/fx"
)

func RegisterRoutes(
	s *Server,
	userCtrl *controller.UserController,
	) {
	api := s.Gin().Group("/api")
	userCtrl.RegisterRoutes(api)
}

var RoutesModule = fx.Invoke(RegisterRoutes)