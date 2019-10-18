package app

import (
	"github.com/cbellee/goShop-customerService/app/controller"
	"github.com/labstack/echo"
)

type server struct {
	*echo.Echo
	customerController controller.CustomerController
}

func newServer(customerController controller.CustomerController) *server {
	s := &server{
		Echo:               echo.New(),
		customerController: customerController,
	}

	initMiddlewares(s)
	initRoutes(s)

	return s
}

func (s *server) CRUD(entity string, crud controller.CRUD) {
	crud.RegisterTo(entity, s.Echo)
}
