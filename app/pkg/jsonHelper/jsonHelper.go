package jsonHelper

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadJson[T any](path string) (*T, error) {
	// Open file
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer jsonFile.Close()

	// Read file
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse JSON
	var result T
	if err := json.Unmarshal(byteValue, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &result, nil
}

func WriteJson[T any](path string, data T) error {
	dataString, err := json.Marshal(data)

	if err != nil {
		fmt.Println("failed to marshal JSON: %w", err)
		return err
	}

	if err := os.WriteFile("data.json", dataString, 0644); err != nil {
		fmt.Println("failed to write file: %w", err)
		fmt.Println("try again? Y/N: %w", err)

		var res string
		fmt.Scanln(&res)

		if res == "N" {
			return nil
		}

		WriteJson(path, data)
	}

	return nil
}
