package user

import (
	"log"

	"github.com/goldenkingstyle/demo-crypto/internal/api"
	"github.com/goldenkingstyle/demo-crypto/internal/crypto"
)

type API interface {
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) CreateUser(name string) *User {
	user := NewUser(name)

	s.repo.Save(user)

	return user
}

func (s *UserService) Buy(id crypto.CryptoID, usd float64, api *api.API) {
	user := s.repo.Get()

	cryptoList, err := api.GetPrice()
	if err != nil {
		log.Fatal(err)
	}

	var crypto crypto.Crypto
	for _, cryptoListItem := range cryptoList {
		if cryptoListItem.ID == id {
			crypto = cryptoListItem
			break
		}
	}

	if crypto.ID == 0 {
		log.Fatal("Unknown id")
	}

	amount := usd / crypto.Price

	if usd > user.Balance {
		log.Fatal("Not enough money")
	}

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

	s.repo.Save(user)
}

func (s *UserService) Sell(id crypto.CryptoID, usd float64, api *api.API) {
	user := s.repo.Get()

	cryptoList, err := api.GetPrice()
	if err != nil {
		log.Fatal(err)
	}

	var crypto crypto.Crypto
	for _, cryptoListItem := range cryptoList {
		if cryptoListItem.ID == id {
			crypto = cryptoListItem
			break
		}
	}

	if crypto.ID == 0 {
		log.Fatal("Unknown id")
	}

	amount := usd / crypto.Price

	index := -1
	for i, walletCrypto := range user.Wallet {
		if walletCrypto.Id == id {
			index = i
		}
	}

	if index == -1 {
		log.Fatal("You don't have this crypto")
	}

	if user.Wallet[index].Amount < amount {
		log.Fatal("Not enough crypto")
	}

	user.Wallet[index].Amount -= amount
	user.Balance += usd

	s.repo.Save(user)
}

func (s *UserService) Profile() {
	user := s.repo.Get()
	user.Profile()
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo}
}
