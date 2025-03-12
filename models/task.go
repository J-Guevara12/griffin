package models

import (
	"fmt"
	"time"
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
    Summary string
    Description string
    Notes string
    DueDate time.Time
    Priority Priority
    Status Status
}

func NewTask(summary string, Description string, notes string, due_date time.Time, priority Priority, status Status) Task {
    switch priority{
    case Highest, High, Medium, Low, Lowest:
    default:
        panic(fmt.Sprintf("Priority value (%v) not suported!", priority))
    }

    switch status{
    case Closed, WorkingOn, NotOnYou, Starting, NotYetStarted:
    default:
        panic(fmt.Sprintf("Status value (%v) not suported!", status))
    }
    return Task{
        Summary: summary,
        Description: Description,
        Notes: notes,
        DueDate: due_date,
        Priority: priority,
        Status: status,
    }
}
