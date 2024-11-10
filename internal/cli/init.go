package cli

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/goldenkingstyle/demo-crypto/internal/storage"
	"github.com/goldenkingstyle/demo-crypto/internal/user"
)

func Init(args []string) {

	initSet := flag.NewFlagSet("init", flag.ExitOnError)
	name := initSet.String("name", "user", "Profile name")

	initSet.Parse(args)

	user := user.NewUser(*name)

	userJson, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	storage.CreateStorage(userJson)
}
