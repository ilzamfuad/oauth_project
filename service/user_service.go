package service

import (
	"html"
	"nexmedis_project/model"
	"nexmedis_project/repository"
	"nexmedis_project/utils"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(model.User) error
	LoginCheck(username, password string) (string, error)
	GetUserByID(userID int) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) CreateUser(user model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return us.userRepo.CreateUser(user)
}

func (us *userService) LoginCheck(username, password string) (string, error) {
	var err error

	user, err := us.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = verifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *userService) GetUserByID(userID int) (*model.User, error) {
	return us.userRepo.GetUserByID(userID)
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
