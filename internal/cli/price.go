package cli

import (
	"flag"
	"log"
	"os"

	"github.com/goldenkingstyle/demo-crypto/internal/api"
	"github.com/goldenkingstyle/demo-crypto/internal/crypto"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Price(args []string, api *api.API) {
	priceSet := flag.NewFlagSet("price", flag.ExitOnError)

	count := priceSet.Int("count", 30, "count of crypto prices")

	priceSet.Parse(args)

	cryptoList, err := api.GetPrice()
	if err != nil {
		log.Fatal(err)
	}

	PrintCryptoList(count, cryptoList)
}

func PrintCryptoList(count *int, cryptoList []crypto.Crypto) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "ID", "Name", "Price (USD)"})

	for i, crypto := range cryptoList {
		if i == *count {
			break
		}
		t.AppendRow(table.Row{i + 1, crypto.ID, crypto.Name, crypto.Price})
	}

	t.SetStyle(table.StyleColoredBlackOnCyanWhite)
	t.Render()
}
