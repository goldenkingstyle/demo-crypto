package cli

import "github.com/goldenkingstyle/demo-crypto/internal/storage"

func Init() {
	storage.CreateStorage()
}
