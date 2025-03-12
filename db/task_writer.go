package db

import "griffin/models"

type TaskWriter interface {
    WriteTask(task *models.Task) error
    GetAllTasks() (*[]models.Task, error)
}

