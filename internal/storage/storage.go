package storage

import (
	"log"
	"os"
)

func CreateStorage(userJson []byte) {
	err := os.WriteFile("./storage/storage.json", userJson, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
