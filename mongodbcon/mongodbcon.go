package mongodbcon

import (
	"context"
	"log"

	httperrors "github.com/myrachanto/erroring"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx = context.TODO()
)

func Mongodb(url, dbname string) (*mongo.Database, httperrors.HttpErr) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, httperrors.NewBadRequestError("Could not connect to mongodb")
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, httperrors.NewBadRequestError("Failed to ping")
	}
	db := client.Database(dbname)
	return db, nil
}
func DbClose(client *mongo.Client) {
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
