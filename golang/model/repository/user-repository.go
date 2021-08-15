package repository

import (
	"context"
	"errors"
	"fmt"
	"golang/model/dto"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userServ struct{}

type UserRepository interface {
	ValidateEmail(string) error
	CreateUser(*dto.User, context.Context) (string, error)
	GetUser(*dto.LoginUser, context.Context) (*dto.User, error)
	GetUserById(string, context.Context) (*dto.GetUserStruct, error)
}

var UserCollection *mongo.Collection = DBService.GetCollection(client, "users")

func NewUserRepository() UserRepository {
	DBService.Initialize()
	return &database{}
}

func (d *database) ValidateEmail(email string) error {
	match := dto.Email_regex.MatchString(email)
	if match {
		return nil
	} else {
		return errors.New("invalid email")
	}
}

func (d *database) GetUser(login *dto.LoginUser, ctx context.Context) (*dto.User, error) {
	res := UserCollection.FindOne(ctx, bson.M{
		"email": login.Email,
	})

	var result dto.User
	decode_err := res.Decode(&result)
	if decode_err != nil {
		return nil, decode_err
	}

}

func (d *database) CreateUser(req_user *dto.User, ctx context.Context) (string, error) {
	req_user.Id = primitive.NewObjectID()
	email_err := d.ValidateEmail(req_user.Email)
	if email_err != nil {
		return "", email_err
	}

	var hash, hash_error = hashPassword(req_user.Password)
	if hash_error != nil {
		return "", errors.New("hashing error")
	}

	if availible := isUsernameAvailable(req_user.Username, ctx); !availible {
		return "", errors.New("username taken")
	}
	req_user.Password = hash
	result, insertError := UserCollection.InsertOne(ctx, &req_user)
	if insertError != nil {
		matched, _ := regexp.MatchString(`duplicate key`, insertError.Error())
		if matched {
			return "", errors.New("409")

		} else {
			return "", errors.New("failed to insert user")
		}
	}
	id := strings.Split(fmt.Sprintf("%v", result), "\"")[1]
	// hex, err := RedirectService.SecureRedirect(ctx, id)
	// if err != nil {
	// 	return "", errors.New("unable to redirect")
	// }

	return id, nil
}

func (d *database) GetUserById(id string, ctx context.Context) (*dto.GetUserStruct, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	res := UserCollection.FindOne(ctx, bson.M{"_id": objId})
	var user dto.GetUserStruct
	decode_err := res.Decode(&user)
	if decode_err != nil {
		return nil, decode_err
	}

	return &user, nil

}
