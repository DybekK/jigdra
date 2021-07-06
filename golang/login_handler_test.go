package main

import (
	"context"
	"golang/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLoginReturns200IfExists(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = godotenv.Load("tests.env")
	r := getRouter()
	h := &handler{}
	middleware, err := getMiddleWare()
	model.Interface.Initialize()
	w := httptest.NewRecorder()
	if err != nil {
		t.Fail()
	}
	test := map[string]struct {
		payload    string
		statuscode int
	}{
		"register": {
			payload: `{
				"username":"test",
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

	r.POST("/v1/login", middleware.LoginHandler)
	r.POST("/v1/register", h.addUser)

	req1, err := http.NewRequest("POST", "/v1/register", strings.NewReader(test["register"].payload))
	r.ServeHTTP(w, req1)
	assert.Nil(t, err)
	assert.Equal(t, 302, w.Result().StatusCode)
	w2 := httptest.NewRecorder()
	req2, err := http.NewRequest("POST", "/v1/login", strings.NewReader(test["login"].payload))
	r.ServeHTTP(w2, req2)
	assert.Nil(t, err)
	assert.Equal(t, 200, w2.Result().StatusCode)
	defer cancel()
	model.Interface.GetCollection("users").Drop(ctx)
}

func TestLoginReturnBadRequest(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = godotenv.Load("tests.env")
	model.Interface.Initialize()
	r := getRouter()
	w := httptest.NewRecorder()
	middleware, err := getMiddleWare()
	if err != nil {
		t.Errorf("failed to get middleware")
	}
	r.POST("/v1/login", middleware.LoginHandler)
	req, _ := http.NewRequest("POST", "/v1/login", strings.NewReader(""))
	want := 400
	r.ServeHTTP(w, req)
	assert.Equal(t, want, w.Result().StatusCode)
	defer cancel()
	model.Interface.GetCollection("users").Drop(ctx)
}

func TestLoginReturnUnauthorized(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = godotenv.Load("tests.env")
	gin.SetMode(gin.TestMode)
	model.Interface.Initialize()
	r := getRouter()
	middleware, _ := getMiddleWare()
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
	r.POST("/v1/login", middleware.LoginHandler)
	req, err := http.NewRequest(http.MethodPost, "/v1/login", strings.NewReader(tests["401"].payload))
	r.ServeHTTP(w, req)
	assert.Equal(t, nil, err)
	assert.NotNil(t, w.Result())
	assert.Equal(t, tests["401"].expectedcode, w.Result().StatusCode)
	defer cancel()
	model.Interface.GetCollection("users").Drop(ctx)
}
