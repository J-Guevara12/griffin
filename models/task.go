package models

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	"github.com/mergestat/timediff"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Priority string
type Status string 

const (
    Highest Priority = "Highest"
    High = "High"
    Medium = "Medium"
    Low = "Low"
    Lowest = "Lowest"
)

const (
    Closed Status = "Closed"
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
    Priority Priority
    Status Status
}

// Task Constructor. Every time you intend to create a new task you should use this function
func NewTask(summary string, Description string, notes string, due_date time.Time, priority Priority, status Status) *Task {
    switch priority{
    case Highest, High, Medium, Low, Lowest:
    default:
        fmt.Fprintf(os.Stderr, "Priority value (%v) not suported!\n", priority)
        os.Exit(1)
    }

    switch status{
    case Closed, WorkingOn, NotOnYou, Starting, NotYetStarted:
    default:
        fmt.Fprintf(os.Stderr, "Status value (%v) not suported!\n", status)
        os.Exit(1)
    }
    return &Task{
        ID: bson.NewObjectID(),
        Summary: summary,
        Description: Description,
        Notes: notes,
        DueDate: due_date,
        Priority: priority,
        Status: status,
    }
}

func (task Task) To_ls_table() []string {
    return []string{
        task.ID.Hex(),
        task.Summary,
        task.DueDate.Format("02-Jan-2006 15:04:05") + fmt.Sprintf("\n(%v)", timediff.TimeDiff(task.DueDate)),
        string(task.Priority),
        string(task.Status),
    }
}

func CreateTaskTable(tasks []Task) string {
        rows := make([][]string, 0)
        for _, task := range tasks {
            rows = append(rows, task.To_ls_table())
        }

        var (
            purple    = lipgloss.Color("99")
            gray      = lipgloss.Color("245")
            lightGray = lipgloss.Color("241")
            white     = lipgloss.Color("#FFFFFF")

            headerStyle  = lipgloss.NewStyle().Foreground(white).Bold(true).Align(lipgloss.Center).Background(purple)

            priorityColor = map[Priority]lipgloss.Color{
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
            switch col {
                case 0:
                    style = style.Width(28)
                case 1:
                    style = style.Width(50).Align(lipgloss.Left)
                case 3:
                    priority := fmt.Sprint(rows[row][col])
                    style = style.Foreground(priorityColor[Priority(priority)])
            }
            return style
            }).
        Headers("ID", "Summary", "Due date", "Priority", "Status").
        Rows(rows...)
        return fmt.Sprint(t)
}
