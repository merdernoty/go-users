package main

import (
    "go-users/app"
    "go-users/internal/config"
    "go-users/internal/database"
    "go-users/internal/user"
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
