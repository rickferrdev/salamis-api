package mongo

import (
	"github.com/rickferrdev/salamis-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type PostSchema struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
	UserID  primitive.ObjectID `bson:"author_id"`
}

func (PostSchema) TableName() string {
	return "posts"
}

func (u *PostSchema) FromDomain() *domain.PostDomain {
	return &domain.PostDomain{
		Title:   u.Title,
		Content: u.Content,
		UserID:  u.UserID.Hex(),
	}
}

func PostFromSchema(post *domain.PostDomain) *PostSchema {
	return &PostSchema{
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.AuthorID,
	}
}

type postStorage struct {
	database *gorm.DB
}
