package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cbellee/goShop-orderService/config"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func serve(s *server, conf config.Config) error {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	// gracefull shutdown
	go func() {
		<-gracefulStop
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		s.Shutdown(ctx)
	}()

	var listenAddress = conf.Address + ":" + conf.Port
	return s.Start(listenAddress)
}

func printConfigDetails(ctx *cli.Context) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type", "Required", "Default"})
	for _, detail := range config.Details() {
		table.Append([]string{detail.Name, detail.Type, detail.Required, detail.Default})
	}
	table.Render()
}
