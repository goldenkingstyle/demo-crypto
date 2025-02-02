package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/goldenkingstyle/demo-crypto/internal/api"
	"github.com/goldenkingstyle/demo-crypto/internal/cli"
	"github.com/goldenkingstyle/demo-crypto/internal/config"
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
	args := os.Args[2:]
	cli := cli.NewCLI(command, args)

	app := newApp(cfg, cli, api)

	app.run()
}

type App struct {
	cfg *config.Config
	cli *cli.CLI
	api *api.API
}

func newApp(cfg *config.Config, cli *cli.CLI, api *api.API) *App {
	return &App{cfg, cli, api}
}

func (app *App) run() {
	app.cli.Command(app.cfg, app.api)
}
