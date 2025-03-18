/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"griffin/models"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Deletes a task",
	Long: `Deletes the task by its ID. to know the task ID execute 'griffin task ls'.`,
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        id, err := bson.ObjectIDFromHex(args[0])
        if err != nil {
            panic(err)
        }
        task := models.Task{ID: id}
        configured_db().DeleteTask(&task)
	},
}

func init() {
	taskCmd.AddCommand(deleteCmd)
}
