package ports

import (
	"context"
	"net/http"
)

// JWT Tokenizer
type TokenPayload struct {
	UserID uint
}

type TokenizerOutput struct {
	Token string
}

type Tokenizer interface {
	Generate(userID uint) (*TokenizerOutput, error)
	Verify(token string) (*TokenPayload, error)
}

// Password Hasher
type Hasher interface {
	Generate(password []byte) ([]byte, error)
	Compare(hash, password []byte) error
}

// HTTP Response Handler
type Response interface {
	Success(ctx context.Context, w http.ResponseWriter, status int, data any)
	Error(ctx context.Context, w http.ResponseWriter, err error)
}
