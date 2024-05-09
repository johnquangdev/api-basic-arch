package sell

import (
	"context"
	"errors"
	"fmt"

	rInterface "github.com/GunNguyen/api-basic-arch/repository/interfaces"
	uModel "github.com/GunNguyen/api-basic-arch/repository/model"
	"github.com/GunNguyen/api-basic-arch/usecase/interfaces"
	"github.com/GunNguyen/api-basic-arch/usecase/model"
	"gorm.io/gorm"
)

type Auth struct {
	repository rInterface.Repository
}

func NewProductUseCase(repo rInterface.Repository) interfaces.Product {
	return &Auth{
		repository: repo,
	}
}

func (a Auth) CreateCategory(ctx context.Context, req model.Category) error {
	//check quyền
	//check token
	if err := a.repository.Product().CreateCategory(context.Background(), uModel.Category{
		Category_name: req.CategoryName,
	}); err != nil {
		return err
	}
	return nil
}

func (a Auth) CreateProduct(ctx context.Context, req model.Product) error {
	if _, err := a.repository.Product().GetCategoryById(ctx, req.CategoryId); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("category dose not exist")
		}
		return err
	}
	if err := a.repository.Product().CreateProduct(ctx, uModel.Product{
		// ProductId:   req.ProductId,
		ProductName:     req.ProductName,
		Price:           req.Price,
		Description:     req.Description,
		ImageUrl:        req.ImageUrl,
		QuantityInStock: req.QuantityInStock,
		CategoryID:      req.CategoryId,
	}); err != nil {
		return err
	}
	return nil
}
