package storage

import (
	"encoding/json"
	"log"
	"os"

	"github.com/goldenkingstyle/demo-crypto/internal/user"
)

func CreateStorage(userJson []byte) {
	err := os.WriteFile("./storage/storage.json", userJson, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadStorage() user.User {
	userJson, err := os.ReadFile("./storage/storage.json")
	if err != nil {
		log.Fatal(err)
	}

	var user user.User
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		log.Fatal(err)
	}

	return user
}
