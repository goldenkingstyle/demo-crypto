package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/goldenkingstyle/demo-crypto/internal/config"
	"github.com/goldenkingstyle/demo-crypto/internal/crypto"
)

type API struct {
	client  *http.Client
	api_key string
}

type CryptoListResponse struct {
	CryptoList []CryptoResponse `json:"data"`
}

type CryptoResponse struct {
	ID     crypto.CryptoID `json:"id"`
	Name   string          `json:"name"`
	Symbol string          `json:"symbol"`
	Quote  struct {
		USD struct {
			Price float64 `json:"price"`
		} `json:"USD"`
	} `json:"quote"`
}

func NewAPI(cfg config.Config) *API {
	return &API{
		client:  &http.Client{},
		api_key: cfg.API_KEY,
	}
}

func (api *API) GetPrice() ([]crypto.Crypto, error) {
	response, err := api.get("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var cryptoListResponse CryptoListResponse
	err = json.Unmarshal(data, &cryptoListResponse)
	if err != nil {
		return nil, err
	}

	var cryptoList []crypto.Crypto

	for _, cryptoResponse := range cryptoListResponse.CryptoList {
		cryptoList = append(cryptoList, crypto.Crypto{
			ID:    cryptoResponse.ID,
			Name:  cryptoResponse.Name,
			Price: cryptoResponse.Quote.USD.Price,
		})
	}

	return cryptoList, nil
}

func (api *API) get(url string, body io.Reader) (*http.Response, error) {
	return api.newRequest(http.MethodGet, url, body)
}

func (api *API) newRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", api.api_key)

	return api.client.Do(req)
}
