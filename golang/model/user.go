package model

import (
	"context"
	"errors"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username *string            `json:"username" validate:"required"`
	Email    *string            `json:"email" validate:"required"`
	Password *string            `json:"password" validate:"required"`
}

type LoginUser struct {
	Email    *string `json:"email" validate: "required"`
	Password *string `json:"password" validate:"required"`
}

var email_regex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

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
		return nil, errors.New("decode error")
	}
	if verifyPassword(*login.Password, *result.Password) {
		return &result, nil
	} else {
		return nil, errors.New("Invalid password")
	}
}

func (d *Database) CreateUser(req_user *User, ctx context.Context) (*mongo.InsertOneResult, error) {
	var user User
	user.ID = primitive.NewObjectID()
	email_err := d.ValidateEmail(*req_user.Email)
	if email_err != nil {
		return nil, email_err
	}
	user.Username = req_user.Username
	user.Email = req_user.Email

	var hash, hash_error = hashPassword(*req_user.Password)
	if hash_error != nil {
		return nil, errors.New("hashing error")
	}

	user.Password = &hash

	result, insertError := d.GetCollection("temp_users").InsertOne(ctx, user)
	if insertError != nil {
		matched, _ := regexp.MatchString(`duplicate key`, insertError.Error())
		if matched {
			return nil, errors.New("email already in use")

		} else {
			return nil, errors.New(insertError.Error())
		}
	}
	return result, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
