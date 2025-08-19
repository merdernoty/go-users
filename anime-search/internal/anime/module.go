package anime

import (
	"github.com/merdernoty/microservices-planner/anime-search/internal/anime/client"
	"github.com/merdernoty/microservices-planner/anime-search/internal/anime/controller"
	"github.com/merdernoty/microservices-planner/anime-search/internal/anime/domain"
	"github.com/merdernoty/microservices-planner/anime-search/internal/anime/service"
	"github.com/merdernoty/microservices-planner/anime-search/internal/anime/transport"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	transport.NewAnimeGRPC,
	fx.Annotate(
		service.NewAnimeService,
		fx.As(new(domain.AnimeService)),
	),
	controller.NewAnimeController,
	client.NewJikanClient,
)
