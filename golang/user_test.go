package main

import (
	"context"
	"golang/model"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser(t *testing.T) {
	_ = godotenv.Load("tests.env")
	gin.SetMode(gin.TestMode)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	model.Interface.Initialize()
	user := &model.User{
		Id:       primitive.NewObjectID(),
		Username: "test",
		Email:    "test@mail.com",
		Password: "strong",
	}
	defer cancel()
	result, err := model.Interface.CreateUser(user, ctx)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	model.Interface.GetCollection("users").Drop(ctx)
}

func TestCreateUserConflict(t *testing.T) {
	_ = godotenv.Load("tests.env")
	gin.SetMode(gin.TestMode)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	model.Interface.Initialize()
	users := map[string]model.User{
		"test1": {
			Id:       primitive.NewObjectID(),
			Username: "test",
			Email:    "test@mail.com",
			Password: "strong",
		},
		"test2": {
			Id:       primitive.NewObjectID(),
			Username: "test2",
			Email:    "test@mail.com",
			Password: "strong",
		},
		"test3": {
			Id:       primitive.NewObjectID(),
			Username: "test",
			Email:    "test3@mail.com",
			Password: "strong",
		},
	}
	//New user
	user1 := users["test1"]
	result1, err := model.Interface.CreateUser(&user1, ctx)
	assert.Nil(t, err)
	assert.NotNil(t, result1)
	//New user with different username but same email
	user2 := users["test2"]
	result2, err := model.Interface.CreateUser(&user2, ctx)
	assert.NotNil(t, err)
	assert.Nil(t, result2)
	assert.Equal(t, "email in use", err.Error())
	//New user with different email but same email
	user3 := users["test3"]
	result3, err := model.Interface.CreateUser(&user3, ctx)
	assert.NotNil(t, err)
	assert.Nil(t, result3)
	assert.Equal(t, "username taken", err.Error())
	defer cancel()
	model.Interface.GetCollection("users").Drop(ctx)

}
