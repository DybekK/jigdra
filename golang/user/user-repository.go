package user

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

//factory

func NewUserRepository(client *mongo.Client) UserRepository {
	coll := client.Database(os.Getenv("MONGO_INITDB_DATABASE")).Collection("users")
	return UserRepository{client: client, coll: coll}
}

//methods

func (ur *UserRepository) GetUser(login *LoginUser, ctx context.Context) (*User, error) {
	res := ur.coll.FindOne(ctx, bson.M{
		"email": login.Email,
	})
	var result User
	decodeErr := res.Decode(&result)
	if decodeErr != nil {
		return nil, decodeErr
	}
	return &result, nil

}

func (ur *UserRepository) CreateUser(reqUser *User, ctx context.Context) (interface{}, error) {
	result, insertError := ur.coll.InsertOne(ctx, &reqUser)
	return result.InsertedID, insertError
}

func (ur *UserRepository) GetUserById(id primitive.ObjectID, ctx context.Context) (*GetUserStruct, error) {
	res := ur.coll.FindOne(ctx, bson.M{"_id": id})
	var user GetUserStruct
	decodeErr := res.Decode(&user)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return &user, nil

}

func (ur *UserRepository) IsUsernameAvailable(username string, ctx context.Context) bool {
	result := ur.coll.FindOne(ctx, bson.M{"username": username})
	return result.Err() == mongo.ErrNoDocuments
}
