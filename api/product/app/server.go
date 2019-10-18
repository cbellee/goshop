package app

import (
	"github.com/cbellee/goShop-productService/app/controller"
	"github.com/labstack/echo"
)

type server struct {
	*echo.Echo
	productController controller.ProductController
}

func newServer(productController controller.ProductController) *server {
	s := &server{
		Echo:              echo.New(),
		productController: productController,
	}

	initMiddlewares(s)
	initRoutes(s)

	return s
}

func (s *server) CRUD(entity string, crud controller.CRUD) {
	crud.RegisterTo(entity, s.Echo)
}
