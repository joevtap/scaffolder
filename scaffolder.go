package scaffolder

import (
	"fmt"

	"github.com/joevtap/scaffolder/config"
)

// Scaffold creates a new project based on the provided config file
// and template data.
//
//   - destinationPath is the path to the directory where the project will be created.
//     destinationPath is relative to the current working directory.
//
//   - configPath is the path to the project config file.
//     configPath is relative to appRootPath.
//
//   - appRootPath is the absolute path to the root of the application.
//     appRootPath is used to resolve relative paths in the project definition file.
//
//   - data is the template data to be used when rendering the project config file.
func Scaffold(destinationPath, configPath, appRootPath string, data interface{}) error {
	var p Project
	if err := config.Read(configPath, &p); err != nil {
		return err
	}

	if err := validateProject(p); err != nil {
		return fmt.Errorf("project config is invalid: %v", err)
	}

	if pathExists(destinationPath) {
		return fmt.Errorf("project directory already exists: %v", destinationPath)
	}

	if err := createDir(destinationPath); err != nil {
		return fmt.Errorf("error creating project directory: %v", err)
	}

	if err := createDirectoryTree(p.Root, destinationPath, appRootPath, data); err != nil {
		return fmt.Errorf("error scaffolding project: %v", err)
	}

	return nil
}
