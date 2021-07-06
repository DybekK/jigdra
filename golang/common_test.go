package main

import (
	"context"
	"fmt"
	"golang/model"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func getMiddleWare() (*jwt.GinJWTMiddleware, error) {
	middleware, err := jwt.New(model.Interface.GetMiddleWare())
	if err != nil {
		return nil, err
	}
	return middleware, nil
}

func getMongoClient() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:password@localhost:27017"))
	if err != nil {
		panic(err)
	}
	return client
}

func createIndex(coll *mongo.Collection) {
	//Username should also be unique
	for _, key := range []string{"username", "email"} {
		index, err := coll.Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys:    bson.D{{Key: key, Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		)

		if err != nil {
			panic(err)
		}

		fmt.Println(index)
	}
}

func GetUserCollection() *mongo.Collection {
	client := getMongoClient()
	coll := client.Database("jigdra_test").Collection("users")
	createIndex(coll)
	return coll
}
