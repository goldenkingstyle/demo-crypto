package user

import (
	"github.com/goldenkingstyle/demo-crypto/internal/crypto"
)

type User struct {
	Name   string            `json:"name"`
	Wallet []crypto.CryptoID `json:"wallet"`
}

func NewUser(name string) *User {
	return &User{
		Name:   name,
		Wallet: []crypto.CryptoID{},
	}
}
