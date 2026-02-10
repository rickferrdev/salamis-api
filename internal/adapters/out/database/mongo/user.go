package mongo

import (
	"context"

	"github.com/rickferrdev/salamis-api/internal/core/domain"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) ports.UserRepository {
	return &userRepository{
		collection: database.Collection("users"),
	}
}

func (u *userRepository) Create(ctx context.Context, user domain.UserDomain) (*domain.UserDomain, error) {
	schema := UserDomainToSchema(user)

	if _, err := u.collection.InsertOne(ctx, schema); err != nil {
		return nil, ErrorFully(err)
	}

	return schema.UserSchemaToDomain(), nil
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain.UserDomain, error) {
	var user UserSchema

	if err := u.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, ErrorFully(err)
	}

	return user.UserSchemaToDomain(), nil
}

func (u *userRepository) UpdateUserByID(ctx context.Context, userID string, user domain.UserDomain) (*domain.UserDomain, error) {
	schema := UserDomainToSchema(user)
	updateData, err := bson.Marshal(schema)
	if err != nil {
		return nil, ErrorFully(err)
	}

	var updateMap bson.M
	if err := bson.Unmarshal(updateData, &updateMap); err != nil {
		return nil, ErrorFully(err)
	}
	delete(updateMap, "_id")

	update := bson.M{"$set": updateMap}
	if _, err := u.collection.UpdateByID(ctx, userID, update); err != nil {
		return nil, ErrorFully(err)
	}

	return schema.UserSchemaToDomain(), nil
}

func (u *userRepository) DeleteUserByID(ctx context.Context, userID string) error {
	filter := bson.M{"_id": userID}

	res, err := u.collection.DeleteOne(ctx, filter)
	if err != nil {
		return ErrorFully(err)
	}

	if res.DeletedCount == 0 {
		return ErrorFully(mongo.ErrNoDocuments)
	}

	return nil
}
