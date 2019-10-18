package app

import (
	"github.com/cbellee/goShop-orderService/app/controller"
	"github.com/labstack/echo"
)

type server struct {
	*echo.Echo
	orderController controller.OrderController
}

func newServer(orderController controller.OrderController) *server {
	s := &server{
		Echo:               echo.New(),
		orderController: orderController,
	}

	initMiddlewares(s)
	initRoutes(s)

	return s
}

func (s *server) CRUD(entity string, crud controller.CRUD) {
	crud.RegisterTo(entity, s.Echo)
}
