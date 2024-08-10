package services

import (
	"encoding/json"
	"fmt"
	"go-layer-pgr-ex/config"
	"go-layer-pgr-ex/models"
	"go-layer-pgr-ex/repositories"
)

func CreateTask(task *models.Task) error {
	err := repositories.CreateTask(task)
	if err == nil {
		config.RDB.Del(config.Ctx, fmt.Sprintf("tasks_%d", task.UserID))
	}
	return err
}

func GetTasks(userID uint) ([]models.Task, error) {
	cacheKey := fmt.Sprintf("tasks_%d", userID)
	val, err := config.RDB.Get(config.Ctx, cacheKey).Result()
	if err == nil {
		var tasks []models.Task
		json.Unmarshal([]byte(val), &tasks)
		return tasks, nil
	}

	tasks, err := repositories.GetTasks(userID)
	if err == nil {
		data, _ := json.Marshal(tasks)
		config.RDB.Set(config.Ctx, cacheKey, data, 0)
	}
	return tasks, err
}
