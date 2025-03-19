/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package task

import (
	"github.com/spf13/cobra"
    "griffin/models"
    "time"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a task",
	Long: `Updates a task in the database. In order to select the appropiate task you must provide its ID run 'griffin task ls' in order to get it`,
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        id := args[0]
        task := configured_db().GetTaskByID(id)
        
        if summary != "" {
            task.Summary = summary
        }
        if description != "" {
            task.Description = description
        }
        if notes != "" {
            task.Notes = notes
        }
        if status != "" {
            new_status := models.Status(status)
            if new_status == models.Closed && task.Status != models.Closed {
                task.Closed = time.Now()
            }

            task.Status = new_status
        }
        if priority != "" {
            task.Priority = models.Priority(priority)
        }

        if timedelta != "" {
            task.DueDate = parse_timedelta(timedelta)
        }

        task.Modified = time.Now()
        configured_db().UpdateTask(&task)

	},
}

func init() {
	taskCmd.AddCommand(updateCmd)

    updateCmd.Flags().StringVarP(&summary, "summary", "S", "", "Task summary (or Title)")
    updateCmd.Flags().StringVarP(&description, "description", "D", "", "Task Description")
    updateCmd.Flags().StringVarP(&notes, "notes", "N", "", "Task Notes")
    updateCmd.Flags().StringVarP(&status, "status", "s", "", "Status")
    updateCmd.Flags().StringVarP(&priority, "priority", "P", "", "Task priority")
    updateCmd.Flags().StringVarP(&timedelta, "time", "T", "", "How many time do you have to complete the task?")
}
