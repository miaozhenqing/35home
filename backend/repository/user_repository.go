package repository

import (
	"35home/models"
	"errors"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByID(id uint64) (*models.User, error)
	UpdateLastLogin(id uint64, loginTime int64) error
	ExistsByEmail(email string) bool
	ExistsByUsername(username string) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id uint64) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateLastLogin(id uint64, loginTime int64) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("last_login_at", loginTime).Error
}

func (r *userRepository) ExistsByEmail(email string) bool {
	var count int
	r.db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (r *userRepository) ExistsByUsername(username string) bool {
	var count int
	r.db.Model(&models.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func (r *userRepository) CreateIfNotExists(user *models.User) error {
	if r.ExistsByEmail(user.Email) {
		return errors.New("email already exists")
	}
	return r.Create(user)
}
