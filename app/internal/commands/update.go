package commands

import "fmt"

type CommandUpdate struct{}

func (c CommandUpdate) Execute(params map[string]string) error {
	fmt.Println("CommandUpdate Execute")
	return nil
}
