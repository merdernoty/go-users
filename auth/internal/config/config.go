package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	 GRPCPort string
	 HTTPPort string
	 UserServiceAddr string
	 JWTSecret string
}

func LoadConfig() *Config{
	_ = godotenv.Load()

	return &Config{
		GRPCPort:   getEnv("GRPC_PORT", "500052"),
		HTTPPort:   getEnv("HTTP_PORT", "8081"),
		UserServiceAddr: getEnv("USER_SERVICE_ADDR", "localhost:50051"),
		JWTSecret: getEnv("JWT_SECRET", "mysecret"),
	}
}

func getEnv(key,defaultValue string) string  {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}