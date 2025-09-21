package symfony

import (
	"fmt"
	"os"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type MergeEnvironmentFiles struct{}

func (s *MergeEnvironmentFiles) Execute(parameters steps.StepParameters) error {
	fmt.Printf("* Merge Environment Files Step for %s\n", parameters.ProjectName)

	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", parameters.ProjectName, err)
	}

	templateFile := "templates/docker/" + parameters.Database + "/env.dist"
	targetPath := parameters.ProjectName + "/.env"

	if err := os.Remove(targetPath); err != nil {
		return fmt.Errorf("failed to remove environment file: %w", err)
	}

	if err := utilities.CopyFile(templateFile, targetPath); err != nil {
		return fmt.Errorf("failed to copy environment file: %w", err)
	}

	return nil
}
