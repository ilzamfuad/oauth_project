package repository

import (
	"nexmedis_project/model"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(model.User) error
	GetUserByUsername(username string) (*model.User, error)
	GetUserByID(userID int) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(user model.User) error {
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()
	err := ur.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := ur.db.Model(model.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) GetUserByID(userID int) (*model.User, error) {
	user := &model.User{}
	err := ur.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
