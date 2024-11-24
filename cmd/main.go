package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/goldenkingstyle/demo-crypto/internal/api"
	"github.com/goldenkingstyle/demo-crypto/internal/cli"
	"github.com/goldenkingstyle/demo-crypto/internal/config"
	"github.com/goldenkingstyle/demo-crypto/internal/storage"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load(ctx)
	if err != nil {
		log.Fatal(err)
	}

	user := storage.ReadStorage()

	api := api.NewAPI(*cfg)

	flag.Parse()

	switch flag.Arg(0) {
	case "init":
		cli.Init(os.Args[2:])
	case "profile":
		cli.Profile(user)
	case "set":
		cli.Set()
	case "price":
		cli.Price(os.Args[2:], api)
	case "buy":
		cli.Buy(os.Args[2:], user, api)
	default:
		fmt.Println("Unknown command")
	}
}
