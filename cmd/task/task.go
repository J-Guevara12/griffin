/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package task

import (
	"github.com/spf13/cobra"
    "griffin/cmd"
    "griffin/db"
)

func configured_db() db.DBConnector{
    return db.NewDBConnector("mongodb://localhost:27017", "test", "Tasks", 5)
}

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task interface for CRUD operations",
	Long: `This interface is designed to perform CRUD operations on the task database.`,
}

func init() {
	cmd.RootCmd.AddCommand(taskCmd)
}
