package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/goldenkingstyle/demo-crypto/internal/cli"
	"github.com/goldenkingstyle/demo-crypto/internal/config"
)

func main() {

	ctx := context.Background()

	cfg, err := config.Load(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("API_KEY:", cfg.API_KEY)

	flag.Parse()

	switch flag.Arg(0) {
	case "init":
		cli.Init(os.Args[2:])
	case "profile":
		cli.Profile()
	case "set":
		cli.Set()
	case "price":
		cli.Price()
	default:
		fmt.Println("Unknown command")
	}

}
