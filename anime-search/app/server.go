package app

import (
	"context"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"google.golang.org/grpc"

	animeSearchpb "github.com/merdernoty/anime-proto/anime"
	"github.com/merdernoty/microservices-planner/anime-search/internal/anime/domain"
	"github.com/merdernoty/microservices-planner/anime-search/internal/anime/transport"
	"github.com/merdernoty/microservices-planner/anime-search/internal/config"
)

type Server struct {
	engine *gin.Engine
	grpc   *grpc.Server
}

func NewServer(lc fx.Lifecycle, as domain.AnimeService) *Server {
	engine := gin.Default()

	grpcServer := grpc.NewServer()
	animeSearchpb.RegisterAnimeServiceServer(grpcServer, transport.NewAnimeGRPC(as))

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
				lis, err := net.Listen("tcp", ":"+config.LoadConfig().GRPCPort)
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
