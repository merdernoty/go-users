package helpers

import (
	userpb "github.com/merdernoty/anime-proto/user"
	"github.com/merdernoty/microservices-planner/users/internal/user/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserProto(user *domain.User) *userpb.User {
	return &userpb.User{
		Id: user.ID,
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
		Bio: user.Bio,
		AvatarUrl: user.AvatarURL,
		FavoriteGenres: user.FavoriteGenres,
		FollowersCount: user.FollowersCount,
		FollowingCount: user.FollowingCount,
		CreatedAt:      timestamppb.New(user.CreatedAt),
		UpdatedAt:      timestamppb.New(user.UpdatedAt),
	}
}
