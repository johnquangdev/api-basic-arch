package sell

import (
	"context"

	"github.com/GunNguyen/api-basic-arch/repository/interfaces"
	"github.com/GunNguyen/api-basic-arch/repository/model"
	"gorm.io/gorm"
)

type Auth struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.Product {
	return &Auth{
		Db: db,
	}
}

func (a Auth) CreateCategory(ctx context.Context, req model.Category) error {
	return a.Db.Create(&req).Error
}

func (a Auth) CreateProduct(ctx context.Context, req model.Product) error {
	return a.Db.Create(&req).Error
}

func (a Auth) GetCategoryById(ctx context.Context, id int) (model.Category, error) {
	var (
		category model.Category
	)
	if err := a.Db.Where("category_id=?", id).Take(&category).Error; err != nil {
		return model.Category{}, err
	}
	return category, nil
}
