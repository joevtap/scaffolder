package scaffolder

import (
	"fmt"
	"os"

	"github.com/joevtap/scaffolder/pkg/definitionparser"
)

// Scaffold creates a new project based on the provided definition file
// and template data.
//
//   - destinationPath is the path to the directory where the project will be created.
//     destinationPath is relative to the current working directory.
//
//   - definitionPath is the path to the project definition file.
//     definitionPath is relative to appRootPath.
//
//   - appRootPath is the absolute path to the root of the application.
//     appRootPath is used to resolve relative paths in the project definition file.
//
//   - data is the template data to be used when rendering the project definition file.
func Scaffold(destinationPath, definitionPath, appRootPath string, data interface{}) error {
	project, err := os.ReadFile(definitionPath)
	if err != nil {
		return fmt.Errorf("error reading project definition file: %v", err)
	}

	var p Project
	err = definitionparser.ParseFile(project, &p)
	if err != nil {
		return fmt.Errorf("error unmarshaling TOML file: %v", err)
	}

	err = validateProject(p)
	if err != nil {
		return fmt.Errorf("project definition is invalid: %v", err)
	}

	if pathExists(destinationPath) {
		return fmt.Errorf("project directory already exists: %v", destinationPath)
	}

	err = createDir(destinationPath)
	if err != nil {
		return fmt.Errorf("error creating project directory: %v", err)
	}

	err = createDirectoryTree(p.Root, destinationPath, appRootPath, data)
	if err != nil {
		return fmt.Errorf("error scaffolding project: %v", err)
	}

	return nil
}
