package mongo

import (
	"context"
	"errors"
	"log/slog"

	"github.com/rickferrdev/salamis-api/internal/core/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

func ErrorFully(err error) error {
	if err == nil {
		return nil
	}

	slog.Error("database error occurred",
		"details", err.Error(),
		"type", slog.Any("type", err),
	)

	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return ports.ErrTimeout
	}

	if mongo.IsDuplicateKeyError(err) {
		return ports.ErrConstraintViolation
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return ports.ErrRecordNotFound
	}

	if mongo.IsNetworkError(err) {
		return ports.ServiceUnavailableError
	}
	return ports.InternalError
}
