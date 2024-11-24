package user

import (
	"fmt"
	"log"

	"github.com/goldenkingstyle/demo-crypto/internal/api"
	"github.com/goldenkingstyle/demo-crypto/internal/crypto"
)

const DEFAULT_BALANCE float64 = 10000

type User struct {
	Name    string       `json:"name"`
	Balance float64      `json:"balance"`
	Wallet  []WalletItem `json:"wallet"`
}

type WalletItem struct {
	id     crypto.CryptoID
	amount float64
}

func NewUser(name string) *User {
	return &User{
		Name:    name,
		Balance: DEFAULT_BALANCE,
		Wallet:  []WalletItem{},
	}
}

func (user *User) Profile() {
	fmt.Println("Profile name:", user.Name)
	fmt.Println("Profile crypto wallet:")
}

func (user *User) Buy(id crypto.CryptoID, usd float64, api *api.API) {
	cryptoList, err := api.GetPrice()
	if err != nil {
		log.Fatal(err)
	}

	var price float64
	for _, crypto := range cryptoList {
		if crypto.ID == id {
			price = crypto.Quote.USD.Price
		}
	}

	if price == 0 {
		log.Fatal("Unknown id")
	}

	amount := price / usd

	user.Balance -= usd

	for _, crypto := range user.Wallet {
		if crypto.id == id {
			crypto.amount += amount
		}
	}

	user.Update()
}

func (user *User) Update() {
}
