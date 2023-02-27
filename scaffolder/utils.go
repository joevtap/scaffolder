package scaffolder

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/go-playground/validator/v10"
)

func createDirectoryTree(block []Block, parentPath string, templateData interface{}) error {
	for _, block := range block {
		switch block.Type {
		case "dir":
			dirPath := filepath.Join(parentPath, block.Name)
			err := createDir(dirPath)
			if err != nil {
				return err
			}

			if len(block.Children) > 0 {
				createDirectoryTree(block.Children, dirPath, templateData)
			}

		case "file":
			filePath := filepath.Join(parentPath, block.Name)
			err := createFile(filePath, block, templateData)
			if err != nil {
				return err
			}

		default:
			return fmt.Errorf("invalid block type %v", block.Type)
		}
	}

	return nil
}

func createDir(dirPath string) error {
	if err := os.MkdirAll(dirPath, os.FileMode(0755)); err != nil {
		return fmt.Errorf("error creating directory %v: %v", dirPath, err)
	}

	return nil
}

func createFile(filePath string, block Block, templateData interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file %v: %v", filePath, err)
	}
	defer file.Close()

	if block.Template != "" {
		templateFilePath := filepath.FromSlash(block.Template)
		templateFile, err := os.ReadFile(templateFilePath)
		if err != nil {
			return fmt.Errorf("error reading template file %v: %v", templateFilePath, err)
		}

		tpl, err := template.New(block.Name).Parse(string(templateFile))
		if err != nil {
			return fmt.Errorf("error parsing template file %v: %v", templateFilePath, err)
		}
		err = tpl.Execute(file, templateData)
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
