package user

import (
	"fmt"

	"github.com/goldenkingstyle/demo-crypto/internal/crypto"
)

const DEFAULT_BALANCE float64 = 10000

type User struct {
	Name    string         `json:"name"`
	Balance float64        `json:"balance"`
	Wallet  []WalletCrypto `json:"wallet"`
}

type WalletCrypto struct {
	Id     crypto.CryptoID `json:"id"`
	Name   string          `json:"name"`
	Amount float64         `json:"amount"`
}

func NewUser(name string) *User {
	user := User{
		Name:    name,
		Balance: DEFAULT_BALANCE,
		Wallet:  []WalletCrypto{},
	}

	return &user
}

func (user *User) Profile() {
	fmt.Println("Name:", user.Name)
	fmt.Println("Balance:", user.Balance)
	fmt.Println("Wallet:")

	for _, crypto := range user.Wallet {
		fmt.Println("Id:", crypto.Id)
		fmt.Println("Name:", crypto.Name)
		fmt.Println("Amount:", crypto.Amount)
		fmt.Println()
	}
}

type UserRepository interface {
	Get() *User
	Save(*User)
}
