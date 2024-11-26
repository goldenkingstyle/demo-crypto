package user

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

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
	user := User{
		Name:    name,
		Balance: DEFAULT_BALANCE,
		Wallet:  []WalletItem{},
	}

	user.Update()

	return &user
}

func GetUser() *User {
	userJson, err := os.ReadFile("./storage/storage.json")
	if err != nil {
		log.Fatal(err)
	}

	var user User
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		log.Fatal(err)
	}

	return &user
}

func (user *User) Profile() {
	fmt.Println("Name:", user.Name)
	fmt.Println("Balance:", user.Balance)
	fmt.Println("Wallet:")
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

	if usd > user.Balance {
		log.Fatal("Not enough money")
	}

	amount := price / usd
	user.Balance -= usd

	index := -1
	for i, crypto := range user.Wallet {
		if crypto.id == id {
			index = i
		}
	}

	if index == -1 {
		user.Wallet = append(user.Wallet, WalletItem{
			id:     crypto.CryptoID(len(user.Wallet)),
			amount: amount,
		})
	} else {
		user.Wallet[index].amount += amount
	}

	user.Update()
}

func (user *User) Update() {
	userJson, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("./storage/storage.json", userJson, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
