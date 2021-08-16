package repository

import (
	"context"
	"fmt"
	"golang/model/dto"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userData []dto.User

type MockUserRepo struct {
	mock.Mock
}

func (r MockUserRepo) CreateUser(user *dto.User, ctx context.Context) (string, error) {
	user.Id = primitive.NewObjectID()
	for _, val := range userData {
		if val.Email == user.Email {
			e := mongo.CommandError{Code: 11000}
			return "", e
		}
	}
	userData = append(userData, *user)

	return fmt.Sprintf("%v", user.Id), nil
}
func (r MockUserRepo) GetUser(login *dto.LoginUser, ctx context.Context) (*dto.User, error) {
	for _, val := range userData {
		if login.Email == val.Email {
			return &val, nil
		}
	}
	return nil, mongo.ErrNoDocuments
}
func (r MockUserRepo) GetUserById(id string, ctx context.Context) (*dto.GetUserStruct, error) {
	for _, val := range userData {
		if id == val.Id.Hex() {
			return &dto.GetUserStruct{
				Id:          val.Id,
				Email:       val.Email,
				Username:    val.Username,
				Name:        val.Name,
				Surname:     val.Surname,
				GenderId:    val.GenderId,
				DateOfBirth: val.DateOfBirth,
			}, nil
		}
	}
	return nil, mongo.ErrNoDocuments
}
func (r MockUserRepo) IsUsernameAvailable(username string, ctx context.Context) bool {
	for _, val := range userData {
		if username == val.Username {
			return false
		}
	}
	return true
}

var redirectData = []dto.Security{}

type MockRedirectRepo struct {
	mock mock.Mock
}

func (r MockRedirectRepo) SecureRedirect(ctx context.Context, id string) (string, error) {
	var sec dto.Security
	sec.Id = id
	randHex, _ := randomHex(20)
	sec.Hex = randHex
	fmt.Println(id)
	redirectData = append(redirectData, sec)
	return randHex, nil
}

func (r MockRedirectRepo) VerifyRedirect(ctx context.Context, hex string) (string, error) {
	id := ""
	index := 0
	for i, val := range redirectData {
		if hex == val.Hex {
			id = val.Id
			index = i
		}
	}
	if id != "" {
		redirectData[index] = redirectData[len(redirectData)-1]
		redirectData = redirectData[:len(redirectData)-1]
		return id, nil
	}

	return "", mongo.ErrNoDocuments
}

func Purge() {
	redirectData = nil
	userData = nil
}
