package main

import (
	"fmt"
	"golang/model/dto"
	"golang/model/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestLoginReturns200IfExists(t *testing.T) {
	c := &gin.Context{}
	r := getRouter()
	h := &handler{}
	w := httptest.NewRecorder()
	defer func() {
		repository.UserCollection.DeleteMany(c, bson.M{})
		repository.RedirectCollection.DeleteMany(c, bson.M{})
	}()
	test := map[string]struct {
		payload    string
		statuscode int
	}{
		"register": {
			payload: `{
				"username":"test",
				"name": "Janusz",
				"surname":"Kowalski",
				"email":"test@mail.com",
				"password":"very_stronk"
			}`,
			statuscode: 302,
		},
		"login": {
			payload: `{
				"email":"test@mail.com",
				"password":"very_stronk"
			}`,
		},
	}

	r.POST("/v1/login", h.login)
	r.POST("/v1/register", h.addUser)

	req, err := http.NewRequest("POST", "/v1/register", strings.NewReader(test["register"].payload))
	r.ServeHTTP(w, req)
	assert.Nil(t, err)
	assert.Equal(t, 302, w.Result().StatusCode)
	w = httptest.NewRecorder()
	req, err = http.NewRequest("POST", "/v1/login", strings.NewReader(test["login"].payload))
	r.ServeHTTP(w, req)
	assert.Nil(t, err)
	assert.Equal(t, 200, w.Result().StatusCode)
}

func TestLoginReturnBadRequest(t *testing.T) {
	c := &gin.Context{}
	h := &handler{}
	r := getRouter()
	w := httptest.NewRecorder()
	defer func() {
		repository.UserCollection.DeleteMany(c, bson.M{})
		repository.RedirectCollection.DeleteMany(c, bson.M{})
	}()
	r.POST("/v1/login", h.login)
	req, _ := http.NewRequest("POST", "/v1/login", strings.NewReader(""))
	want := 400
	r.ServeHTTP(w, req)
	assert.Equal(t, want, w.Result().StatusCode)
}

func TestLoginReturnUnauthorized(t *testing.T) {
	c := &gin.Context{}
	h := &handler{}
	gin.SetMode(gin.TestMode)
	r := getRouter()
	w := httptest.NewRecorder()
	defer func() {
		repository.UserCollection.DeleteMany(c, bson.M{})
		repository.RedirectCollection.DeleteMany(c, bson.M{})
	}()
	tests := map[string]struct {
		payload      string
		expectedcode int
	}{
		"401": {
			payload:      `{"email": "doesntexist@mail.com", "password":"1234" }`,
			expectedcode: 401,
		},
	}
	r.POST("/v1/login", h.login)
	req, err := http.NewRequest(http.MethodPost, "/v1/login", strings.NewReader(tests["401"].payload))
	r.ServeHTTP(w, req)
	assert.Equal(t, nil, err)
	assert.NotNil(t, w.Result())
	assert.Equal(t, tests["401"].expectedcode, w.Result().StatusCode)
}

//Same hex value for redirect should return 401 after used once
func TestRedirectHexExpires(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c := &gin.Context{}
	h := &handler{}
	r := getRouter()
	w := httptest.NewRecorder()
	defer func() {
		repository.UserCollection.DeleteMany(c, bson.M{})
		repository.RedirectCollection.DeleteMany(c, bson.M{})
	}()
	userToRegister := dto.User{
		Username: "someusername",
		Name:     "Janusz",
		Surname:  "Kowalski",
		Email:    "uwa@mail.com",
		Password: "verystrongpasswd",
	}

	id, err := userRepo.CreateUser(&userToRegister, c)
	hex, err := redirectRepo.SecureRedirect(c, id)
	assert.Nil(t, err)
	r.GET("/v1/login", h.login)
	urlString := fmt.Sprintf("/v1/login?redirect=%s", hex)
	req, _ := http.NewRequest("GET", urlString, nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
	req, _ = http.NewRequest("GET", urlString, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)
}
