package app

import (
	"context"
	"go-users/internal/config"
	"go-users/internal/user/domain"
	"go-users/internal/user/transport"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"google.golang.org/grpc"

	userpb "github.com/merdernoty/anime-proto/user"
)

type Server struct {
	engine *gin.Engine
	grpc   *grpc.Server
}

func NewServer(lc fx.Lifecycle, us domain.UserService) *Server {
	engine := gin.Default()

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, transport.NewUserGRPCServer(us))

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
