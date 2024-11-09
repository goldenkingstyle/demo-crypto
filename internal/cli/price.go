package cli

import (
	"fmt"
	"log"

	"github.com/goldenkingstyle/demo-crypto/internal/api"
)

func Price(api *api.API) {
	cryptoList, err := api.GetPrice()
	if err != nil {
		log.Fatal(err)
	}

	for _, crypto := range cryptoList {
		fmt.Printf("%s: %f$\n", crypto.Name, crypto.Quote.USD.Price)
	}
}
