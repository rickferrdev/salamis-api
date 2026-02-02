package service

import (
	"context"
	"errors"

	"github.com/rickferrdev/salamis-api/internal/core/domain"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
)

type authService struct {
	tokenizer ports.Tokenizer
	hasher    ports.Hasher
	storage   ports.AuthStorage
}

func NewAuthService(tokenizer ports.Tokenizer, hasher ports.Hasher, storage ports.AuthStorage) ports.AuthService {
	return &authService{
		hasher:    hasher,
		tokenizer: tokenizer,
		storage:   storage,
	}
}

func (u *authService) Login(ctx context.Context, input ports.AuthInput) (*ports.AuthOutput, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	receive, err := u.storage.FindUserByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, ports.ErrRecordNotFound) {
			return nil, ports.ErrUserNotFound
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, ports.ErrTimeout
		}

		return nil, err
	}

	err = u.hasher.Compare([]byte(receive.Password), []byte(input.Password))
	if err != nil {
		if errors.Is(err, ports.ErrPasswordMismatch) {
			return nil, ports.ErrInvalidCredentials
		}

		return nil, err
	}

	token, err := u.tokenizer.Generate(receive.ID)
	if err != nil {
		if errors.Is(err, ports.ErrTokenGeneration) {
			return nil, ports.ErrTokenGeneration
		}

		return nil, err
	}

	return &ports.AuthOutput{
		User: ports.UserOutput{
			Nickname: receive.Username,
			Username: receive.Nickname,
		},
		Token: token.Token,
	}, nil
}

func (u *authService) Register(ctx context.Context, input ports.AuthInput) (*ports.AuthOutput, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	hash, err := u.hasher.Generate([]byte(input.Password))
	if err != nil {
		return nil, err
	}

	user, err := u.storage.Create(ctx, domain.UserDomain{
		Nickname: input.Nickname,
		Password: string(hash),
		Username: input.Username,
		Email:    input.Email,
	})
	if err != nil {
		if errors.Is(err, ports.ErrConstraintViolation) {
			return nil, ports.ErrUserAlreadyExists
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, ports.ErrTimeout
		}

		return nil, err
	}

	return &ports.AuthOutput{
		User: ports.UserOutput{
			Nickname: user.Nickname,
			Username: user.Username,
		},
	}, nil
}
