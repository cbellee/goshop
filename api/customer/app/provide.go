package app

import (
	"github.com/cbellee/goShop-customerService/app/controller"
	"github.com/cbellee/goShop-customerService/app/repository"
	"github.com/cbellee/goShop-customerService/config"
	"github.com/cbellee/goShop-customerService/db"
	"github.com/urfave/cli"
	"go.uber.org/dig"
)

func container() *dig.Container {
	container := dig.New()
	container.Provide(newServer)
	container.Provide(config.LoadConfig)
	container.Provide(db.Connect)
	container.Provide(controller.NewCustomerController)
	container.Provide(repository.NewCustomerRepository)

	return container
}

func triggerAction(invokeFunc interface{}) interface{} {
	return func(ctx *cli.Context) error {
		container := container()
		container.Provide(ctx.Args)
		return container.Invoke(invokeFunc)
	}
}
