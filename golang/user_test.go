package main

import (
	"golang/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser(t *testing.T) {
	c := &gin.Context{}
	_ = godotenv.Load("tests.env")
	gin.SetMode(gin.TestMode)
	model.Interface.Initialize()
	user := &model.User{
		Id:       primitive.NewObjectID(),
		Username: "test",
		Name:     "Janusz",
		Surname:  "Kowalski",
		Email:    "test@mail.com",
		Password: "strongpasswd",
	}
	result, err := model.Interface.CreateUser(user, c)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	model.Interface.GetCollection("users").Drop(c)
}

func TestCreateUserConflict(t *testing.T) {
	c := &gin.Context{}
	_ = godotenv.Load("tests.env")
	gin.SetMode(gin.TestMode)
	model.Interface.Initialize()
	users := map[string]model.User{
		"test1": {
			Id:       primitive.NewObjectID(),
			Username: "test",
			Name:     "Janusz",
			Surname:  "Kowalski",
			Email:    "test@mail.com",
			Password: "strong",
		},
		"test2": {
			Id:       primitive.NewObjectID(),
			Username: "test2",
			Name:     "Janusz",
			Surname:  "Kowalski",
			Email:    "test@mail.com",
			Password: "strong",
		},
		"test3": {
			Id:       primitive.NewObjectID(),
			Username: "test",
			Name:     "Janusz",
			Surname:  "Kowalski",
			Email:    "test3@mail.com",
			Password: "strong",
		},
	}
	//New user
	user1 := users["test1"]
	result1, err := model.Interface.CreateUser(&user1, c)
	assert.Nil(t, err)
	assert.NotNil(t, result1)
	//New user with different username but same email
	user2 := users["test2"]
	result2, err := model.Interface.CreateUser(&user2, c)
	assert.NotNil(t, err)
	assert.Empty(t, result2)
	assert.Equal(t, "email in use", err.Error())
	//New user with different email but same email
	user3 := users["test3"]
	result3, err := model.Interface.CreateUser(&user3, c)
	assert.NotNil(t, err)
	assert.Empty(t, result3)
	assert.Equal(t, "username taken", err.Error())
	model.Interface.GetCollection("users").Drop(c)

}
