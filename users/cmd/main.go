package main

import (
    "github.com/merdernoty/microservices-planner/users/app"
    "github.com/merdernoty/microservices-planner/users/internal/config"
    "github.com/merdernoty/microservices-planner/users/internal/database"
    "github.com/merdernoty/microservices-planner/users/internal/user"
    "go.uber.org/fx"
)

func main() {
    fx.New(
        config.Module,
        database.Module,
        app.Module,
        user.Module,
        fx.Invoke(database.Migrate),
    ).Run()
}
