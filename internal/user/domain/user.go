package domain

import "time"

type User struct {
	ID        int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string `gorm:"size:50;unique;not null" json:"username"`
	Email     string `gorm:"size:255;unique;not null" json:"email"`
	Password  string `gorm:"size:255;not null" json:"-"`
	AvatarURL string `gorm:"size:500" json:"avatar_url"`
	Bio       string `gorm:"size:500" json:"bio"`


	FavoriteGenres []string `gorm:"-" json:"favorite_genres,omitempty"` //TODO вынести в отдельную таблицу
//	Watchlist      []Anime  `gorm:"-" json:"watchlist,omitempty"`       //TODO связи многие-ко-многим
	FollowersCount int64    `gorm:"default:0" json:"followers_count"`
	FollowingCount int64    `gorm:"default:0" json:"following_count"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id int64) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int64) error
}

type UserService interface {
	GetAllUsers() ([]User, error)
	GetUser(id int64) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id int64) error
}
