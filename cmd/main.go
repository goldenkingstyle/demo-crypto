package main

import (
	"flag"
	"fmt"

	"github.com/goldenkingstyle/demo-crypto/internal/cli"
)

func main() {

	flag.Parse()

	switch flag.Arg(0) {
	case "init":
		cli.Init()
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
