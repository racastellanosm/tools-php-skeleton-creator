package symfony

import (
	"fmt"
	"os"

	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type ReorganizeNeededFoldersStep struct{}

func (s *ReorganizeNeededFoldersStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Reorganize Needed Folders Step" + parameters.ProjectName)

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
