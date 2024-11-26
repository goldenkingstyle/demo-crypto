package cli

import (
	"github.com/goldenkingstyle/demo-crypto/internal/user"
)

func Profile() {
	user := user.GetUser()

	user.Profile()
}
