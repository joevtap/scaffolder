package scaffolder

import (
	"os"

	"github.com/go-playground/validator/v10"
)

// validateProject validates the project config file based on the tags defined on type Project.
func validateProject(project Project) error {
	v := validator.New()
	return v.Struct(project)
}

// pathExists returns true if the provided path exists.
func pathExists(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}
