package interfaces

import (
	"context"

	"github.com/GunNguyen/api-basic-arch/repository/model"
)

type Auth interface {
	CreateUser(context.Context, model.User) error
	CheckGmail(context.Context, string) (model.User, error)
}

type Product interface {
	CreateCategory(context.Context, model.Category) error
	CreateProduct(context.Context, model.Product) error
	GetCategoryById(context.Context, int) (model.Category, error)
}

type Repository interface {
	Auth() Auth
	Product() Product
}
