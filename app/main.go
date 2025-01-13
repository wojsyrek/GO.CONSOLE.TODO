package main

import (
	"fmt"
	"os"

	config "GO.CONSOLE.TODO/pkg"
	"GO.CONSOLE.TODO/pkg/commands"
)

func main() {
	Welcome()
	Init()
}

func Welcome() {
	fmt.Println(config.Constants.AppName + " " + config.Constants.Version)
	fmt.Println(config.Constants.WelcomeMsg)
}

func Init() {
	var com string
	fmt.Scanln(&com)

	if com == "exit" {
		os.Exit(1)
	} else {
		result := commands.Execute(com)

		if result != nil {
			println(result.Error())
		}
	}

	Init()
}
