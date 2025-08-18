package main

import (
	"github.com/merdernoty/microservices-planner/auth/app"
	"github.com/merdernoty/microservices-planner/auth/internal/auth"
	"github.com/merdernoty/microservices-planner/auth/internal/auth/client"
	"github.com/merdernoty/microservices-planner/auth/internal/config"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		app.Module,
		auth.Module,
		fx.Provide(client.NewUserClient),
	).Run()
}
