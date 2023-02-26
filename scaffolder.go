package scaffolder

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Dir struct {
	Dirs  map[string]Dir `yaml:"dirs"`
	Files []File         `yaml:"files"`
}

type File struct {
	Name    string `yaml:"name"`
	Content string `yaml:"content,omitempty"`
}

func Scaffold(templatePath, outputDirPath string, data map[string]string) error {
	log.SetPrefix("func Scaffold(string, string, map[string]string) error:")
	var dir Dir

	yamlData, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlData, &dir)
	if err != nil {
		log.Fatal(err)
	}

	return dir.Scaffold(outputDirPath, data)
}

func (d Dir) Scaffold(parentDirPath string, data map[string]string) error {
	log.SetPrefix("func (d Dir) Scaffold(string, map[string]string) error:")

	for key, subDir := range d.Dirs {
		subDirPath := filepath.Join(parentDirPath, key)

		err := os.MkdirAll(subDirPath, 0755)
		if err != nil {
			log.Fatal(err)
		}

		err = subDir.Scaffold(subDirPath, data)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := createFiles(d.Files, parentDirPath, data)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func createFiles(files []File, parentDirPath string, data map[string]string) error {
	log.SetPrefix("func createFiles([]File, string, map[string]string) error:")

	for _, file := range files {
		err := createFile(file, parentDirPath, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func createFile(file File, parentDirPath string, data map[string]string) error {
	log.SetPrefix("func createFile(File, string, map[string]string) error:")

	filePath := filepath.Join(parentDirPath, file.Name)

	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}

	newFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer newFile.Close()

	if file.Content == "" {
		return nil
	}

	tmpl, err := template.New(file.Name).Parse(file.Content)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(newFile, data)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
