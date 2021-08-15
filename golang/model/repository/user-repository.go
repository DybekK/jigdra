package repository

import (
	"context"
	"fmt"
	"golang/model/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(*dto.User, context.Context) (string, error)
	GetUser(*dto.LoginUser, context.Context) (*dto.User, error)
	GetUserById(string, context.Context) (*dto.GetUserStruct, error)
}

var UserCollection *mongo.Collection = DBService.GetCollection(client, "users")

func NewUserRepository() UserRepository {
	DBService.Initialize()
	return &database{}
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
	return &result, nil

}

func (d *database) CreateUser(req_user *dto.User, ctx context.Context) (string, error) {
	result, insertError := UserCollection.InsertOne(ctx, &req_user)
	return fmt.Sprintf("%v", result), insertError
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
