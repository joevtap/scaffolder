package scaffolder

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/go-playground/validator/v10"
)

// createDirectoryTree creates a directory tree based on the provided definition.
//
//   - projectRoot is the field in the Project struct that represents the root of the directory tree.
//   - destinationPath is the path to the directory where the directory tree will be created.
//     The path is relative to the current working directory.
//   - appRootPath is the absolute path to the root of the application.
//   - data is the template data to be used when rendering the project definition file.
func createDirectoryTree(projectRoot []Block, destinationPath, appRootPath string, data interface{}) error {
	for _, block := range projectRoot {
		switch block.Type {
		case "dir":
			dirPath := filepath.Join(destinationPath, block.Name)
			err := createDir(dirPath)
			if err != nil {
				return err
			}

			if len(block.Children) > 0 {
				createDirectoryTree(block.Children, dirPath, appRootPath, data)
			}

		case "file":
			filePath := filepath.Join(destinationPath, block.Name)
			err := createFile(filePath, block, appRootPath, data)
			if err != nil {
				return err
			}

		default:
			return fmt.Errorf("invalid block type %v", block.Type)
		}
	}

	return nil
}

// createDir creates a directory at the provided path.
//
//   - dirPath is the path to the directory to be created.
func createDir(dirPath string) error {
	if err := os.MkdirAll(dirPath, os.FileMode(0755)); err != nil {
		return fmt.Errorf("error creating directory %v: %v", dirPath, err)
	}

	return nil
}

// createFile creates a file at the provided destinationPath.
//
//   - destinationPath is the path to the file to be created.
//   - definition is the definition of the file to be created.
//   - appRoot is the absolute path to the root of the application.
//   - data is the template data to be used when rendering the project definition file.
func createFile(destinationPath string, definition Block, appRootPath string, data interface{}) error {
	file, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("error creating file %v: %v", destinationPath, err)
	}
	defer file.Close()

	if definition.Template != "" {
		templateFilePath := filepath.Join(appRootPath, definition.Template)
		templateFile, err := os.ReadFile(templateFilePath)
		if err != nil {
			return fmt.Errorf("error reading template file %v: %v", templateFilePath, err)
		}

		tpl, err := template.New(definition.Name).Parse(string(templateFile))
		if err != nil {
			return fmt.Errorf("error parsing template file %v: %v", templateFilePath, err)
		}
		err = tpl.Execute(file, data)
		if err != nil {
			return fmt.Errorf("error executing template file %v: %v", templateFilePath, err)
		}
	}

	return nil
}

func validateProject(project Project) error {
	v := validator.New()
	return v.Struct(project)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}
