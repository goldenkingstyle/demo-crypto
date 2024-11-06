package entity

type User struct {
	Name   string     `json:"name"`
	Wallet []CryptoID `json:"wallet"`
}

func NewUser(name string) *User {
	return &User{name, []CryptoID{}}
}
