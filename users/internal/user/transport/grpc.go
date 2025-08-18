package transport

import (
	"context"
	userpb "github.com/merdernoty/anime-proto/user"
	"github.com/merdernoty/microservices-planner/users/internal/user/domain"
	"github.com/merdernoty/microservices-planner/users/internal/user/helpers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserGRPCServer struct {
	userpb.UnimplementedUserServiceServer
	us domain.UserService
}

func NewUserGRPCServer(us domain.UserService) *UserGRPCServer {
	return &UserGRPCServer{us: us}
}

func (s *UserGRPCServer) GetUser(сtx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := s.us.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &userpb.GetUserResponse{
		User: helpers.ToUserProto(user),
	}, nil
}

func (s *UserGRPCServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := s.us.GetAllUsers()
	if err != nil {
		return nil, err
	}
	resp := &userpb.ListUsersResponse{}
	for _, u := range users {
		resp.Users = append(resp.Users, helpers.ToUserProto(&u))
	}
	return resp, nil
}

func (s *UserGRPCServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := s.us.CreateUser(user); err != nil {
		return nil, err
	}
	return &userpb.CreateUserResponse{
		User: helpers.ToUserProto(user),
	}, nil
}

func (s *UserGRPCServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	user := &domain.User{
		ID:       req.Id,
		Email:    req.Email,
		Username: req.Username,
		Bio:      req.Bio,
	}

	err := s.us.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: helpers.ToUserProto(user),
	}, nil
}

func (s *UserGRPCServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	err := s.us.DeleteUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{
		Success: true,
	}, nil
}

func (s *UserGRPCServer) GetUserByEmail(ctx context.Context, req *userpb.GetUserByEmailRequest) (*userpb.GetUserByEmailResponse, error) {
	user, err := s.us.GetUserByEmail(req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	if user == nil {
		return &userpb.GetUserByEmailResponse{User: nil}, nil
	}

	return &userpb.GetUserByEmailResponse{
		User: helpers.ToUserProto(user),
	}, nil
}
