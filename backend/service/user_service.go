package service

import (
	"errors"
	"time"

	"35home/common"
	"35home/dto"
	"35home/models"
	"35home/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// 验证密码是否一致
	if req.Password != req.ConfirmPassword {
		return nil, errors.New(common.ErrorPasswordsDoNotMatch)
	}

	// 检查邮箱是否已存在
	if s.userRepo.ExistsByEmail(req.Email) {
		return nil, errors.New(common.ErrorEmailAlreadyRegistered)
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(common.ErrorFailedToHashPassword)
	}

	// 创建用户
	user := &models.User{
		Username:   req.Username,
		Email:      req.Email,
		Password:   string(hashedPassword),
		Gender:     req.Gender,
		Age:        req.Age,
		Occupation: req.Occupation,
		Hobbies:    req.Hobbies,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New(common.ErrorFailedToCreateUser)
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
	}

	return response, nil
}

func (s *userService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// 查找用户
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New(common.ErrorInvalidEmailOrPassword)
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New(common.ErrorInvalidEmailOrPassword)
	}

	// 更新最后登录时间
	loginTime := time.Now()
	if err := s.userRepo.UpdateLastLogin(user.ID, loginTime); err != nil {
		// 记录日志但不影响登录
		// log.Printf("Failed to update last login time: %v", err)
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
		Token: "dummy-token", // 实际项目中应该生成JWT token
	}

	return response, nil
}
