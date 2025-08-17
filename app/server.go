package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(lc fx.Lifecycle) *Server {
	engine := gin.Default()
	s := &Server{engine: engine}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Start server on :8080")
				if err := engine.Run(":8080");err != nil {
					log.Fatal(err)
				}

			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping server...")
			return nil
		},
	})

	return s
}

func (s *Server)Gin() *gin.Engine{
	return s.engine
}