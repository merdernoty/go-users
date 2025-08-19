package app

import (
	"github.com/merdernoty/microservices-planner/anime-search/internal/anime/controller"
	"go.uber.org/fx"
)

func RegisterRoutes(
	s *Server,
	animeSearchCtrl *controller.AnimeController,
) {
	api := s.Gin().Group("/api")
	animeSearchCtrl.RegisterRoutes(api)
}

var RoutesModule = fx.Invoke(RegisterRoutes)
