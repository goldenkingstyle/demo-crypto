package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/goldenkingstyle/demo-crypto/internal/api"
	"github.com/goldenkingstyle/demo-crypto/internal/crypto"
	"github.com/goldenkingstyle/demo-crypto/internal/user"
	"github.com/jedib0t/go-pretty/v6/table"
)

type CLI struct {
	command     string
	args        []string
	api         *api.API
	userService *user.UserService
}

func NewCLI(command string, args []string, api *api.API, userService *user.UserService) *CLI {
	return &CLI{command, args, api, userService}
}

func (cli *CLI) Run() {
	switch cli.command {
	case "init":
		cli.Init()
	case "profile":
		cli.Profile()
	case "set":
		cli.Set()
	case "price":
		cli.Price(cli.api)
	case "buy":
		cli.Buy(cli.api)
	case "sell":
		cli.Sell(cli.api)
	default:
		fmt.Println("Unknown command")
	}
}

func (cli *CLI) Init() {
	initSet := flag.NewFlagSet("init", flag.ExitOnError)
	name := initSet.String("name", "user", "Profile name")

	initSet.Parse(cli.args)

	cli.userService.CreateUser(*name)
}

func (cli *CLI) Profile() {
	cli.userService.Profile()
}

func (cli *CLI) Set() {}

func (cli *CLI) Price(api *api.API) {
	priceSet := flag.NewFlagSet("price", flag.ExitOnError)

	count := priceSet.Int("count", 30, "count of crypto prices")

	priceSet.Parse(cli.args)

	cryptoList, err := cli.api.GetPrice()
	if err != nil {
		log.Fatal(err)
	}

	PrintCryptoList(count, cryptoList)
}

func (cli *CLI) Buy(api *api.API) {
	buySet := flag.NewFlagSet("buy", flag.ExitOnError)

	id := buySet.Int("id", 0, "crypto id for buying")
	usd := buySet.Float64("usd", 0, "amount of usd for buying")

	buySet.Parse(cli.args)

	if *id < 1 {
		log.Fatal("Incorrect id")
	}

	if *usd <= 0 {
		log.Fatal("Incorrect amount of usd")
	}

	cli.userService.Buy(crypto.CryptoID(*id), *usd, api)
}

func (cli *CLI) Sell(api *api.API) {
	sellSet := flag.NewFlagSet("sell", flag.ExitOnError)

	id := sellSet.Int("id", 0, "crypto id for selling")
	usd := sellSet.Float64("usd", 0, "amount of usd for selling")

	sellSet.Parse(cli.args)

	if *id < 1 {
		log.Fatal("Incorrect id")
	}

	if *usd <= 0 {
		log.Fatal("Incorrect amount of usd")
	}

	cli.userService.Sell(crypto.CryptoID(*id), *usd, api)
}

func PrintCryptoList(count *int, cryptoList []crypto.Crypto) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "ID", "Name", "Price (USD)"})

	for i, crypto := range cryptoList {
		if i == *count {
			break
		}
		t.AppendRow(table.Row{i + 1, crypto.ID, crypto.Name, crypto.Price})
	}

	t.SetStyle(table.StyleColoredBlackOnCyanWhite)
	t.Render()
}
