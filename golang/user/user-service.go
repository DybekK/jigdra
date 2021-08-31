package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo UserRepository
}

//factory

func NewUserService(repo UserRepository) UserService {
	return UserService{Repo: repo}
}

//methods

func (us *UserService) GetUser(login *LoginUser, ctx context.Context) (*User, error) {
	user, err := us.Repo.GetUser(login, ctx)
	if err != nil {
		return nil, err
	}
	if verifyPassword(login.Password, user.Password) {
		return user, nil
	} else {
		return nil, errors.New("invalid password")
	}
}

func (us *UserService) CreateUser(reqUser *User, ctx context.Context) (string, error) {
	reqUser.Id = primitive.NewObjectID()
	emailErr := validateEmail(reqUser.Email)
	if emailErr != nil {
		return "", emailErr
	}

	var hash, hashError = hashPassword(reqUser.Password)
	if hashError != nil {
		return "", errors.New("hashing error")
	}

	if available := us.Repo.IsUsernameAvailable(reqUser.Username, ctx); !available {
		return "", errors.New("username taken")
	}
	reqUser.Password = hash
	res, err := us.Repo.CreateUser(reqUser, ctx)
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

func (us *UserService) GetUserById(id string, ctx context.Context) (*GetUserStruct, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	return us.Repo.GetUserById(objId, ctx)
}

func (us *UserService) IsUsernameAvailable(username string, ctx context.Context) bool {
	return us.Repo.IsUsernameAvailable(username, ctx)
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
	match := EmailRegex.MatchString(email)
	if match {
		return nil
	} else {
		return errors.New("invalid email")
	}
}
