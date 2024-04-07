package main

import (
	"fmt"

	"asset-management.com/config"
	"asset-management.com/database"
	"asset-management.com/router"
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

	router.NewRoute(app)

	app.Listen(":3000")
}
