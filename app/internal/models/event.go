package models

import "time"

type Event struct {
	EventType        string    `json:"eventType"`
	EventDate        time.Time `json:"eventDate"`
	EventDescription string    `json:"eventDescription"`
}

const (
	Create = "CREATED"
	Update = "UPDATED"
	Delete = "DELETED"
)
