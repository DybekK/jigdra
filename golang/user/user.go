package user

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	Email    string `json:"email" validate:"required"`
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
