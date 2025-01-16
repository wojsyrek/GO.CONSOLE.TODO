package commands

import "fmt"

type CommandDelete struct{}

func (c CommandDelete) Execute(params map[string]string) error {
	fmt.Println("CommandDelete Execute")
	return nil
}
