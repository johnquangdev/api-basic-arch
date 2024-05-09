package model

import "github.com/go-playground/validator/v10"

type CreateUserResponse struct {
	Gmail    string `json:"gmail" validate:"required"`
	UserName string `json:"username" validate:"required"`
	Age      int    `json:"age"`
	Password string `json:"-" validate:"min=8,containsany=!@#?"`
}

func (c *CreateUserResponse) Validate() error {
	var validate = validator.New()
	return validate.Struct(c)
}

type LoginReply struct {
	UserName     string `json:"username"`
	Age          int    `json:"age,omitempty"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (l *LoginReply) Validate() error {
	var validate = validator.New()
	return validate.Struct(l)
}
