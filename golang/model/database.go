package model

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	Client *mongo.Client
}

type Security struct {
	Id  string `json:"id"`
	Hex string `json:"hex"`
}

type UserInterface interface {
	Initialize() (*mongo.Client, error)
	ValidateEmail(string) error
	CreateUser(*User, context.Context) (string, error)
	GetCollection(string) *mongo.Collection
	GetUser(*LoginUser, context.Context) (*User, error)
	GetUserById(string, context.Context) (*GetUserStruct, error)
	SecureRedirect(context.Context, string) (string, error)
	VerifyRedirect(context.Context, string) (string, error)
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

func (d *Database) Initialize() (*mongo.Client, error) {
	db_user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	db_passwd := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	db_host := os.Getenv("MONGO_HOST")
	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017", db_user, db_passwd, db_host)
	var err error
	d.Client, err = getConnection(uri)
	createIndex(d.GetCollection("users"))
	return d.Client, err
}

func (d *Database) GetCollection(collectionName string) *mongo.Collection {
	var collection *mongo.Collection = (*mongo.Collection)(d.Client.Database(os.Getenv("MONGO_INITDB_DATABASE")).Collection(collectionName))
	return collection
}

func (d *Database) SecureRedirect(ctx context.Context, id string) (string, error) {
	coll := d.GetCollection("redirect")
	var sec Security
	sec.Id = id
	randHex, _ := randomHex(20)
	sec.Hex = randHex
	_, err := coll.InsertOne(ctx, sec)
	if err != nil {
		return "", err
	}

	return randHex, nil
}

func (d *Database) VerifyRedirect(ctx context.Context, hex string) (string, error) {
	coll := d.GetCollection("redirect")
	var sec Security
	res := coll.FindOneAndDelete(ctx, bson.M{"hex": hex})
	decode_err := res.Decode(&sec)
	if decode_err != nil {
		return "", decode_err
	}
	return sec.Id, nil
}

//this shouldn't be done programatically, will be removed later
func createIndex(coll *mongo.Collection) {
	_, err := coll.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		panic(err)
	}
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
