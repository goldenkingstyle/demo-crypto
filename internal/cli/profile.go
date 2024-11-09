package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/goldenkingstyle/demo-crypto/internal/entity"
)

func Profile() {
	userJson, err := os.ReadFile("./storage/storage.json")
	if err != nil {
		log.Fatal(err)
	}

	var user entity.User
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Profile name:", user.Name)
	fmt.Println("Profile crypto wallet:")

}
