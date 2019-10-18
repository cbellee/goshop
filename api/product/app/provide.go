package app

import (
	"github.com/cbellee/goShop-productService/app/controller"
	"github.com/cbellee/goShop-productService/app/repository"
	"github.com/cbellee/goShop-productService/config"
	"github.com/cbellee/goShop-productService/db"
	"github.com/urfave/cli"
	"go.uber.org/dig"
)

func container() *dig.Container {
	container := dig.New()
	container.Provide(newServer)
	container.Provide(config.LoadConfig)
	container.Provide(db.Connect)
	container.Provide(controller.NewProductController)
	container.Provide(repository.NewProductRepository)

	return container
}

func triggerAction(invokeFunc interface{}) interface{} {
	return func(ctx *cli.Context) error {
		container := container()
		container.Provide(ctx.Args)
		return container.Invoke(invokeFunc)
	}
}
