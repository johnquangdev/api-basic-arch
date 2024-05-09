package model

import (
	"github.com/go-playground/validator/v10"
)

type CreateUserResquest struct {
	Gmail    string `json:"gmail" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age"`
	Password string `json:"password" validate:"min=8,containsany=!@#?"`
}

func (c *CreateUserResquest) Validate() error {
	var validate = validator.New()
	return validate.Struct(c)
}

type LoginRequest struct {
	Gmail    string
	Password string
}

func (l *LoginRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(l)
}

type Category struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type Product struct {
	ProductName     string `json:"product_name"`
	CategoryId      int    `json:"category_id"`
	QuantityInStock int    `json:"quantity_in_stock"`
	Price           int    `json:"price"`
	Description     string `json:"description"`
	ImageUrl        string `json:"image_url"`
}
