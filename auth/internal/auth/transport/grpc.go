package transport

import (
	"context"

	authpb "github.com/merdernoty/anime-proto/auth"
	"github.com/merdernoty/microservices-planner/auth/internal/auth/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthGRPCServer struct {
	authpb.UnimplementedAuthServiceServer
	service domain.AuthService
}

func NewAuthGRPCServer(s domain.AuthService) *AuthGRPCServer {
	return &AuthGRPCServer{service: s}
}

func (s *AuthGRPCServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	token, userID, err := s.service.Register(req.Username, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &authpb.RegisterResponse{
		Token: &authpb.Token{
			Value:     token.Value,
			ExpiresAt: timestamppb.New(token.ExpiresAt),
		},
		UserId: userID,
	}, nil
}

func (s *AuthGRPCServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	token, userID, err := s.service.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &authpb.LoginResponse{
		Token: &authpb.Token{
			Value:     token.Value,
			ExpiresAt: timestamppb.New(token.ExpiresAt),
		},
		UserId: userID,
	}, nil
}

func (s *AuthGRPCServer) VerifyToken(ctx context.Context, req *authpb.VerifyTokenRequest) (*authpb.VerifyTokenResponse, error) {
	valid, userID, exp, err := s.service.VerifyToken(req.Token)
	if err != nil {
		return nil, err
	}

	var expiresAt *timestamppb.Timestamp
	if !exp.IsZero() {
		expiresAt = timestamppb.New(exp)
	}

	return &authpb.VerifyTokenResponse{
		Valid:     valid,
		UserId:    userID,
		ExpiresAt: expiresAt,
	}, nil
}
