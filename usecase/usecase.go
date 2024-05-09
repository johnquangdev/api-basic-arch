package usecase

import (
	"time"

	rInterface "github.com/GunNguyen/api-basic-arch/repository/interfaces"
	"github.com/GunNguyen/api-basic-arch/usecase/auth"
	"github.com/GunNguyen/api-basic-arch/usecase/auth/sell"
	uInterface "github.com/GunNguyen/api-basic-arch/usecase/interfaces"
)

type UseCase struct {
	repository    rInterface.Repository
	tokenTimeLife time.Duration
	secretKey     string
}

func (u *UseCase) Auth() uInterface.Auth {
	return auth.NewAuthUseCase(u.repository, u.tokenTimeLife, u.secretKey)
}

func (u *UseCase) Product() uInterface.Product {
	return sell.NewProductUseCase(u.repository)
}

func NewUseCase(repo rInterface.Repository, tokenTimeLife time.Duration, secretKey string) uInterface.Usercase {
	return &UseCase{
		repository:    repo,
		tokenTimeLife: tokenTimeLife,
		secretKey:     secretKey,
	}
}
