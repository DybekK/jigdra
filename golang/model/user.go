package model

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Username    string             `json:"username" validate:"required,max=32"`
	Name        string             `json:"name" validate:"required,max=32"`
	Surname     string             `json:"surname" validate:"required,max=32"`
	Email       string             `json:"email" validate:"required,max=255"`
	Password    string             `json:"password" validate:"required,min=6,max=20"`
	GenderId    int                `json:"genderId"`
	DateOfBirth string             `json:"dateOfBirth"`
}

type LoginUser struct {
	Email    string `json:"email" validate: "required"`
	Password string `json:"password" validate:"required"`
}

type GetUserStruct struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Username    string             `json:"username" validate:"required"`
	Name        string             `json:"name" validate:"required,max=32"`
	Surname     string             `json:"surname" validate:"required,max=32"`
	Email       string             `json:"email" validate:"required"`
	GenderId    int                `json:"genderId"`
	DateOfBirth string             `json:"dateOfBirth"`
}

func (d *Database) ValidateEmail(email string) error {
	match := email_regex.MatchString(email)
	if match {
		return nil
	} else {
		return errors.New("invalid email")
	}
}

func (d *Database) GetUser(login *LoginUser, ctx context.Context) (*User, error) {
	coll := d.GetCollection("users")

	res := coll.FindOne(ctx, bson.M{
		"email": login.Email,
	})

	var result User
	decode_err := res.Decode(&result)
	if decode_err != nil {
		return nil, decode_err
	}
	if verifyPassword(login.Password, result.Password) {
		return &result, nil
	} else {
		return nil, errors.New("invalid password")
	}
}

func (d *Database) CreateUser(req_user *User, ctx context.Context) (string, error) {
	var user User
	user.Id = primitive.NewObjectID()
	email_err := d.ValidateEmail(req_user.Email)
	if email_err != nil {
		return "", email_err
	}
	user.Username = req_user.Username
	user.Email = req_user.Email
	user.DateOfBirth = req_user.DateOfBirth
	user.GenderId = req_user.GenderId
	user.Name = req_user.Name
	user.Surname = req_user.Surname

	var hash, hash_error = hashPassword(req_user.Password)
	if hash_error != nil {
		return "", errors.New("hashing error")
	}

	if availible := d.isUsernameAvailable(user.Username, ctx); !availible {
		return "", errors.New("username taken")
	}
	user.Password = hash
	result, insertError := d.GetCollection("users").InsertOne(ctx, user)
	if insertError != nil {
		matched, _ := regexp.MatchString(`duplicate key`, insertError.Error())
		if matched {
			return "", errors.New("409")

		} else {
			return "", errors.New("failed to insert user")
		}
	}
	id := strings.Split(fmt.Sprintf("%v", result), "\"")[1]
	hex, err := d.SecureRedirect(ctx, id)
	if err != nil {
		return "", errors.New("unable to redirect")
	}

	return hex, nil
}

var validate = validator.New()

func (d *Database) GetUserById(id string, ctx context.Context) (*GetUserStruct, error) {
	coll := d.GetCollection("users")
	objId, _ := primitive.ObjectIDFromHex(id)
	res := coll.FindOne(ctx, bson.M{"_id": objId})
	var user GetUserStruct
	decode_err := res.Decode(&user)
	if decode_err != nil {
		return nil, decode_err
	}

	return &user, nil

}

func (d *Database) isUsernameAvailable(username string, ctx context.Context) bool {
	coll := d.GetCollection("users")
	result := coll.FindOne(ctx, bson.M{"username": username})
	return result.Err() == mongo.ErrNoDocuments

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
