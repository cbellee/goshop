package main

import (
	"log"
	"os"

	"github.com/cbellee/goShop-customerService/app"
	"github.com/cbellee/goShop-customerService/config"
	"github.com/urfave/cli"
)

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = config.App.Name
	cliApp.Usage = config.App.Usage
	cliApp.Version = config.App.Version
	cliApp.Commands = app.Commands()

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
