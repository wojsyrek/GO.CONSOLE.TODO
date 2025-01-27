package commands

import (
	"fmt"
	"slices"
	"strconv"

	"GO.CONSOLE.TODO/internal/models"
	"GO.CONSOLE.TODO/internal/storage"
)

type CommandDelete struct{}

func (c CommandDelete) Execute(params map[string]string) error {

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return fmt.Errorf("Wrong id passed")
	}

	data, err := storage.ReadJson[[]models.TaskRow]("data.json")

	if err != nil {
		data = &[]models.TaskRow{}
	}

	for v := range *data {
		if (*data)[v].TaskId == id {
			*data = slices.Delete(*data, v, v+1)
			fmt.Printf("Task with ID %d was successfully deleted\n", id)
			storage.WriteJson("data.json", *data)
			return nil
		}
	}

	fmt.Printf("Task with ID %d was not found\n", id)

	return nil
}
