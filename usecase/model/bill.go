package model

type Category struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type Product struct {
	ProductId       int    `json:"product_id"`
	ProductName     string `json:"product_name"`
	CategoryId      int    `json:"category_id"`
	Price           int    `json:"price"`
	Description     string `json:"description"`
	ImageUrl        string `json:"image_url"`
	QuantityInStock int    `json:"quantidy_in_stock"`
}
