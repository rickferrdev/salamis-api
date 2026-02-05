package mongo

import (
	"context"
	"errors"

	"github.com/rickferrdev/salamis-api/internal/core/domain"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type userRepository struct {
	database *mongo.Database
}

func MewUserRepository(database *mongo.Database) (ports.UserRepository, error) {
	return &userRepository{
		database: database,
	}, nil
}

func (u *userRepository) UpdateByID(ctx context.Context, id uint, user domain.UserDomain) (*domain.UserDomain, error) {
	model := UserFromSchema(&user)

	err := u.database.WithContext(ctx).Model(&UserSchema{}).Where("id = ?", id).Updates(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ports.ErrRecordNotFound
		}
		return nil, err
	}

	return model.FromDomain(), nil
}
func (u *userRepository) FindByID(ctx context.Context, id uint) (*domain.UserDomain, error) {
	model := &UserSchema{}

	err := u.database.WithContext(ctx).Where("id = ?", id).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ports.ErrRecordNotFound
		}
		return nil, err
	}

	return model.FromDomain(), nil
}

func (u *userRepository) Create(ctx context.Context, user domain.UserDomain) (*domain.UserDomain, error) {
	model := UserFromSchema(&user)

	if model.Profile.Status == "" {
		model.Profile.Status = "Hey there! I am using Salamis."
	}

	if model.Profile.AvatarURL == "" {
		model.Profile.AvatarURL = "https://example.com/default-avatar.png"
	}

	if model.Profile.Description == "" {
		model.Profile.Description = "This is my profile description."
	}

	res, err := u.database.Collection("users").InsertOne(ctx, bson.D{})

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, ports.ErrConstraintViolation
		}
	}

	return model.FromDomain(), nil
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain.UserDomain, error) {
	var model UserSchema

	err := u.database.WithContext(ctx).Where("email = ?", email).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ports.ErrRecordNotFound
		}
		return nil, err
	}

	return model.FromDomain(), nil
}

func (u *userRepository) DeleteByID(ctx context.Context, id uint) error {
	err := u.database.WithContext(ctx).Delete(&UserSchema{}, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ports.ErrRecordNotFound
		}
		return err
	}

	return nil
}
