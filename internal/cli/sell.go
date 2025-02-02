package cli

import (
	"flag"
	"log"

	"github.com/goldenkingstyle/demo-crypto/internal/api"
	"github.com/goldenkingstyle/demo-crypto/internal/crypto"
	"github.com/goldenkingstyle/demo-crypto/internal/user"
)

func Sell(args []string, api *api.API) {
	user := user.GetUser()

	sellSet := flag.NewFlagSet("sell", flag.ExitOnError)

	id := sellSet.Int("id", 0, "crypto id for selling")
	usd := sellSet.Float64("usd", 0, "amount of usd for selling")

	sellSet.Parse(args)

	if *id < 1 {
		log.Fatal("Incorrect id")
	}

	if *usd <= 0 {
		log.Fatal("Incorrect amount of usd")
	}

	user.Sell(crypto.CryptoID(*id), *usd, api)
}
