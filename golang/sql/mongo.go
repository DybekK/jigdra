package sql

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getConnection(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
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

func InitMongoDatabase() *mongo.Client {
	if strings.HasSuffix(os.Args[0], ".test") {
		_ = godotenv.Load("tests.env")
	}
	db_user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	db_passwd := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	db_host := os.Getenv("MONGO_HOST")
	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017", db_user, db_passwd, db_host)
	client, err := getConnection(uri)
	if err != nil {
		panic(err)
	}
	createIndex(GetCollection(client, "users"))
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = (*mongo.Collection)(client.Database(os.Getenv("MONGO_INITDB_DATABASE")).Collection(collectionName))
	return collection
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
