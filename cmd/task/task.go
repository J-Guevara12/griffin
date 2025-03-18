/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package task

import (
	"github.com/spf13/cobra"
    "griffin/cmd"
    "griffin/db"
	"time"
	"strconv"
	"regexp"
	s "strings"
    "os"
    "fmt"
)

const date_regex = `^\s*(-?\d+\.?\d*)\s*(d|h|m|w|months?|days?|hours?|minutes?|weeks?)\s*$`

var summary string
var description string
var notes string
var timedelta string
var priority string
var status string

func parse_timedelta(timedelta string) time.Time {
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
    return due_time
}

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
