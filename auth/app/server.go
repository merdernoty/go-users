package app

import (
	"context"
	"github.com/merdernoty/microservices-planner/auth/internal/config"
	"github.com/merdernoty/microservices-planner/auth/internal/auth/domain"
	"github.com/merdernoty/microservices-planner/auth/internal/auth/transport"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"google.golang.org/grpc"

	authpb "github.com/merdernoty/anime-proto/auth"
)

type Server struct {
	engine *gin.Engine
	grpc   *grpc.Server
}

func NewServer(lc fx.Lifecycle, us domain.AuthService) *Server {
	engine := gin.Default()

	grpcServer := grpc.NewServer()
	authpb.RegisterAuthServiceServer(grpcServer, transport.NewAuthGRPCServer(us))

	s := &Server{engine: engine, grpc: grpcServer}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Start REST server on :", config.LoadConfig().HTTPPort)
				if err := engine.Run(":" + config.LoadConfig().HTTPPort); err != nil {
					log.Fatal(err)
				}

			}()

			go func() {
				lis, err := net.Listen("tcp", ":" + config.LoadConfig().GRPCPort)
				if err != nil {
					log.Fatal(err)
				}
				log.Println("Start gRPC server on :", config.LoadConfig().GRPCPort)
				if err := grpcServer.Serve(lis); err != nil {
					log.Fatal(err)
				}

			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping server...")
			s.grpc.GracefulStop()
			return nil
		},
	})

	return s
}

func (s *Server) Gin() *gin.Engine {
	return s.engine
}
