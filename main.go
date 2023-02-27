package main

import (
	"log"

	"github.com/joevtap/scaffolder/scaffolder"
)

func main() {
	err := scaffolder.Scaffold("teste", "project.toml", map[string]string{
		"name":        "Joel",
		"projectName": "Joel Test Project",
		"message":     "Hello World",
	})

	if err != nil {
		log.Fatal(err)
	}
}
