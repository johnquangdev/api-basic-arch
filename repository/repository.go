package repository

import (
	"fmt"

	"github.com/GunNguyen/api-basic-arch/repository/auth"
	sell "github.com/GunNguyen/api-basic-arch/repository/auth/Sell"
	"github.com/GunNguyen/api-basic-arch/repository/interfaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

func (d *DB) Auth() interfaces.Auth {
	return auth.NewAuthRepository(d.Db)
}

func (d *DB) Product() interfaces.Product {
	return sell.NewProductRepository(d.Db)
}

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
type Database struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
}

func NewDatabaseConnection(d Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		d.Host,
		d.User,
		d.Password,
		d.Name,
		d.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewRepository(d Database) (interfaces.Repository, error) {
	db, err := NewDatabaseConnection(d)
	if err != nil {
		return nil, err
	}
	return &DB{
		Db: db,
	}, nil
}
