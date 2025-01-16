package commands

import (
	"cmp"
	"encoding/json"
	"fmt"
	"slices"
	"time"

	"GO.CONSOLE.TODO/internal/models"
	"GO.CONSOLE.TODO/internal/storage"
)

type CommandAdd struct{}

func (c CommandAdd) Execute(params map[string]string) error {
	data, err := storage.ReadJson[[]models.TaskRow]("data.json")
	if err != nil {
		data = &[]models.TaskRow{}
	}

	last := 0
	if len(*data) != 0 {
		last = slices.MaxFunc(*data, func(a, b models.TaskRow) int {
			return cmp.Compare(a.TaskId, b.TaskId)
		}).TaskId
	}

	newTask := models.TaskRow{}
	newTask.TaskId = last + 1

	for k, v := range params {
		switch k {
		case "name":
			newTask.TaskName = v
		case "desc":
			newTask.TaskDescription = v
		case "deadline":
			deadline, err := time.Parse("02/01/2006 15:04:05", v)
			if err != nil {
				return err
			}
			if deadline.Before(time.Now()) {
				return fmt.Errorf("deadline %s must be in the future", params["value"])
			}
			newTask.TaskDeadline = deadline
		default:
			fmt.Printf("unknown parameter %s, skipping...\n", k)
		}
	}

	newTask.Events = append(newTask.Events, models.Event{
		EventType:        models.Create,
		EventDate:        time.Now(),
		EventDescription: "",
	})

	*data = append(*data, newTask)

	storage.WriteJson("data.json", *data)
	fmt.Println("Success!")
	json, err := json.MarshalIndent(newTask, "", "   ")

	if err != nil {
		return err
	}

	fmt.Println(string(json))

	return nil
}
