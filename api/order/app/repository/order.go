package repository

import (
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

// OrderItem model
type OrderItem struct {
	ID            int64   `json:"id" validate:"required"`
	Quantity      int64   `json:"quantity" validate:"required"`
	Format        string  `json:"format" validate:"required"`
	UnitPrice     float64 `json:"unitprice" validate:"required"`
	IsBackOrdered bool    `json:"isbackordered"`
}

// Order model
type Order struct {
	ID         int64       `json:"id"`
	CustomerID int64       `json:"customerid" validate:"required"`
	Items      []OrderItem `json:"items" validate:"required"`
	UpdatedAt  time.Time   `json:"updatedat,omit_empty"`
	CreatedAt  time.Time   `json:"createdat,omit_empty"`
	Total      float64     `json:"total" validate:"required"`
}

// Validate Order
func (p *Order) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
