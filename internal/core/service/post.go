package service

import (
	"context"
	"errors"

	"github.com/rickferrdev/salamis-api/internal/core/domain"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
)

type postService struct {
	repository ports.PostStorage
}

func NewPostService(repository ports.PostStorage) ports.PostService {
	return &postService{
		repository: repository,
	}
}

func (u *postService) Publish(ctx context.Context, post ports.PostInput) (*ports.PostOutput, error) {
	newPost := domain.PostDomain{
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID,
	}

	createdPost, err := u.repository.CreatePost(ctx, newPost)
	if err != nil {
		if errors.Is(err, ports.ErrConstraintViolation) {
			return nil, ports.ErrFailedToPublishPost
		}
		return nil, err
	}

	return &ports.PostOutput{
		ID:    createdPost.ID,
		Title: createdPost.Title,
	}, nil
}

func (u *postService) Delete(ctx context.Context, id uint) error {
	success, err := u.repository.DeletePostByID(ctx, id)
	if err != nil {
		if errors.Is(err, ports.ErrRecordNotFound) {
			return ports.ErrPostNotFound
		}
	}

	if success == nil {
		return ports.ErrPostNotFound
	}

	return nil
}
