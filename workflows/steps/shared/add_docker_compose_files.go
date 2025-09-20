package shared

import (
	"fmt"
	"os"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type AddDockerComposeFileStep struct{}

func (s *AddDockerComposeFileStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Add Docker Compose Files Step" + parameters.ProjectName)

	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", parameters.ProjectName, err)
	}

	files := []string{"docker-compose.yaml", "docker-compose.override.yaml", "Dockerfile", "env.dist"}

	switch parameters.Database {
	case "postgresql":
		sourceDir := "templates/docker/postgresql"
		destinationDir := fmt.Sprintf("%s", parameters.ProjectName)
		if err := utilities.CopyFileFromDirectory(sourceDir, destinationDir, files); err != nil {
			return fmt.Errorf("failed to copy docker with postgresql files: %w", err)
		}
	case "mysql":
		sourceDir := "templates/docker/mysql"
		destinationDir := fmt.Sprintf("%s", parameters.ProjectName)
		if err := utilities.CopyFileFromDirectory(sourceDir, destinationDir, files); err != nil {
			return fmt.Errorf("failed to copy docker with mysql files: %w", err)
		}
	}

	return nil
}
