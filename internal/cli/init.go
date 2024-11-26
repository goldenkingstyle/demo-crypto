package cli

import (
	"flag"

	"github.com/goldenkingstyle/demo-crypto/internal/user"
)

func Init(args []string) {
	initSet := flag.NewFlagSet("init", flag.ExitOnError)
	name := initSet.String("name", "user", "Profile name")

	initSet.Parse(args)

	_ = user.NewUser(*name)
}
