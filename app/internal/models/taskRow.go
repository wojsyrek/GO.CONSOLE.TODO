package models

import "time"

type TaskRow struct {
	TaskId          int       `json:"taskId"`
	TaskName        string    `json:"taskName"`
	TaskDescription string    `json:"taskDescription"`
	TaskDeadline    time.Time `json:"taskDeadline"`
	Events          []Event   `json:"events"`
}
