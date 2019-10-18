package repository

import (
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

// Address model
type Address struct {
	IsShippingAddress bool   `json:"isshippingaddress" validate:"required"`
	StreetAddress     string `json:"streetaddress" validate:"required"`
	Suburb            string `json:"suburb" validate:"required"`
	PostCode          int64  `json:"postcode" validate:"required"`
	City              string `json:"city" validate:"required"`
	State             string `json:"state" validate:"required"`
	Country           string `json:"country" validate:"required"`
	CountyCode        string `json:"countrycode" validate:"required"`
}

// PhoneNumber model
type PhoneNumber struct {
	Home   string `json:"home"`
	Mobile string `json:"mobile"`
	Work   string `json:"work"`
}

// Customer model
type Customer struct {
	ID         int64       `json:"id" validate:"required"`
	Email      string      `json:"email" validate:"required"`
	FirstName  string      `json:"firstname" validate:"required"`
	LastName   string      `json:"lastname" validate:"required"`
	MiddleName string      `json:"middlename" validate:"required"`
	Title      string      `json:"title" validate:"required"`
	Phone      PhoneNumber `json:"phone"`
	Addresses  []Address   `json:"addresses" validate:"required"`
	UpdatedAt  time.Time   `json:"updatedat,omit_empty"`
	CreatedAt  time.Time   `json:"createdat,omit_empty"`
}

// Validate customer
func (p *Customer) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
