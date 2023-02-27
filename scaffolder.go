package scaffolder

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func Scaffold(path, definition string, templateData interface{}) error {
	project, err := os.ReadFile(definition)
	if err != nil {
		return fmt.Errorf("error reading project definition file: %v", err)
	}

	var p Project
	err = toml.Unmarshal(project, &p)
	if err != nil {
		return fmt.Errorf("error unmarshaling TOML file: %v", err)
	}

	err = validateProject(p)
	if err != nil {
		return fmt.Errorf("project definition is invalid: %v", err)
	}

	if pathExists(path) {
		return fmt.Errorf("project directory already exists: %v", path)
	}

	err = createDir(path)
	if err != nil {
		return fmt.Errorf("error creating project directory: %v", err)
	}

	err = createDirectoryTree(p.Root, path, templateData)
	if err != nil {
		return fmt.Errorf("error scaffolding project: %v", err)
	}

	return nil
}
