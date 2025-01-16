package commands

import (
	"fmt"

	"GO.CONSOLE.TODO/internal/models"
	"GO.CONSOLE.TODO/internal/storage"
)

type CommandRead struct{}

func (c CommandRead) Execute(params map[string]string) error {
	data, err := storage.ReadJson[[]models.TaskRow]("data.json")

	if len(*data) == 0 {
		fmt.Println("No tasks found")
	} else {
		PrintObject(data)

		if err != nil {
			return nil
		}
	}
	fmt.Println()
	return nil
}
