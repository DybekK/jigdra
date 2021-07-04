package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username *string            `json:"username" validate:"required"`
	Email    *string            `json:"email" validate:"required"`
	Password *string            `json:"password" validate:"required"`
}

var validate = validator.New()
var client = getConnection()
var userCollection = getCollection(client, "users")

func getUwa(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "success",
	})
}

func addUser(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var req_user User
	if err := c.BindJSON(&req_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationError := validate.Struct(req_user)
	if validationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
		return
	}
	var hash, hash_error = HashPassword(*req_user.Password)
	if hash_error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "shit broke once more",
		})
		return
	}
	var insert_user User
	insert_user.ID = primitive.NewObjectID()
	insert_user.Username = req_user.Username
	insert_user.Email = req_user.Email
	insert_user.Password = &hash

	result, insertError := userCollection.InsertOne(ctx, insert_user)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": insertError.Error(),
		})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result)

}

func main() {
	fmt.Println("lets go")
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/v1", getUwa)
	r.POST("/v1/add", addUser)
	r.GET("/v1/users", func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		cursor, error := userCollection.Find(ctx, bson.D{})
		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			return
		}
		defer cancel()
		var result []User

		for cursor.Next(context.TODO()) {
			var el User
			err := cursor.Decode(&el)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, el)
		}

		cursor.Close(context.TODO())

		c.JSON(http.StatusOK, result)
	})
	createIndex(userCollection)

	log.Fatal(r.Run(":4201"))
}

//this shouldn't be done programatically, will be removed later
func createIndex(coll *mongo.Collection) {
	index, err := coll.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(index)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getConnection() *mongo.Client {
	//needs to be changed to "mongodb://mongodb:27017" if you want to run it in docker
	uri := "mongodb://localhost:27017"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Connected successfully")
	return client
}

func getCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = (*mongo.Collection)(client.Database("temp_users").Collection(collectionName))
	return collection
}
