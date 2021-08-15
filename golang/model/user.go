package model

import (
	"context"
	"errors"
	"fmt"
	"golang/model/dto"
	"golang/model/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	ValidateEmail(string) error
	CreateUser(*dto.User, context.Context) (string, error)
	GetUser(*dto.LoginUser, context.Context) (*dto.User, error)
	GetUserById(string, context.Context) (*dto.GetUserStruct, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) ValidateEmail(email string) error {
	return u.repo.ValidateEmail(email)

}

func (u *userService) GetUser(login *dto.LoginUser, ctx context.Context) (*dto.User, error) {
	user, err := u.repo.GetUser(login, ctx)
	if err != nil {
		return nil, err
	}
	if verifyPassword(login.Password, user.Password) {
		return user, nil
	} else {
		return nil, errors.New("invalid password")
	}
}

func (u *userService) CreateUser(req_user *dto.User, ctx context.Context) (string, error) {
	return u.repo.CreateUser(req_user, ctx)
}

func (u *userService) GetUserById(id string, ctx context.Context) (*dto.GetUserStruct, error) {
	return u.repo.GetUserById(id, ctx)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	fmt.Println(hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func validateEmail(email string) error {
	match := dto.Email_regex.MatchString(email)
	if match {
		return nil
	} else {
		return errors.New("invalid email")
	}
}
