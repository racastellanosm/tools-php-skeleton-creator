package create

import (
	"fmt"
	"os"
)

type ReorganizeNeededFoldersStep struct{}

func (s *ReorganizeNeededFoldersStep) Execute(projectName string) error {
	fmt.Println("* Reorganize Needed Folders Step" + projectName)

	if err := os.RemoveAll("src/Controller"); err != nil {
		return fmt.Errorf("failed to remove directory src/Controllers: %w", err)
	}

	dirs := []string{
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
