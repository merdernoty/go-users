package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	 DbHost string
	 DBPort string
	 DBUser string
	 DBPassword string
	 DBName string
	 DBSSLMode string
}

func LoadConfig() *Config{
	_ = godotenv.Load()

	return &Config{
		DbHost: 	getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "users_db"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key,defaultValue string) string  {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}