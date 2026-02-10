package ports

import "errors"

// package hasher
var (
	ErrPasswordMismatch = errors.New("password mismatch")
)

// package tokenizer
var (
	ErrTokenGeneration = errors.New("failed to generate token")
	ErrInvalidToken    = errors.New("invalid token")
	ErrTokenExpired    = errors.New("token has expired")
)

// profile storage errors
var (
	ErrProfileNotFound     = errors.New("profile not found")
	ErrFailedUpdateProfile = errors.New("failed to update profile")
)

// database record errors
var (
	ErrRecordNotFound       = errors.New("record not found")
	ErrConstraintViolation  = errors.New("constraint violation")
	InternalError           = errors.New("an unexpected database error occurred")
	ServiceUnavailableError = errors.New("database is unreachable")
)

// context errors
var (
	ErrTimeout = errors.New("operation timed out")
)

// auth service errors
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
)

// handler errors
var (
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrInternalServerError = errors.New("internal server error")
)

// post service errors
var (
	ErrPostNotFound        = errors.New("post not found")
	ErrFailedToPublishPost = errors.New("failed to publish post")
)
