package commonclient

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongodbClient(uri, username, password string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := options.Credential{
		Username: username,
		Password: password,
	}

	return mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(credential))
}
