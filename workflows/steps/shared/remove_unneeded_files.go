package shared

import (
	"fmt"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type RemoveUnneededFiles struct{}

func (s *RemoveUnneededFiles) Execute(parameters steps.StepParameters) error {
	// scan the given path, and if a match is found inside deleteThisFile array then delete it

	deleteThisFiles := []string{".editorconfig", ".env.test", ".env.dev", ".env.local"}

	for _, file := range deleteThisFiles {
		if err := utilities.DeleteFile(parameters.ProjectName + "/" + file); err != nil {
			return fmt.Errorf("failed to delete file: %w", err)
		}
	}

	return nil
}
