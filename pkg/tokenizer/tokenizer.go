package tokenizer

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rickferrdev/salamis-api/config"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
)

type Tokenizer struct {
	Secret []byte
}

func NewTokenizer(Env *config.Env) ports.Tokenizer {
	return &Tokenizer{
		Secret: []byte(Env.AppSecretJWT),
	}
}

type UserClaims struct {
	UserID uint
	jwt.RegisteredClaims
}

func (u *Tokenizer) Generate(userID uint) (*ports.TokenizerOutput, error) {
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "salamis-api",
			Subject:   "user",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Minute).UTC()),
		},
	})

	token, err := tokenJWT.SignedString(u.Secret)
	if err != nil {
		return nil, ports.ErrTokenGeneration
	}

	return &ports.TokenizerOutput{Token: token}, nil
}

func (u *Tokenizer) Verify(token string) (*ports.TokenPayload, error) {
	tokenJWT, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ports.ErrInvalidToken
		}

		return u.Secret, nil
	})
	if err != nil {
		return nil, ports.ErrInvalidToken
	}

	claims, ok := tokenJWT.Claims.(*UserClaims)
	if !ok {
		return nil, ports.ErrInvalidToken
	}

	return &ports.TokenPayload{
		UserID: claims.UserID,
	}, nil
}
