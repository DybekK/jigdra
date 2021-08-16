package main

import (
	"golang/model/dto"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c := &gin.Context{}
	_ = getRouter() //This inits userService and redirectService so yeah idc
	user := &dto.User{
		Username: "test",
		Name:     "Janusz",
		Surname:  "Kowalski",
		Email:    "test@mail.com",
		Password: "strongpasswd",
	}
	result, err := userService.CreateUser(user, c)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	insertedUser, err := userService.GetUserById(result, c)
	assert.Nil(t, err)
	assert.Equal(t, user.Name, insertedUser.Name)
	assert.Equal(t, user.Surname, insertedUser.Surname)
	assert.Equal(t, user.Username, insertedUser.Username)
	assert.Equal(t, user.Email, insertedUser.Email)
}

func TestCreateUserConflict(t *testing.T) {
	c := &gin.Context{}
	gin.SetMode(gin.TestMode)
	_ = getRouter()
	users := map[string]dto.User{
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
	result1, err := userService.CreateUser(&user1, c)
	assert.Nil(t, err)
	assert.NotNil(t, result1)
	//New user with different username but same email
	user2 := users["test2"]
	result2, err := userService.CreateUser(&user2, c)
	assert.NotNil(t, err)
	assert.Empty(t, result2)
	assert.Equal(t, "409", err.Error())
	//New user with different email but same username
	user3 := users["test3"]
	result3, err := userService.CreateUser(&user3, c)
	assert.NotNil(t, err)
	assert.Empty(t, result3)
	assert.Equal(t, "username taken", err.Error())

}

func TestGetUserById(t *testing.T) {
	c := &gin.Context{}
	_ = getRouter()
	gin.SetMode(gin.TestMode)
	user := &dto.User{
		Username: "test",
		Name:     "Janusz",
		Surname:  "Kowalski",
		Email:    "test@mail.com",
		Password: "strongpasswd",
	}
	uid, err := userService.CreateUser(user, c)
	assert.NotNil(t, uid)
	hex, err := redirectRepo.SecureRedirect(c, uid)
	assert.Nil(t, err)
	assert.NotNil(t, hex)
	id, err := redirectService.VerifyRedirect(c, hex)
	assert.Nil(t, err)
	assert.NotNil(t, id)
	user_resp, err := userService.GetUserById(id, c)
	assert.Nil(t, err)
	assert.NotNil(t, user_resp)
	assert.Equal(t, user.Username, user_resp.Username)
	assert.Equal(t, user.Email, user_resp.Email)
	assert.Equal(t, user.Name, user_resp.Name)
	assert.Equal(t, user.Surname, user_resp.Surname)

}
