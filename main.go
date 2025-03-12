/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

//import "griffin/cmd"
import (
	"griffin/db"
	"griffin/models"
	"time"
)

func main() {
	//cmd.Execute()
    task := models.NewTask(
        "Title 1",
        "# Description 1:\nMay be a longer test\n## Has over one title",
        "* Various notes\n* Various Notes line 2",
        time.Now().Add(time.Duration(24) * time.Hour),
        models.Highest,
        models.Starting,
    )
    conn := db.NewDBConnector("mongodb://localhost:27017", "test", "Tasks", 5)
    conn.WriteTask(task)
    task.DueDate = time.Now()
    task.Description = "A Solo aparecer una vez! jajaj"
    conn.UpdateTask(task)
    conn.DeleteTask(task)
}
