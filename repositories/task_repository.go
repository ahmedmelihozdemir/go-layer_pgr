package repositories

import (
	"go-layer-pgr-ex/models"
	"go-layer-pgr-ex/config"
)

func CreateTask(task *models.Task) error {
	return config.DB.Create(task).Error
}

func GetTasks(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}
