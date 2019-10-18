package repository

import (
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

// Price model
type Price struct {
	CD  float64 `json:"cd"`
	LP  float64 `json:"lp"`
	MP3 float64 `json:"mp3"`
}

// Product model
type Product struct {
	ID          int64     `json:"id"`
	Album       string    `json:"album" validate:"required"`
	Artist      string    `json:"artist" validate:"required"`
	ReleaseYear string    `json:"releaseyear" validate:"required"`
	Description string    `json:"description" validate:"required"`
	UpdatedAt   time.Time `json:"updatedat,omit_empty"`
	CreatedAt   time.Time `json:"createdat,omit_empty"`
	ImageURL    string    `json:"imageurl,omit_empty"`
	Price       Price     `json:"price" validate:"required"`
}

// Validate product
func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
