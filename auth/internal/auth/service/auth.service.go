package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	userpb "github.com/merdernoty/anime-proto/user"
	"github.com/merdernoty/microservices-planner/auth/internal/auth/domain"
	"github.com/merdernoty/microservices-planner/auth/internal/config"
	"github.com/merdernoty/microservices-planner/auth/pkg"
)

var jwtKey = config.LoadConfig().JWTSecret

type authService struct {
	userClient userpb.UserServiceClient
}

func NewAuthService(userClient userpb.UserServiceClient) domain.AuthService {
	return &authService{userClient: userClient}
}

func (s *authService) Register(username, email, password string) (*domain.Token, int64, error) {
	ctx := context.Background()

	hashedPassword, err := pkg.HashPassword(password)
	if err != nil {
		return nil, 0, err
	}

	resp, err := s.userClient.CreateUser(ctx, &userpb.CreateUserRequest{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, 0, err
	}

	token, err := pkg.GenerateToken(resp.User.Id)
	if err != nil {
		return nil, 0, err
	}

	return token, resp.User.Id, nil
}

func (s *authService) Login(email, password string) (*domain.Token, int64, error) {
	ctx := context.Background()

	resp, err := s.userClient.GetUserByEmail(ctx, &userpb.GetUserByEmailRequest{Email: email})
	if err != nil || resp.User == nil {
		fmt.Printf("Error fetching user by email: %v\n", err)
		return nil, 0, errors.New("user not found")
	}

	if !pkg.CheckPasswordHash(password, resp.User.Password) {
		fmt.Printf("Error fetching user by email: %v\n", err)
		return nil, 0, errors.New("invalid password")
	}

	token, err := pkg.GenerateToken(resp.User.Id)
	if err != nil {
		return nil, 0, err
	}

	return token, resp.User.Id, nil
}

func (s *authService) VerifyToken(tokenStr string) (bool, int64, time.Time, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil || !token.Valid {
		return false, 0, time.Time{}, errors.New("invalid token")
	}
	uid, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return false, 0, time.Time{}, err
	}
	var exp time.Time
	if claims.ExpiresAt != nil {
		exp = claims.ExpiresAt.Time
	}

	return true, uid, exp, nil
}
