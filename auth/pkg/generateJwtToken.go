package pkg

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/merdernoty/microservices-planner/auth/internal/auth/domain"
	"github.com/merdernoty/microservices-planner/auth/internal/config"
)

var jwtKey = config.LoadConfig().JWTSecret

func GenerateToken(userID int64) (*domain.Token, error) {
	expiration := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(userID, 10),
		ExpiresAt: jwt.NewNumericDate(expiration),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := t.SignedString([]byte(jwtKey))
	if err != nil {
		return nil, err
	}

	return &domain.Token{Value: tokenStr, ExpiresAt: expiration}, nil
}
