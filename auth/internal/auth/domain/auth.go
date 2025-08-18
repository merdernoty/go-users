package domain

import "time"

type AuthRequest struct {
	Username string
	Email    string
	Password string
}

type AuthResponse struct {
	UserID int64
	Token  string
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type VerifyRequest struct {
	Token string `json:"token" binding:"required"`
}

type Token struct {
	Value     string
	ExpiresAt time.Time
}

type AuthService interface {
	Register(username, email, password string) (*Token, int64, error)
	Login(email, password string) (*Token, int64, error)
	VerifyToken(token string) (bool, int64, time.Time, error)
}
