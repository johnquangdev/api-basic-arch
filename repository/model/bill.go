package model

type Product struct {
	ProductId       int    `json:"product_id"`
	ProductName     string `json:"product_name"`
	CategoryID      int    `json:"category_id"`
	Price           int    `json:"price"`
	Description     string `json:"description"`
	ImageUrl        string `json:"image_url"`
	QuantityInStock int    `json:"quantidy_in_stock"`
}

func (Product) TableName() string {
	return "products"
}

type Category struct {
	CategoryId    int    `json:"category_id"`
	Category_name string `json:"category_name"`
}

func (Category) TableName() string {
	return "categories"
}
