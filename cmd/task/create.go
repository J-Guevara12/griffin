/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"fmt"
	"griffin/models"
	"os"

	"github.com/spf13/cobra"
)


// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new task",
	Long: `Creates a new Task. You must provide at least the Summary (Title) and a Description.

In order to set the due date, you can provide a comma separated list of increments (supports
negative floating point numbers) with a unit (minutes, hour, days, weeks, months) to the -T flag.

For every unit (except months) you can provide just the initial letter (m, h, d, w)

Returns the created task ID.
Example:
--------
Add a task of high priority, which is being worked on and is due to one month minus one and a half weeks:

    griffin task create -S "Sample Summary" \
                        -D "A short description" \
                        -P "High" -s "Working on" \
                        -T "1 month, -1.5w"
`,
	Run: func(cmd *cobra.Command, args []string) {
        due_time := parse_timedelta(timedelta)
        task := models.NewTask(summary, description, notes, due_time, priority, status)
        fmt.Println(configured_db().WriteTask(task))
        os.Exit(0)
	},
}

func init() {
	taskCmd.AddCommand(createCmd)
    createCmd.Flags().StringVarP(&summary, "summary", "S", "", "Task summary (or Title) (Required)")
    createCmd.Flags().StringVarP(&description, "description", "D", "", "Task Description (Required)")
    createCmd.Flags().StringVarP(&notes, "notes", "N", "", "Task Notes (Optional)")
    createCmd.Flags().StringVarP(&status, "status", "s", "Starting", "Status (default Starting)")
    createCmd.Flags().StringVarP(&priority, "priority", "P", "Lowest", "Task priority (default Lowest)")
    createCmd.Flags().StringVarP(&timedelta, "time", "T", "1d", "How many time do you have to complete the task?")

    createCmd.MarkFlagRequired("summary")
    createCmd.MarkFlagRequired("description")
}
