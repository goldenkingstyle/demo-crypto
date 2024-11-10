package crypto

type CryptoID int

type Crypto struct {
	ID    CryptoID `json:"id"`
	Name  string   `json:"name"`
	Price string   `json:"price"`
}
