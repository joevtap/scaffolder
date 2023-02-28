package scaffolder

import (
	"os"
	"path/filepath"
	"testing"
)

func setupAppRoot() string {
	appRoot, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	return appRoot
}

func cleanup(path string) {
	if err := os.RemoveAll(path); err != nil {
		panic(err)
	}
}

func TestScaffold(t *testing.T) {
	t.Run("should scaffold a project with single file", func(t *testing.T) {
		err := Scaffold("testdata/single-file", "testdata/file.toml", setupAppRoot(), nil)
		if err != nil {
			t.Errorf("error scaffolding project: %v", err)
		}

		cleanup("testdata/single-file")
	})

	t.Run("should scaffold a project with multiple files", func(t *testing.T) {
		err := Scaffold("testdata/multiple-files", "testdata/files.toml", setupAppRoot(), nil)
		if err != nil {
			t.Errorf("error scaffolding project: %v", err)
		}

		cleanup("testdata/multiple-files")
	})

	t.Run("should scaffold a project with multiple files and templates", func(t *testing.T) {
		err := Scaffold("testdata/multiple-files-and-templates", "testdata/files-and-templates.toml", setupAppRoot(), nil)
		if err != nil {
			t.Errorf("error scaffolding project: %v", err)
		}

		cleanup("testdata/multiple-files-and-templates")
	})

	t.Run("should scaffold a project with multiple files and templates with variables", func(t *testing.T) {
		err := Scaffold("testdata/multiple-files-and-templates-with-variables", "testdata/files-and-templates-with-variables.toml", setupAppRoot(), map[string]string{
			"testvar": "ok",
		})
		if err != nil {
			t.Errorf("error scaffolding project: %v", err)
		}

		cleanup("testdata/multiple-files-and-templates-with-variables")
	})

	t.Run("should scaffold a project with nested directories", func(t *testing.T) {
		err := Scaffold("testdata/nested-directories", "testdata/nested-directories.toml", setupAppRoot(), nil)
		if err != nil {
			t.Errorf("error scaffolding project: %v", err)
		}

		cleanup("testdata/nested-directories")
	})

	t.Run("should scaffold a project with nested directories and files", func(t *testing.T) {
		err := Scaffold("testdata/nested-directories-and-files", "testdata/nested-directories-and-files.toml", setupAppRoot(), nil)
		if err != nil {
			t.Errorf("error scaffolding project: %v", err)
		}

		cleanup("testdata/nested-directories-and-files")
	})
}
