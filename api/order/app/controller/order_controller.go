package controller

import (
	"fmt"
	"net/http"

	"github.com/cbellee/goShop-orderService/app/helper/strkit"
	"github.com/cbellee/goShop-orderService/app/repository"
	"github.com/labstack/echo"
)

// OrderController handle input related to Order
type OrderController interface {
	CRUD
}

type orderController struct {
	orderRepository repository.OrderRepository
}

// NewOrderController return new instance of order controller
func NewOrderController(orderRepository repository.OrderRepository) OrderController {
	return &orderController{
		orderRepository: orderRepository,
	}
}

func (c *orderController) Create(ctx echo.Context) (err error) {
	var order repository.Order

	err = ctx.Bind(&order)
	if err != nil {
		return err
	}

	err = order.Validate()
	if err != nil {
		return invalidMessage(ctx, err)
	}

	result, err := c.orderRepository.Insert(order)
	if err != nil {
		return err
	}

	return insertSuccess(ctx, result)

}

func (c *orderController) List(ctx echo.Context) error {
	orders, err := c.orderRepository.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, orders)
}

func (c *orderController) Get(ctx echo.Context) error {
	id, err := strkit.ToInt64(ctx.Param("id"))
	if err != nil {
		return invalidID(ctx, err)
	}

	order, err := c.orderRepository.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, order)
}

func (c *orderController) Delete(ctx echo.Context) error {
	id, err := strkit.ToInt64(ctx.Param("id"))
	if err != nil {
		return invalidID(ctx, err)
	}

	err = c.orderRepository.Delete(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Delete #%d done\n", id)})
}

func (c *orderController) Update(ctx echo.Context) (err error) {
	var order repository.Order

	err = ctx.Bind(&order)
	if err != nil {
		return err
	}

	if order.ID <= 0 {
		return invalidID(ctx, err)
	}

	err = order.Validate()
	if err != nil {
		return invalidMessage(ctx, err)
	}

	result, err := c.orderRepository.Update(order, order.ID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("updated %d items\n", result)})
}

func (c *orderController) BeforeActionFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		if c.orderRepository == nil {
			return fmt.Errorf("OrderRepository is missing\n")
		}
		return next(e)
	}
}

func (c *orderController) RegisterTo(entity string, e *echo.Echo) {
	e.GET(fmt.Sprintf("/%s", entity), c.List, c.BeforeActionFunc)
	e.POST(fmt.Sprintf("/%s", entity), c.Create, c.BeforeActionFunc)
	e.GET(fmt.Sprintf("/%s/:id", entity), c.Get, c.BeforeActionFunc)
	e.PUT(fmt.Sprintf("/%s", entity), c.Update, c.BeforeActionFunc)
	e.DELETE(fmt.Sprintf("/%s/:id", entity), c.Delete, c.BeforeActionFunc)
}
