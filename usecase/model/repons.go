package model

import (
	"time"
)

type CreateUser struct {
	Id         int       `json:"-"`
	Gmail      string    `json:"gmail"`
	Name       string    `json:"name"`
	Age        int       `json:"age,omitempty"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"-"`
	UppdatedAt time.Time `json:"-"`
}

type Login struct {
	Gmail    string
	Password string
}

type LoginReply struct {
	Name         string `json:"name"`
	Age          int    `json:"age,omitempty"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
