// routes/routes.go

package routes

import (
	"go-layer-pgr-ex/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	cfg := config.LoadConfig()
	prefix := cfg.Prefix
	// CORS settings
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	// Routes
	app.Get(prefix+"/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, from Bearber!")
	})
	// Setup routes
	SetupTaskRoutes(app, db, prefix)

	//Error handling
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"message": "Not Found"})
	})
}
