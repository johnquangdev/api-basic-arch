package delivery

import (
	"time"

	"github.com/GunNguyen/api-basic-arch/config"
	"github.com/GunNguyen/api-basic-arch/delivery/auth"
	"github.com/GunNguyen/api-basic-arch/repository"
	"github.com/GunNguyen/api-basic-arch/usecase"
	"github.com/labstack/echo/v4"
)

func NewDelivery(e *echo.Echo, cfg config.Config) error {
	repo, err := repository.NewRepository(repository.Database{
		Host:     cfg.DB.Host,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
		Port:     cfg.DB.Port,
	})
	if err != nil {
		return err
	}

	usecase := usecase.NewUseCase(repo, time.Duration(cfg.JWT.TokenTimeLife), cfg.JWT.SecretKey)

	auth.RegisterAuthRouter(e.Group("/users"), usecase)

	return nil
}
