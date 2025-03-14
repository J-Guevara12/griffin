package models

import (
	"fmt"
	"os"
	"time"

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
