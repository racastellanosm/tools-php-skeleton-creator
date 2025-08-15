package create

import (
	"fmt"
	"os"

	"github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/utilities"
)

type AddDockerComposeFileStep struct{}

func (s *AddDockerComposeFileStep) Execute(projectName string) error {
	fmt.Println("* Add Docker Compose Files Step" + projectName)

	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", projectName, err)
	}

	destinationPath := fmt.Sprintf("%s/docker-compose.yaml", projectName)
	sourceFile := "templates/docker/docker-compose.yaml"
	if err := utilities.CopyFile(sourceFile, destinationPath); err != nil {
		return fmt.Errorf("failed to copy docker-compose.yaml file: %w", err)
	}

	destinationPath = fmt.Sprintf("%s/Dockerfile", projectName)
	sourceFile = "templates/docker/Dockerfile"
	if err := utilities.CopyFile(sourceFile, destinationPath)
		err != nil {
		return fmt.Errorf("failed to copy Dockerfile file: %w", err)
	}

	return nil
}
