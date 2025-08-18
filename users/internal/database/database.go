package database

import (
	"fmt"
	"github.com/merdernoty/microservices-planner/users/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDatabase(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DbHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connect to database", err)
	}
	return db
}
