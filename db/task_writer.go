package db

import "griffin/models"

type TaskWriter interface {
    WriteTask(task *models.Task) string
    GetAllTasks() []models.Task
    UpdateTask(task *models.Task)
    DeleteTask(task *models.Task)
    GetTaskByID(id string) models.Task
}

