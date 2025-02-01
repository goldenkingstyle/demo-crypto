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

	for _, crypto := range user.Wallet {
		fmt.Println("Id:", crypto.Id)
		fmt.Println("Name:", crypto.Name)
		fmt.Println("Amount:", crypto.Amount)
		fmt.Println()
	}
}

func (user *User) Buy(id crypto.CryptoID, usd float64, api *api.API) {
	cryptoList, err := api.GetPrice()
	if err != nil {
		log.Fatal(err)
	}

	var crypto crypto.Crypto
	for _, cryptoListItem := range cryptoList {
		if cryptoListItem.ID == id {
			crypto.ID = cryptoListItem.ID
			crypto.Name = cryptoListItem.Name
			crypto.Price = cryptoListItem.Quote.USD.Price
			break
		}
	}

	if crypto.ID == 0 {
		log.Fatal("Unknown id")
	}

	if usd > user.Balance {
		log.Fatal("Not enough money")
	}

	amount := usd / crypto.Price
	user.Balance -= usd

	index := -1
	for i, walletCrypto := range user.Wallet {
		if walletCrypto.Id == id {
			index = i
		}
	}

	if index == -1 {
		user.Wallet = append(user.Wallet, WalletCrypto{
			Id:     crypto.ID,
			Name:   crypto.Name,
			Amount: amount,
		})
	} else {
		user.Wallet[index].Amount += amount
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
