package shared

import (
	"fmt"
	"os"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type MergeEnvironmentFilesStep struct{}

func (s *MergeEnvironmentFilesStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Merge Environment Files Step")

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

	dockerignore := "templates/docker/" + parameters.Database + "/dockerignore.dist"
	if err := utilities.CopyFile(dockerignore, parameters.ProjectName+"/.dockerignore"); err != nil {
		return fmt.Errorf("failed to copy dockerignore file: %w", err)
	}

	envFile := parameters.ProjectName + "/.env"
	oldString := "project_name_placeholder"
	newString := parameters.ProjectName

	if err := utilities.ReplaceAndCreate(envFile, oldString, newString); err != nil {
		return fmt.Errorf("failed to replace and create file: %w", err)
	}

	return nil
}
