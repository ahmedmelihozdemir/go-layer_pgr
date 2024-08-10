package routes

import (
	"go-layer-pgr-ex/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupTaskRoutes(app *fiber.App, db *gorm.DB, prefix string) {
	app.Post(prefix+"/tasks", handlers.CreateTask)
	app.Get(prefix+"/tasks/:userID", handlers.GetTasks)
}
