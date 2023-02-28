package scaffolder

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// createDirectoryTree creates a directory tree based on the provided config.
//
//   - projectRoot is the field in the Project struct that represents the root of the directory tree.
//   - destinationPath is the path to the directory where the directory tree will be created.
//     The path is relative to the current working directory.
//   - appRootPath is the absolute path to the root of the application.
//   - data is the template data to be used when rendering the project config file.
func createDirectoryTree(projectRoot []Block, destinationPath, appRootPath string, data interface{}) error {
	for _, block := range projectRoot {
		switch block.Type {
		case "dir":
			dirPath := filepath.Join(destinationPath, block.Name)
			if err := createDir(dirPath); err != nil {
				return err
			}

			if len(block.Children) > 0 {
				createDirectoryTree(block.Children, dirPath, appRootPath, data)
			}

		case "file":
			filePath := filepath.Join(destinationPath, block.Name)
			if err := createFile(filePath, block, appRootPath, data); err != nil {
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
//   - block is the definition of the file to be created.
//   - appRoot is the absolute path to the root of the application.
//   - data is the template data to be used when rendering the project definition file.
func createFile(destinationPath string, block Block, appRootPath string, data interface{}) error {
	file, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("error creating file %v: %v", destinationPath, err)
	}
	defer file.Close()

	if block.Template != "" {
		templateFilePath := filepath.Join(appRootPath, block.Template)

		templateFile, err := os.ReadFile(templateFilePath)
		if err != nil {
			return fmt.Errorf("error reading template file %v: %v", templateFilePath, err)
		}

		tpl, err := template.New(block.Name).Parse(string(templateFile))
		if err != nil {
			return fmt.Errorf("error parsing template file %v: %v", templateFilePath, err)
		}

		if err := tpl.Execute(file, data); err != nil {
			return fmt.Errorf("error executing template file %v: %v", templateFilePath, err)
		}
	}

	return nil
}
