package app

import (
	"github.com/merdernoty/microservices-planner/auth/internal/auth/controller"
	"go.uber.org/fx"
)

func RegisterRoutes(
	s *Server,
	authCtrl *controller.AuthController,
) {
	api := s.Gin().Group("/api")
	authCtrl.RegisterRoutes(api)
}

var RoutesModule = fx.Invoke(RegisterRoutes)
