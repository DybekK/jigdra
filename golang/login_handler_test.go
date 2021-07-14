package main

import (
	"fmt"
	"golang/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLoginReturns200IfExists(t *testing.T) {
	c := &gin.Context{}
	_ = godotenv.Load("tests.env")
	r := getRouter()
	h := &handler{}
	model.Interface.Initialize()
	w := httptest.NewRecorder()
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
	model.Interface.GetCollection("users").Drop(c)
}

func TestLoginReturnBadRequest(t *testing.T) {
	c := &gin.Context{}
	_ = godotenv.Load("tests.env")
	h := &handler{}
	model.Interface.Initialize()
	r := getRouter()
	w := httptest.NewRecorder()
	r.POST("/v1/login", h.login)
	req, _ := http.NewRequest("POST", "/v1/login", strings.NewReader(""))
	want := 400
	r.ServeHTTP(w, req)
	assert.Equal(t, want, w.Result().StatusCode)
	model.Interface.GetCollection("users").Drop(c)
}

func TestLoginReturnUnauthorized(t *testing.T) {
	c := &gin.Context{}
	_ = godotenv.Load("tests.env")
	h := &handler{}
	gin.SetMode(gin.TestMode)
	model.Interface.Initialize()
	r := getRouter()
	w := httptest.NewRecorder()
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
	model.Interface.GetCollection("users").Drop(c)
}

//Same hex value for redirect should return 401 after used once
func TestRedirectHexExpires(t *testing.T) {
	_ = godotenv.Load("tests.env")
	gin.SetMode(gin.TestMode)
	c := &gin.Context{}
	h := &handler{}
	model.Interface.Initialize()
	r := getRouter()
	w := httptest.NewRecorder()
	userToRegister := model.User{
		Username: "someusername",
		Name:     "Janusz",
		Surname:  "Kowalski",
		Email:    "uwa@mail.com",
		Password: "verystrongpasswd",
	}

	hex, err := model.Interface.CreateUser(&userToRegister, c)
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
	model.Interface.GetCollection("users").Drop(c)
}
