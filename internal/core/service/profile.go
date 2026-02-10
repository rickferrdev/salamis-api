package service

import (
	"context"
	"errors"

	"github.com/rickferrdev/salamis-api/internal/core/domain"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
)

type profileService struct {
	storage ports.ProfileRepository
}

func NewProfileService(storage ports.ProfileRepository) ports.ProfileService {
	return &profileService{
		storage: storage,
	}
}

func (u *profileService) UpdateProfile(ctx context.Context, profile ports.ProfileInput) (*ports.ProfileOutput, error) {
	updated, err := u.storage.UpdateProfile(ctx, domain.ProfileDomain{
		Status:      profile.Status,
		AvatarURL:   profile.AvatarURL,
		Description: profile.Description,
	})
	if err != nil {
		if errors.Is(err, ports.ErrRecordNotFound) {
			return nil, ports.ErrProfileNotFound
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, ports.ErrTimeout
		}
		return nil, err
	}

	return &ports.ProfileOutput{
		Status:      updated.Status,
		AvatarURL:   updated.AvatarURL,
		Description: updated.Description,
	}, nil
}

func (u *profileService) GetProfileByUserID(ctx context.Context, userID string) (*ports.ProfileOutput, error) {
	exists, err := u.storage.FindProfileByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, ports.ErrProfileNotFound) {
			return nil, ports.ErrProfileNotFound
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, ports.ErrTimeout
		}
		return nil, err
	}

	if exists != nil {
		return &ports.ProfileOutput{
			Status:      exists.Status,
			AvatarURL:   exists.AvatarURL,
			Description: exists.Description,
		}, nil
	}

	return nil, ports.ErrProfileNotFound
}
