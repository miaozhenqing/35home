package dto

import "time"

type RegisterRequest struct {
	Username        string `json:"username" binding:"required,min=2,max=50"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	Gender          string `json:"gender"`
	Age             int    `json:"age"`
	Occupation      string `json:"occupation"`
	Hobbies         string `json:"hobbies"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Gender      string    `json:"gender"`
	Age         int       `json:"age"`
	Occupation  string    `json:"occupation"`
	Hobbies     string    `json:"hobbies"`
	CreatedAt   time.Time `json:"created_at"`
	LastLoginAt time.Time `json:"last_login_at"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type RegisterResponse struct {
	User UserResponse `json:"user"`
}
