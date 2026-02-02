package hasher

import (
	"github.com/rickferrdev/salamis-api/internal/core/ports"
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct {
	Cost int
}

func NewHasher() ports.Hasher {
	return &Hasher{
		Cost: bcrypt.DefaultCost,
	}
}

func (u *Hasher) Generate(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, u.Cost)
}

func (u *Hasher) Compare(hash []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return ports.ErrPasswordMismatch
	}
	return nil
}
