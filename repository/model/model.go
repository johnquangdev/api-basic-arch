package model

import "time"

type User struct {
	Id        int       `json:"id"`
	Gmail     string    `json:"gmail"`
	Name      string    `json:"name"`
	Age       int       `json:"age,omitempty"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (User) TableName() string {
	return "user"
}
