package cli

import (
	"flag"
	"log"

	"github.com/goldenkingstyle/demo-crypto/internal/api"
	"github.com/goldenkingstyle/demo-crypto/internal/crypto"
	"github.com/goldenkingstyle/demo-crypto/internal/user"
)

func Buy(args []string, api *api.API) {
	user := user.GetUser()

	buySet := flag.NewFlagSet("buy", flag.ExitOnError)

	id := buySet.Int("id", 0, "crypto id for buying")
	usd := buySet.Float64("usd", 0, "amount of usd for buying")

	buySet.Parse(args)

	if *id < 1 {
		log.Fatal("Incorrect id")
	}

	if *usd <= 0 {
		log.Fatal("Incorrect amount of usd")
	}

	user.Buy(crypto.CryptoID(*id), *usd, api)
}
