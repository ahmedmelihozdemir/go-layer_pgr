package handlers

import (
	"go-layer-pgr-ex/models"
	"go-layer-pgr-ex/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := services.CreateTask(task); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot create task"})
	}

	return c.Status(201).JSON(task)
}

func GetTasks(c *fiber.Ctx) error {
	userID, _ := strconv.Atoi(c.Params("userID"))

	tasks, err := services.GetTasks(uint(userID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot fetch tasks"})
	}

	return c.JSON(tasks)
}
