package commands

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Command interface {
	Execute(params map[string]string) error
}

func PrintObject[T any](data T) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(jsonData))
}

var commands = map[string]Command{
	"help":   CommandHelp{},
	"add":    CommandAdd{},
	"delete": CommandDelete{},
	"update": CommandUpdate{},
	"read":   CommandRead{},
}

func Execute(command string) error {
	if command == "" {
		return fmt.Errorf("empty command")
	}

	params, err := ReadParams(command)
	if err != nil {
		return err
	}

	if len(params) > 0 {
		idx := strings.Index(command, " -")
		if idx == -1 {
			return fmt.Errorf("invalid command format")
		}
		command = command[0:idx]
	}

	if cmd, exists := commands[command]; exists {
		return cmd.Execute(params)
	}
	return fmt.Errorf("command %s not found", command)
}

func ReadParams(paramString string) (map[string]string, error) {
	parts := strings.SplitN(paramString, " -", 2)
	if len(parts) < 2 {
		return nil, nil
	}

	var params = map[string]string{}
	ps := parts[1]

	for _, param := range strings.Split(ps, "-") {
		p := strings.SplitN(strings.TrimSpace(param), " ", 2)
		if len(p) != 2 {
			return nil, fmt.Errorf("invalid parameter format: %s", param)
		}

		paramName := strings.TrimSpace(p[0])
		if paramName == "" {
			return nil, fmt.Errorf("empty parameter name")
		}

		params[paramName] = strings.TrimSpace(p[1])
	}

	return params, nil
}
