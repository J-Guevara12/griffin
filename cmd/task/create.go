/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"fmt"
	"griffin/models"
	"os"
	"regexp"
	"strconv"
	s "strings"
	"time"

	"github.com/spf13/cobra"
)

const date_regex = `^\s*(-?\d+\.?\d*)\s*(d|h|m|w|months?|days?|hours?|minutes?|weeks?)\s*$`

var summary string
var description string
var notes string
var timedelta string
var priority string
var status string

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
        reg, err := regexp.Compile(date_regex)
        if err != nil {
            panic(err)
        }
        tokens := s.Split(timedelta, ",")
        due_time := time.Now()
        for _, token := range tokens {
            values := reg.FindStringSubmatch(token)
            if values != nil {
                quantity, err := strconv.ParseFloat(values[1], 32)
                if err!=nil{
                    panic(err)
                }
                var duration time.Duration
                switch values[2] {
                    case "d", "day", "days":
                        duration = time.Duration(int64(float64(time.Second*60*60*24)*quantity))
                    case "h", "hour", "hours":
                        duration = time.Duration(int64(float64(time.Second*60*60)*quantity))
                    case "m", "minute", "minutes":
                        duration = time.Duration(int64(float64(time.Second*60)*quantity))
                    case "w", "week", "weeks":
                        duration = time.Duration(int64(float64(time.Second*60*60*24*7)*quantity))
                    case "month", "months":
                        duration = (time.Now().AddDate(0, int(quantity), 0).Sub(time.Now()))
                    default:
                        fmt.Fprintf(os.Stderr, "Error parsing the relative time token (%v)\n", token)
                        os.Exit(1)
                }
                due_time = due_time.Add(duration)
            } else {
                fmt.Fprintf(os.Stderr, "Error parsing the relative time token (%v)\n", token)
                os.Exit(1)
            }

        }
        task := models.NewTask(summary, description, notes, due_time, models.Priority(priority), models.Status(status))
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
    createCmd.Flags().StringVarP(&timedelta, "time", "T", "1d", "How do you have to complete the task?")

    createCmd.MarkFlagRequired("summary")
    createCmd.MarkFlagRequired("description")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
