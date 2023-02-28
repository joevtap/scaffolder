package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Read reads the project config file at configPath and unmarshals it into the provided v interface.
func Read(configPath string, v interface{}) error {
	project, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading project config file: %v", err)
	}

	if err = unmarshal(project, &v); err != nil {
		return fmt.Errorf("error unmarshaling config file: %v", err)
	}

	return nil
}

func unmarshal(data []byte, v interface{}) error {
	return toml.Unmarshal(data, &v)
}
