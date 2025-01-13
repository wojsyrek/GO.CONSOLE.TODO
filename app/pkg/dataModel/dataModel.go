package dataModel

type TaskRowWrapper struct {
	Datas []TaskRow `json:"data"`
}

type TaskRow struct {
	TaskId          string  `json:"taskId"`
	TaskName        string  `json:"taskName"`
	TaskDescription string  `json:"taskDescription"`
	TaskDeadline    string  `json:"taskDeadline"`
	Events          []Event `json:"events"`
}

type Event struct {
	EventType        string `json:"eventType"`
	EventData        string `json:"eventData"`
	EventDescription string `json:"eventDescription"`
}
