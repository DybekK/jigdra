package repository

import (
	"context"
	"golang/model/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(*dto.User, context.Context) (interface{}, error)
	GetUser(*dto.LoginUser, context.Context) (*dto.User, error)
	GetUserById(primitive.ObjectID, context.Context) (*dto.GetUserStruct, error)
	IsUsernameAvailable(string, context.Context) bool
}

var UserCollection *mongo.Collection

func NewUserRepository() UserRepository {
	client = DBService.Initialize()
	UserCollection = DBService.GetCollection(client, "users")
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

func (d *database) CreateUser(req_user *dto.User, ctx context.Context) (interface{}, error) {
	result, insertError := UserCollection.InsertOne(ctx, &req_user)
	return result.InsertedID, insertError
}

func (d *database) GetUserById(id primitive.ObjectID, ctx context.Context) (*dto.GetUserStruct, error) {
	res := UserCollection.FindOne(ctx, bson.M{"_id": id})
	var user dto.GetUserStruct
	decode_err := res.Decode(&user)
	if decode_err != nil {
		return nil, decode_err
	}

	return &user, nil

}

func (d *database) IsUsernameAvailable(username string, ctx context.Context) bool {
	result := UserCollection.FindOne(ctx, bson.M{"username": username})
	return result.Err() == mongo.ErrNoDocuments
}
