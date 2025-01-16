package commands

import "fmt"

type CommandHelp struct{}

func (c CommandHelp) Execute(params map[string]string) error {
	fmt.Println("CommandHelp Execute")
	return nil
}
