package service

import (
	"context"
	"errors"

	"github.com/rickferrdev/salamis-api/internal/core/domain"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
)

type postService struct {
	repository ports.PostRepository
}

func NewPostService(repository ports.PostRepository) ports.PostService {
	return &postService{
		repository: repository,
	}
}

func (u *postService) Publish(ctx context.Context, post ports.PostInput) (*ports.PostOutput, error) {
	newPost := domain.PostDomain{
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.User,
	}

	createdPost, err := u.repository.CreatePost(ctx, newPost)
	if err != nil {
		if errors.Is(err, ports.ErrConstraintViolation) {
			return nil, ports.ErrFailedToPublishPost
		}
		return nil, err
	}

	return &ports.PostOutput{
		ID:      createdPost.ID,
		Title:   createdPost.Title,
		Content: createdPost.Content,
		User:    ports.UserOutput{},
	}, nil
}

func (u *postService) Delete(ctx context.Context, postID string) error {
	err := u.repository.DeletePostByID(ctx, postID)
	if err != nil {
		if errors.Is(err, ports.ErrRecordNotFound) {
			return ports.ErrPostNotFound
		}

		return err
	}
	return nil
}
