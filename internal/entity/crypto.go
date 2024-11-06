package entity

type CryptoID int

type Crypto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}
