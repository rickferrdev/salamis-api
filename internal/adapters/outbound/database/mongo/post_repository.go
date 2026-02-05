package mongo

import (
	"context"
	"errors"

	"github.com/rickferrdev/salamis-api/internal/core/domain"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
	"gorm.io/gorm"
)

func NewPostStorage(database *database) (ports.PostStorage, error) {
	err := database.AutoMigrate(&PostSchema{})
	if err != nil {
		return nil, err
	}

	return &postStorage{
		database: database,
	}, nil
}

func (u *postStorage) CreatePost(ctx context.Context, post domain.PostDomain) (*domain.PostDomain, error) {
	model := PostFromSchema(&post)

	err := u.database.WithContext(ctx).Create(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ports.ErrConstraintViolation
		}
		return nil, err
	}

	return model.FromDomain(), nil
}
func (u *postStorage) DeletePostByID(ctx context.Context, id uint) error {
	err := u.database.WithContext(ctx).Delete(&PostSchema{}, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ports.ErrRecordNotFound
		}
		return err
	}

	return nil
}
