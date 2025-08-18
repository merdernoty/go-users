package database

import (
	"github.com/merdernoty/microservices-planner/users/internal/user/domain"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&domain.User{}); if err != nil {
		log.Fatalf("cannot migrate models", err)
	}
	log.Printf("Database migrated successfully")
}