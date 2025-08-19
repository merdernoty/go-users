package main

import (
	"github.com/merdernoty/microservices-planner/anime-search/app"
	"github.com/merdernoty/microservices-planner/anime-search/internal/anime"
	"github.com/merdernoty/microservices-planner/anime-search/internal/config"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		app.Module,
		anime.Module,
	).Run()
}
