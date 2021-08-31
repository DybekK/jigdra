package user

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokenPair(objectId string) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	//TODO: remove redis logic to separate service
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	id := uuid.NewString()
	err := rdb.Set(ctx, id, objectId, 0).Err()
	if err != nil {
		panic(err)
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["identityKey"] = objectId
	claims["logoutKey"] = id
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()

	t, err := token.SignedString([]byte("test"))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["identityKey"] = objectId
	rtClaims["exp"] = time.Now().Add(time.Hour * 10).Unix()

	rt, err := refreshToken.SignedString([]byte("test"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}
