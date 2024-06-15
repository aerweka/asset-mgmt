package main

import (
	"fmt"

	"asset-management.com/config"
	"asset-management.com/internal/server"
	"asset-management.com/pkg/cloudinary"
	"asset-management.com/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	cfg, err := config.Init()
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	db, err := database.ConnectDB(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	err = database.AutoMigrate(db)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	cloudinaryApp, err := cloudinary.InitInstance(*cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	server.NewHandler(app, db, cloudinaryApp)

	app.Listen(":3000")
}
