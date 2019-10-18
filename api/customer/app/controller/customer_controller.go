package controller

import (
	"fmt"
	"net/http"

	"github.com/cbellee/goShop-customerService/app/helper/strkit"
	"github.com/cbellee/goShop-customerService/app/repository"
	"github.com/labstack/echo"
)

// CustomerController handle input related to Customer
type CustomerController interface {
	CRUD
}

type customerController struct {
	customerRepository repository.CustomerRepository
}

// NewCustomerController return new instance of customer controller
func NewCustomerController(customerRepository repository.CustomerRepository) CustomerController {
	return &customerController{
		customerRepository: customerRepository,
	}
}

func (c *customerController) Create(ctx echo.Context) (err error) {
	var customer repository.Customer

	err = ctx.Bind(&customer)
	if err != nil {
		return err
	}

	err = customer.Validate()
	if err != nil {
		return invalidMessage(ctx, err)
	}

	result, err := c.customerRepository.Insert(customer)
	if err != nil {
		return err
	}

	return insertSuccess(ctx, result)

}

func (c *customerController) List(ctx echo.Context) error {
	customers, err := c.customerRepository.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, customers)
}

func (c *customerController) Get(ctx echo.Context) error {
	id, err := strkit.ToInt64(ctx.Param("id"))
	if err != nil {
		return invalidID(ctx, err)
	}

	customer, err := c.customerRepository.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, customer)
}

func (c *customerController) Delete(ctx echo.Context) error {
	id, err := strkit.ToInt64(ctx.Param("id"))
	if err != nil {
		return invalidID(ctx, err)
	}

	err = c.customerRepository.Delete(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Delete #%d done\n", id)})
}

func (c *customerController) Update(ctx echo.Context) (err error) {
	var customer repository.Customer

	err = ctx.Bind(&customer)
	if err != nil {
		return err
	}

	if customer.ID <= 0 {
		return invalidID(ctx, err)
	}

	err = customer.Validate()
	if err != nil {
		return invalidMessage(ctx, err)
	}

	result, err := c.customerRepository.Update(customer, customer.ID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("updated %d items\n", result)})
}

func (c *customerController) BeforeActionFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		if c.customerRepository == nil {
			return fmt.Errorf("CustomerRepository is missing\n")
		}
		return next(e)
	}
}

func (c *customerController) RegisterTo(entity string, e *echo.Echo) {
	e.GET(fmt.Sprintf("/%s", entity), c.List, c.BeforeActionFunc)
	e.POST(fmt.Sprintf("/%s", entity), c.Create, c.BeforeActionFunc)
	e.GET(fmt.Sprintf("/%s/:id", entity), c.Get, c.BeforeActionFunc)
	e.PUT(fmt.Sprintf("/%s", entity), c.Update, c.BeforeActionFunc)
	e.DELETE(fmt.Sprintf("/%s/:id", entity), c.Delete, c.BeforeActionFunc)
}
