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

	destinationPath := fmt.Sprintf("%s/docker-compose.yaml", parameters.ProjectName)
	sourceFile := "templates/docker/docker-compose.yaml"
	if err := utilities.CopyFile(sourceFile, destinationPath); err != nil {
		return fmt.Errorf("failed to copy docker-compose.yaml file: %w", err)
	}

	destinationPath = fmt.Sprintf("%s/Dockerfile", parameters.ProjectName)
	sourceFile = "templates/docker/Dockerfile"
	if err := utilities.CopyFile(sourceFile, destinationPath); err != nil {
		return fmt.Errorf("failed to copy Dockerfile file: %w", err)
	}

	return nil
}
