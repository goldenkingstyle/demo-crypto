package crypto

type CryptoID int

type Crypto struct {
	ID    CryptoID `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
}
