package model

import (
	"context"
	"fmt"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	Client *mongo.Client
}

type UserInterface interface {
	Initialize(string) (*mongo.Client, error)
	ValidateEmail(string) error
	CreateUser(*User, context.Context) (*mongo.InsertOneResult, error)
	GetCollection(string) *mongo.Collection
	GetUser(*LoginUser, context.Context) (*User, error)
	GetMiddleWare() *jwt.GinJWTMiddleware
}

var (
	Interface UserInterface = &Database{}
)

func getConnection(uri string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Connected successfully")
	return client, nil
}

func (d *Database) Initialize(uri string) (*mongo.Client, error) {
	var err error
	d.Client, err = getConnection(uri)
	createIndex(d.GetCollection("users"))
	return d.Client, err
}

func (d *Database) GetCollection(collectionName string) *mongo.Collection {
	var collection *mongo.Collection = (*mongo.Collection)(d.Client.Database(os.Getenv("MONGO_INITDB_DATABASE")).Collection(collectionName))
	return collection
}

//this shouldn't be done programatically, will be removed later
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
