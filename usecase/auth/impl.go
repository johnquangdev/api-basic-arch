package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	rInterfaces "github.com/GunNguyen/api-basic-arch/repository/interfaces"
	rModel "github.com/GunNguyen/api-basic-arch/repository/model"
	uInterfaces "github.com/GunNguyen/api-basic-arch/usecase/interfaces"
	"github.com/GunNguyen/api-basic-arch/usecase/model"
	"github.com/GunNguyen/api-basic-arch/utils/hash"
	"github.com/GunNguyen/api-basic-arch/utils/token"
	"gorm.io/gorm"
)

type Auth struct {
	repository    rInterfaces.Repository
	tokenTimeLife time.Duration
	secretKey     string
}

func NewAuthUseCase(repo rInterfaces.Repository, tl time.Duration, sr string) uInterfaces.Auth {
	return &Auth{
		repository:    repo,
		tokenTimeLife: tl,
		secretKey:     sr,
	}
}

func (a Auth) CreateUser(ctx context.Context, req model.CreateUser) error {
	//ktr acccout
	if _, err := a.repository.Auth().CheckGmail(ctx, req.Gmail); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else {
		return fmt.Errorf("gmail existed")
	}
	// hash password
	hashPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return err
	}
	//create user
	if err := a.repository.Auth().CreateUser(ctx, rModel.User{
		Id:       req.Id,
		Gmail:    req.Gmail,
		Name:     req.Name,
		Password: hashPassword,
		Age:      req.Age,
	}); err != nil {
		return err
	}
	return nil
}

func (a Auth) LoginUser(ctx context.Context, req model.Login) (*model.LoginReply, error) {
	//check gmail
	user, err := a.repository.Auth().CheckGmail(ctx, req.Gmail)
	if err != nil {
		return nil, fmt.Errorf("wrong gmail")
	}
	//check password
	if !hash.CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("wrong password")
	}

	//generate access token
	accessToken, err := token.GenerateJwt(ctx, a.tokenTimeLife, a.secretKey, user.Name, user.Gmail)
	if err != nil {
		fmt.Println("generate accesstoken wrong")
		return nil, err
	}

	//generate refresh token
	refreshToken, err := token.GenerateJwt(ctx, a.tokenTimeLife+time.Minute, a.secretKey, user.Name, user.Gmail)
	if err != nil {
		fmt.Println("generate refreshtoken wrong")
		return nil, err
	}

	return &model.LoginReply{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
