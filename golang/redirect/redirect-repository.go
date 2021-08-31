package redirect

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RedirectRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

//factory

func NewRedirectRepository(client *mongo.Client) RedirectRepository {

	coll := client.Database(os.Getenv("MONGO_INITDB_DATABASE")).Collection("redirects")
	return RedirectRepository{client: client, coll: coll}
}

//methods

func (rr *RedirectRepository) SecureRedirect(ctx context.Context, sec Security) (string, error) {
	_, err := rr.coll.InsertOne(ctx, sec)
	if err != nil {
		return "", err
	}

	return sec.Hex, nil
}

func (rr *RedirectRepository) VerifyRedirect(ctx context.Context, hex string) (string, error) {
	var sec Security
	res := rr.coll.FindOneAndDelete(ctx, bson.M{"hex": hex})
	decode_err := res.Decode(&sec)
	if decode_err != nil {
		return "", decode_err
	}
	return sec.Id, nil
}
