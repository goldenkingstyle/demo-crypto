package cli

import (
	"fmt"

	"github.com/goldenkingstyle/demo-crypto/internal/storage"
)

func Profile() {

	user := storage.ReadStorage()

	fmt.Println("Profile name:", user.Name)
	fmt.Println("Profile crypto wallet:")

}
