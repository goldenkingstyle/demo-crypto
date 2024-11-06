package cli

import (
	"encoding/json"
	"log"

	"github.com/goldenkingstyle/demo-crypto/internal/entity"
	"github.com/goldenkingstyle/demo-crypto/internal/storage"
)

func Init() {

	name := "user"

	user := entity.NewUser(name)

	userJson, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	storage.CreateStorage(userJson)
}
