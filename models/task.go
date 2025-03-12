package models

import "time"

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
    DueDate time.Time
    Priority Priority
    Status Status
}
