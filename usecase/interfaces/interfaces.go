package interfaces

import (
	"context"

	"github.com/GunNguyen/api-basic-arch/usecase/model"
)

type Auth interface {
	CreateUser(context.Context, model.CreateUser) error
	LoginUser(context.Context, model.Login) (*model.LoginReply, error)
}

type Product interface {
	CreateCategory(context.Context, model.Category) error
	CreateProduct(context.Context, model.Product) error
}

type Usercase interface {
	Auth() Auth
	Product() Product
}
