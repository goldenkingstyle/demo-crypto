package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/goldenkingstyle/demo-crypto/internal/api"
	"github.com/goldenkingstyle/demo-crypto/internal/config"
	"github.com/goldenkingstyle/demo-crypto/internal/crypto"
	"github.com/goldenkingstyle/demo-crypto/internal/user"
	"github.com/jedib0t/go-pretty/v6/table"
)

type CLI struct {
	command string
	args    []string
}

func NewCLI(command string, args []string) *CLI {
	return &CLI{command: command, args: args}
}

func (cli *CLI) Command(cfg *config.Config, api *api.API) {
	switch cli.command {
	case "init":
		cli.Init()
	case "profile":
		cli.Profile()
	case "set":
		cli.Set()
	case "price":
		cli.Price(api)
	case "buy":
		cli.Buy(api)
	case "sell":
		cli.Sell(api)
	default:
		fmt.Println("Unknown command")
	}
}

func (cli *CLI) Init() {
	initSet := flag.NewFlagSet("init", flag.ExitOnError)
	name := initSet.String("name", "user", "Profile name")

	initSet.Parse(cli.args)

	storagePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	storagePath += "/crypto-storage"

	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		err = os.Mkdir(storagePath, 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

	_ = user.NewUser(*name)
}

func (cli *CLI) Profile() {
	user := user.GetUser()

	user.Profile()
}

func (cli *CLI) Set() {}

func (cli *CLI) Price(api *api.API) {
	priceSet := flag.NewFlagSet("price", flag.ExitOnError)

	count := priceSet.Int("count", 30, "count of crypto prices")

	priceSet.Parse(cli.args)

	cryptoList, err := api.GetPrice()
	if err != nil {
		log.Fatal(err)
	}

	PrintCryptoList(count, cryptoList)
}

func (cli *CLI) Buy(api *api.API) {
	user := user.GetUser()

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

	user.Buy(crypto.CryptoID(*id), *usd, api)
}

func (cli *CLI) Sell(api *api.API) {
	user := user.GetUser()

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

	user.Sell(crypto.CryptoID(*id), *usd, api)
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
