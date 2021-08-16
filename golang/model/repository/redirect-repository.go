package repository

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"golang/model/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RedirectRepository interface {
	SecureRedirect(context.Context, string) (string, error)
	VerifyRedirect(context.Context, string) (string, error)
}

var RedirectCollection *mongo.Collection

func NewRedirectRepository() RedirectRepository {
	DBService.Initialize()
	RedirectCollection = DBService.GetCollection(client, "redirect")
	return &database{}
}

func (d *database) SecureRedirect(ctx context.Context, id string) (string, error) {
	var sec dto.Security
	sec.Id = id
	randHex, _ := randomHex(20)
	sec.Hex = randHex
	_, err := RedirectCollection.InsertOne(ctx, sec)
	if err != nil {
		return "", err
	}

	return randHex, nil
}

func (d *database) VerifyRedirect(ctx context.Context, hex string) (string, error) {
	var sec dto.Security
	res := RedirectCollection.FindOneAndDelete(ctx, bson.M{"hex": hex})
	decode_err := res.Decode(&sec)
	if decode_err != nil {
		return "", decode_err
	}
	return sec.Id, nil
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
