package create

import (
	"fmt"
	"os"
)

type ReorganizeNeededFoldersStep struct{}

func (s *ReorganizeNeededFoldersStep) Execute(projectName string) error {
	fmt.Println("* Reorganize Needed Folders Step " + projectName)

	dirs := []string{
		"public",
		"src",
		"src/Application",
		"src/Domain",
		"src/Infrastructure",
		"src/UI/Http/Controller",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return nil
}
