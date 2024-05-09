package auth

import (
	"context"
	"net/http"

	model "github.com/GunNguyen/api-basic-arch/model/response"
	modelr "github.com/GunNguyen/api-basic-arch/model/resquest"
	"github.com/GunNguyen/api-basic-arch/usecase/interfaces"
	uModel "github.com/GunNguyen/api-basic-arch/usecase/model"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	usecase interfaces.Usercase
}

func RegisterAuthRouter(g *echo.Group, useCase interfaces.Usercase) {
	a := Auth{
		usecase: useCase,
	}
	g.POST("/sign-up", a.CreateUser)
	g.POST("/sign-in", a.LoginUser)

	g.POST("/category", a.CreateCategory)
	g.POST("/product", a.CreateProduct)
}

func (a Auth) CreateUser(e echo.Context) error {
	var (
		req modelr.CreateUserResquest
	)
	//kiem tra du lieu nhan vao
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	//kiem tra dieu kien du lieu
	if err := req.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	//create
	if err := a.usecase.Auth().CreateUser(context.Background(), uModel.CreateUser{
		Gmail:    req.Gmail,
		Name:     req.Name,
		Age:      req.Age,
		Password: req.Password,
	}); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "ok",
	})
}

func (a Auth) LoginUser(e echo.Context) error {
	var (
		req modelr.LoginRequest
	)
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	if err := req.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	reply, err := a.usecase.Auth().LoginUser(context.Background(), uModel.Login{
		Gmail:    req.Gmail,
		Password: req.Password,
	})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, model.LoginReply{
		AccessToken:  reply.AccessToken,
		RefreshToken: reply.RefreshToken,
	})
}

func (a Auth) CreateCategory(e echo.Context) error {
	var (
		req modelr.Category
	)
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"massege": err.Error(),
		})
	}
	if err := a.usecase.Product().CreateCategory(context.Background(), uModel.Category{
		CategoryName: req.CategoryName,
	}); err != nil {
		return err
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"massege": req.CategoryName,
	})
}

func (a Auth) CreateProduct(e echo.Context) error {
	var (
		req modelr.Product
	)
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"massege": err.Error(),
		})
	}

	if err := a.usecase.Product().CreateProduct(context.Background(), uModel.Product{
		ProductName:     req.ProductName,
		Price:           req.Price,
		QuantityInStock: req.QuantityInStock,
		Description:     req.Description,
		ImageUrl:        req.ImageUrl,
		CategoryId:      req.CategoryId,
	}); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "192209",
			"massege": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"massege": "Ok done",
	})
}
