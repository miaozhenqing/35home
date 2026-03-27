package service

import (
	"errors"
	"time"

	"35home/config"
	"35home/dto"
	"35home/models"
	"35home/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
	config   *config.Config
}

func NewUserService(userRepo repository.UserRepository, cfg *config.Config) UserService {
	return &userService{
		userRepo: userRepo,
		config:   cfg,
	}
}

func (s *userService) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// 验证密码是否一致
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("passwords do not match")
	}

	// 检查用户名是否已存在
	if s.userRepo.ExistsByUsername(req.Username) {
		return nil, errors.New("username already registered")
	}

	// 检查邮箱是否已存在
	if s.userRepo.ExistsByEmail(req.Email) {
		return nil, errors.New("email already registered")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// 创建用户
	user := &models.User{
		Username:    req.Username,
		Email:       req.Email,
		Password:    string(hashedPassword),
		Gender:      req.Gender,
		Age:         req.Age,
		Occupation:  req.Occupation,
		Hobbies:     req.Hobbies,
		LastLoginAt: time.Now().UnixMilli(),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	// 生成JWT令牌
	token, err := s.generateJWT(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// 构建响应
	response := &dto.RegisterResponse{
		User: dto.UserResponse{
			ID:         user.ID,
			Username:   user.Username,
			Email:      user.Email,
			Gender:     user.Gender,
			Age:        user.Age,
			Occupation: user.Occupation,
			Hobbies:    user.Hobbies,
			CreatedAt:  user.CreatedAt,
		},
		Token: token,
	}

	return response, nil
}

const (
	jwtExpiration = 24 * time.Hour
)

type Claims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func (s *userService) generateJWT(user *models.User) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "35home",
			Subject:   user.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *userService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// 查找用户（通过用户名）
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// 更新最后登录时间
	loginTime := time.Now().UnixMilli()
	if err := s.userRepo.UpdateLastLogin(user.ID, loginTime); err != nil {
		// 记录日志但不影响登录
		// log.Printf("Failed to update last login time: %v", err)
	}

	// 生成JWT令牌
	token, err := s.generateJWT(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// 构建响应
	response := &dto.LoginResponse{
		User: dto.UserResponse{
			ID:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			Gender:      user.Gender,
			Age:         user.Age,
			Occupation:  user.Occupation,
			Hobbies:     user.Hobbies,
			CreatedAt:   user.CreatedAt,
			LastLoginAt: loginTime,
		},
		Token: token,
	}

	return response, nil
}
