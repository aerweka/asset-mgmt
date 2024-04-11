package main

import (
	"fmt"

	"asset-management.com/config"
	"asset-management.com/database"
	"asset-management.com/internal/server"
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

	server.NewHandler(app, db)

	app.Listen(":3000")
}
