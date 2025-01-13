package commands

import (
	"fmt"
	"reflect"

	"GO.CONSOLE.TODO/pkg/dataModel"
	"GO.CONSOLE.TODO/pkg/jsonHelper"
)

type Command interface {
	Execute(params string) error
}

type CommandHelp struct{}
type CommandAdd struct{}
type CommandDelete struct{}
type CommandUpdate struct{}
type CommandRead struct{}

func (c CommandHelp) Execute(params string) error {
	fmt.Println("CommandHelp Execute")
	return nil
}

func (c CommandAdd) Execute(params string) error {
	fmt.Println("CommandAdd Execute")
	return nil
}

func (c CommandDelete) Execute(params string) error {
	fmt.Println("CommandDelete Execute")
	return nil
}

func (c CommandUpdate) Execute(params string) error {
	fmt.Println("CommandUpdate Execute")
	return nil
}

func (c CommandRead) Execute(params string) error {
	data, err := jsonHelper.ReadJson[dataModel.TaskRowWrapper]("data.json")

	t := reflect.TypeOf(dataModel.TaskRow{})
	fmt.Println()
	for _, sData := range data.Datas {
		printObject(sData, t)
	}

	if err != nil {
		return nil
	}

	return nil
}

func printObject[T any](data T, t reflect.Type) {

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.Kind() == reflect.Array {
			//todo, to split it after
			continue
		}
		v := reflect.ValueOf(data)
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}

func GetData() dataModel.TaskRowWrapper {
	result, err := jsonHelper.ReadJson[dataModel.TaskRowWrapper]("data.json")

	if err != nil {
		fmt.Println(err.Error())
		result = &dataModel.TaskRowWrapper{
			Datas: []dataModel.TaskRow{},
		}
	}

	return *result
}

var commands = map[string]Command{
	"help":   CommandHelp{},
	"add":    CommandAdd{},
	"delete": CommandDelete{},
	"update": CommandUpdate{},
	"read":   CommandRead{},
}

func Execute(commandName string) error {
	if cmd, exists := commands[commandName]; exists {
		return cmd.Execute(commandName)
	}
	return fmt.Errorf("command %s not found", commandName)
}
