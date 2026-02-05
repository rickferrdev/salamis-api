package mongo

import (
	"context"
	"time"

	"github.com/rickferrdev/salamis-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDatabase(env *config.Env) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(env.AppDBUri)
	client, err := mongo.Connect(ctx, clientOptions)

	return client.Database("salamis"), err
}
