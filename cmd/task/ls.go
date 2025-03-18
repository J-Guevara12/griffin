/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"fmt"
	"griffin/models"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all tasks",
	Long: `Displays a table with all the tasks in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
        tasks := configured_db().GetAllTasks()

        fmt.Println(models.CreateTaskTable(tasks))
        fmt.Println(configured_db().GetTaskByID("67d3aa9f77491e00cc9e20cc"))
	},
}

func init() {
	taskCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
