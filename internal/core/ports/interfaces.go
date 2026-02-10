package ports

import (
	"context"

	"github.com/rickferrdev/salamis-api/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.UserDomain) (*domain.UserDomain, error)
	FindUserByEmail(ctx context.Context, email string) (*domain.UserDomain, error)
	UpdateUserByID(ctx context.Context, userID string, user domain.UserDomain) (*domain.UserDomain, error)
	DeleteUserByID(ctx context.Context, userID string) error
}

type AuthService interface {
	Login(ctx context.Context, input AuthInput) (*AuthOutput, error)
	Register(ctx context.Context, input AuthInput) (*AuthOutput, error)
}

type PostService interface {
	Publish(ctx context.Context, post PostInput) (*PostOutput, error)
	Delete(ctx context.Context, postID string) error
}

type PostRepository interface {
	CreatePost(ctx context.Context, post domain.PostDomain) (*domain.PostDomain, error)
	DeletePostByID(ctx context.Context, postID string) error
}

type ProfileService interface {
	UpdateProfile(ctx context.Context, profile ProfileInput) (*ProfileOutput, error)
	GetProfileByUserID(ctx context.Context, userID string) (*ProfileOutput, error)
}

type ProfileRepository interface {
	UpdateProfile(ctx context.Context, profile domain.ProfileDomain) (*domain.ProfileDomain, error)
	FindProfileByUserID(ctx context.Context, userID string) (*domain.ProfileDomain, error)
}
