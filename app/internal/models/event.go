package models

type Event struct {
	EventType        EventType `json:"eventType"`
	EventData        string    `json:"eventData"`
	EventDescription string    `json:"eventDescription"`
}

type EventType int

const (
	Create EventType = iota // 0
	Update                  // 1
	Delete                  // 2
	Read                    // 3
)

// Add String method for better printing
func (e EventType) String() string {
	return [...]string{"CREATE", "UPDATE", "DELETE", "READ"}[e]
}
