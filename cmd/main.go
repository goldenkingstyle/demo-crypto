package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/goldenkingstyle/demo-crypto/internal/api"
	"github.com/goldenkingstyle/demo-crypto/internal/cli"
	"github.com/goldenkingstyle/demo-crypto/internal/config"
	"github.com/goldenkingstyle/demo-crypto/internal/user"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load(ctx)
	if err != nil {
		log.Fatal(err)
	}

	api := api.NewAPI(*cfg)

	flag.Parse()
	command := flag.Arg(0)

	var args []string
	if len(os.Args) >= 2 {
		args = os.Args[2:]
	}

	repo := user.NewJsonUserRepository(cfg.Filepath)

	userService := user.NewUserService(repo)

	cli := cli.NewCLI(command, args, api, userService)

	cli.Run()
}
