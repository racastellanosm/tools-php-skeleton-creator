package shared

import (
	"fmt"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type AddDockerComposeFileStep struct{}

func (s *AddDockerComposeFileStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Add Docker Compose Files Step")

	files := []string{"docker-compose.yaml", "docker-compose.override.yaml", "Dockerfile"}

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
