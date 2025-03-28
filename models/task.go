package models

import (
	"fmt"
	"os"
	"time"

    s "strings"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	"github.com/mergestat/timediff"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type priority string
type status string 

func Priority(str string) priority {
    switch priority(str){
    case Highest, High, Medium, Low, Lowest:
        return priority(str)
    default:
        fmt.Fprintf(os.Stderr, "Priority value (%v) not suported!\n", str)
        os.Exit(1)
        return ""
    }
}

func Status(str string) status {
    switch status(str){
    case Closed, WorkingOn, NotOnYou, Starting, NotYetStarted:
        return status(str)
    default:
        fmt.Fprintf(os.Stderr, "Status value (%v) not suported!\n", str)
        os.Exit(1)
        return ""
    }
}

const (
    Highest priority = "Highest"
    High = "High"
    Medium = "Medium"
    Low = "Low"
    Lowest = "Lowest"
)

const (
    Closed status = "Closed"
    WorkingOn =     "Working on"
    NotOnYou =      "Not on you"
    Starting =      "Starting"
    NotYetStarted = "Not yet started"
)
 
type Task struct {
    ID bson.ObjectID `bson:"_id"`
    Summary string
    Description string
    Notes string
    DueDate time.Time
    Created time.Time
    Closed time.Time
    Modified time.Time
    Priority priority
    Status status
}

// Task Constructor. Every time you intend to create a new task you should use this function
func NewTask(summary string, description string, notes string, due_date time.Time, priority string, status string) *Task {
    _priority := Priority(priority)
    _status := Status(status)
    return &Task{
        ID: bson.NewObjectID(),
        Summary: summary,
        Description: description,
        Notes: notes,
        DueDate: due_date,
        Created: time.Now(),
        Priority: _priority,
        Status: _status,
    }
}

func (task Task) To_ls_table(columns[]string) []string {
    output := make([]string, len(columns))
    for i, column := range columns{
        switch s.ToLower(column) {
        case "id":
            output[i] = task.ID.Hex()
        case "summary":
            output[i] = task.Summary
        case "description":
            output[i] = task.Description
        case "notes":
            output[i] = task.Notes
        case "due_date":
            output[i] = task.DueDate.Local().Format("02-Jan-2006 15:04:05") + fmt.Sprintf("\n(%v)", timediff.TimeDiff(task.DueDate))
        case "created":
            output[i] = task.Created.Local().Format("02-Jan-2006 15:04:05") + fmt.Sprintf("\n(%v)", timediff.TimeDiff(task.Created))
        case "modified":
            output[i] = task.Modified.Local().Format("02-Jan-2006 15:04:05") + fmt.Sprintf("\n(%v)", timediff.TimeDiff(task.Modified))
        case "closed":
            output[i] = task.Closed.Local().Format("02-Jan-2006 15:04:05") + fmt.Sprintf("\n(%v)", timediff.TimeDiff(task.Closed))
        case "priority":
            output[i] = string(task.Priority)
        case "status":
            output[i] = string(task.Status)
        default:
            fmt.Fprintf(os.Stderr, "Column value '%v' not recognized!\n", column)
            os.Exit(1)

        }

    }
    return output
}

func CreateTaskTable(tasks []Task, columns []string) string {
        rows := make([][]string, 0)

        for _, task := range tasks {
            rows = append(rows, task.To_ls_table(columns))
        }

        headers := make([]string, len(columns))

        for i, column := range columns {
            headers[i] = s.Replace(s.ToUpper(column), "_", " ", -1)

        }

        var (
            purple    = lipgloss.Color("99")
            gray      = lipgloss.Color("245")
            lightGray = lipgloss.Color("243")
            white     = lipgloss.Color("#FFFFFF")

            headerStyle  = lipgloss.NewStyle().Foreground(white).Bold(true).Align(lipgloss.Center).Background(purple)

            priorityColor = map[priority]lipgloss.Color{
                Lowest:  lipgloss.Color("#006c7d"),
                Low:     lipgloss.Color("#04d18d"),
                Medium:  lipgloss.Color("#ebf705"),
                High:    lipgloss.Color("#f78205"),
                Highest: lipgloss.Color("#f72d05"),
            }
        )
        t:= table.New().
        Border(lipgloss.NormalBorder()).
        BorderStyle(lipgloss.NewStyle().Foreground(purple)).StyleFunc(func(row, col int) lipgloss.Style {
            style := lipgloss.NewStyle().Padding(1, 2, 0).Align(lipgloss.Center)
            switch {
                case row == table.HeaderRow:
                    return headerStyle
                case row%2 == 0:
                    style = style.Foreground(lightGray)
                default:
                    style = style.Foreground(gray)
                }
            switch s.ToLower(headers[col]) {
                case "id":
                    style = style.Width(28)
                case "summary":
                    style = style.Width(50).Align(lipgloss.Left)
                case "priority":
                    priority := fmt.Sprint(rows[row][col])
                    style = style.Foreground(priorityColor[Priority(priority)])
            }
            return style
            }).
        Headers(headers...).
        Rows(rows...)
        return fmt.Sprint(t)
}
