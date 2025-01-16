package main

import (
	"bufio"
	"fmt"
	"os"

	"GO.CONSOLE.TODO/internal/commands"
	"GO.CONSOLE.TODO/internal/models"
)

func main() {
	config, err := models.GetConfig(models.DEFAULT_CONFIG_FILE)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(config.AppName + " " + config.Version)
	fmt.Println(config.WelcomeMsg)

	Init()
}

func Init() {
	var com string

	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			com = scanner.Text()
		}

		if com == "exit" {
			os.Exit(1)
		} else {
			err := commands.Execute(com)

			if err != nil {
				println(err.Error())
			}
		}
	}
}
