package config

import "go.uber.org/fx"

var Module = fx.Provide(
	LoadConfig,
)
