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

// TODO: Renaming
type CryptoListResponse struct {
	CryptoList []CryptoResponse `json:"data"`
}

type CryptoResponse struct {
	ID    crypto.CryptoID `json:"id"`
	Name  string          `json:"name"`
	Quote struct {
		USD struct {
			Price float32 `json:"price"`
		} `json:"USD"`
	} `json:"quote"`
}

func NewAPI(cfg config.Config) *API {
	return &API{
		client:  &http.Client{},
		api_key: cfg.API_KEY,
	}
}

func (api *API) GetPrice() ([]CryptoResponse, error) {
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

	cryptoList := cryptoListResponse.CryptoList

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
