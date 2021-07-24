package model

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type redirect struct{}

type redirectService interface {
	SecureRedirect(context.Context, string) (string, error)
	VerifyRedirect(context.Context, string) (string, error)
}

var RedirectService redirectService = &redirect{}
var RedirectCollection *mongo.Collection = DBService.GetCollection(client, "redirect")

func (r *redirect) SecureRedirect(ctx context.Context, id string) (string, error) {
	var sec Security
	sec.Id = id
	randHex, _ := randomHex(20)
	sec.Hex = randHex
	_, err := RedirectCollection.InsertOne(ctx, sec)
	if err != nil {
		return "", err
	}

	return randHex, nil
}

func (r *redirect) VerifyRedirect(ctx context.Context, hex string) (string, error) {
	var sec Security
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
