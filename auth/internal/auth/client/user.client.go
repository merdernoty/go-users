package client

import (
	"context"
	"log"

	userpb "github.com/merdernoty/anime-proto/user"
	"github.com/merdernoty/microservices-planner/auth/internal/config"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func NewUserClient(lc fx.Lifecycle, cfg *config.Config) userpb.UserServiceClient {
	conn, err := grpc.Dial(cfg.UserServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to users service: %v", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return userpb.NewUserServiceClient(conn)
}
