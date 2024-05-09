package main

import (
	"log"

	"github.com/GunNguyen/api-basic-arch/config"
	"github.com/GunNguyen/api-basic-arch/delivery"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	s := echo.New()
	//load config
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config, error: %v", err)
	}

	//custom logger
	// s.Use(middleware.Logger())
	// custom middleware
	s.Use(middleware.Recover())

	// register router
	if err := delivery.NewDelivery(s, *config); err != nil {
		log.Fatalf("failed to new delivery, error: %v", err)
	}
	//run server
	if err := s.Start(":8888"); err != nil {
		log.Fatal(err)
	}
}
