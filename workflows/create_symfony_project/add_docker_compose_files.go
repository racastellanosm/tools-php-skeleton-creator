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

	if err := utilities.CopyFile("docker-compose.yaml", projectName+"/docker-compose.yaml"); err != nil {
		return fmt.Errorf("failed to copy docker-compose.yaml file: %w", err)
	}

	if err := utilities.CopyFile("Dockerfile", projectName+"/Dockerfile")
		err != nil {
		return fmt.Errorf("failed to copy Dockerfile file: %w", err)
	}

	return nil
}
