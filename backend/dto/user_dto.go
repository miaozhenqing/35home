package dto

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
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID          uint64 `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Occupation  string `json:"occupation"`
	Hobbies     string `json:"hobbies"`
	CreatedAt   int64  `json:"created_at"`
	LastLoginAt int64  `json:"last_login_at"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type RegisterResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
