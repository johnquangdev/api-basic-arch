package auth

import (
	"context"

	"github.com/GunNguyen/api-basic-arch/repository/interfaces"
	"github.com/GunNguyen/api-basic-arch/repository/model"
	"gorm.io/gorm"
)

type Auth struct {
	Db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) interfaces.Auth {
	return &Auth{
		Db: db,
	}
}

func (a Auth) CreateUser(ctx context.Context, req model.User) error {
	return a.Db.Create(&req).Error
}

func (a Auth) CheckGmail(ctx context.Context, gmail string) (model.User, error) {
	var user model.User

	if err := a.Db.Where("gmail=?", gmail).Take(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
