package app

import (
    "go-users/internal/user/controller"
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