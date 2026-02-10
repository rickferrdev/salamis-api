package mongo

import (
	"context"

	"github.com/rickferrdev/salamis-api/internal/core/domain"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type postRepository struct {
	collection *mongo.Collection
}

func NewPostRepository(database *mongo.Database) ports.PostRepository {
	return &postRepository{
		collection: database.Collection("posts"),
	}
}

func (u *postRepository) CreatePost(ctx context.Context, post domain.PostDomain) (*domain.PostDomain, error) {
	schema := PostDomainToSchema(post)

	if _, err := u.collection.InsertOne(ctx, schema); err != nil {
		return nil, ErrorFully(err)
	}

	return schema.PostSchemaToDomain(), nil
}

func (u *postRepository) DeletePostByID(ctx context.Context, postID string) error {
	filter := bson.M{"_id": postID}

	res, err := u.collection.DeleteOne(ctx, filter)
	if err != nil {
		return ErrorFully(err)
	}

	if res.DeletedCount == 0 {
		return ErrorFully(mongo.ErrNoDocuments)
	}

	return nil
}
