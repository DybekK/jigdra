package model

import (
	"context"
	"errors"
	"fmt"
	"golang/model/dto"
	"golang/model/repository"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(*dto.User, context.Context) (string, error)
	GetUser(*dto.LoginUser, context.Context) (*dto.User, error)
	GetUserById(string, context.Context) (*dto.GetUserStruct, error)
	IsUsernameAvailable(string, context.Context) bool
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
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
	req_user.Id = primitive.NewObjectID()
	email_err := validateEmail(req_user.Email)
	if email_err != nil {
		return "", email_err
	}

	var hash, hash_error = hashPassword(req_user.Password)
	if hash_error != nil {
		return "", errors.New("hashing error")
	}

	if availible := u.repo.IsUsernameAvailable(req_user.Username, ctx); !availible {
		return "", errors.New("username taken")
	}
	req_user.Password = hash
	res, err := u.repo.CreateUser(req_user, ctx)
	if err != nil {
		matched := mongo.IsDuplicateKeyError(err)
		if matched {
			return "", errors.New("409")

		} else {
			return "", errors.New("failed to insert user")
		}
	}
	str := fmt.Sprintf("%v", res)
	id := strings.Split(str, "\"")[1]
	return id, nil
}

func (u *userService) GetUserById(id string, ctx context.Context) (*dto.GetUserStruct, error) {
	objid, _ := primitive.ObjectIDFromHex(id)
	return u.repo.GetUserById(objid, ctx)
}

func (u *userService) IsUsernameAvailable(username string, ctx context.Context) bool {
	return u.repo.IsUsernameAvailable(username, ctx)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
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
