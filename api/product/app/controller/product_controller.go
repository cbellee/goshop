package controller

import (
	"fmt"
	"net/http"

	"github.com/cbellee/goShop-productService/app/helper/strkit"
	"github.com/cbellee/goShop-productService/app/repository"
	"github.com/labstack/echo"
)

// ProductController handle input related to Product
type ProductController interface {
	CRUD
}

type productController struct {
	productRepository repository.ProductRepository
}

// NewProductController return new instance of product controller
func NewProductController(productRepository repository.ProductRepository) ProductController {
	return &productController{
		productRepository: productRepository,
	}
}

func (c *productController) Create(ctx echo.Context) (err error) {
	var product repository.Product

	err = ctx.Bind(&product)
	if err != nil {
		return err
	}

	err = product.Validate()
	if err != nil {
		return invalidMessage(ctx, err)
	}

	result, err := c.productRepository.Insert(product)
	if err != nil {
		return err
	}

	return insertSuccess(ctx, result)

}

func (c *productController) List(ctx echo.Context) error {
	products, err := c.productRepository.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, products)
}

func (c *productController) Get(ctx echo.Context) error {
	id, err := strkit.ToInt64(ctx.Param("id"))
	if err != nil {
		return invalidID(ctx, err)
	}

	product, err := c.productRepository.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, product)
}

func (c *productController) Delete(ctx echo.Context) error {
	id, err := strkit.ToInt64(ctx.Param("id"))
	if err != nil {
		return invalidID(ctx, err)
	}

	err = c.productRepository.Delete(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Delete #%d done\n", id)})
}

func (c *productController) Update(ctx echo.Context) (err error) {
	var product repository.Product

	err = ctx.Bind(&product)
	if err != nil {
		return err
	}

	if product.ID <= 0 {
		return invalidID(ctx, err)
	}

	err = product.Validate()
	if err != nil {
		return invalidMessage(ctx, err)
	}

	result, err := c.productRepository.Update(product, product.ID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("updated %d items\n", result)})
}

func (c *productController) BeforeActionFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		if c.productRepository == nil {
			return fmt.Errorf("ProductRepository is missing\n")
		}
		return next(e)
	}
}

func (c *productController) RegisterTo(entity string, e *echo.Echo) {
	e.GET(fmt.Sprintf("/%s", entity), c.List, c.BeforeActionFunc)
	e.POST(fmt.Sprintf("/%s", entity), c.Create, c.BeforeActionFunc)
	e.GET(fmt.Sprintf("/%s/:id", entity), c.Get, c.BeforeActionFunc)
	e.PUT(fmt.Sprintf("/%s", entity), c.Update, c.BeforeActionFunc)
	e.DELETE(fmt.Sprintf("/%s/:id", entity), c.Delete, c.BeforeActionFunc)
}
