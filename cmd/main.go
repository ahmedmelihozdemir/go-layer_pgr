package main

import (
	"go-layer-pgr-ex/config"
	"go-layer-pgr-ex/models"
	"go-layer-pgr-ex/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.LoadConfig()
	config.DB.AutoMigrate(&models.User{}, &models.Task{})

	routes.SetupRoutes(app, config.DB)
	app.Listen(":3000")
}
