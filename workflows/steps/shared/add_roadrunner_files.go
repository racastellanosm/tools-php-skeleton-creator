package shared

import (
	"fmt"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type AddRoadrunnerFiles struct{}

func (s *AddRoadrunnerFiles) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Add RoadRunner Files Step")

	files := []string{".rr.yaml", ".rr.dev.yaml"}

	for _, file := range files {
		sourcePath := "templates/roadrunner/" + file
		destinationPath := parameters.ProjectName + "/" + file

		if err := utilities.CopyFile(sourcePath, destinationPath); err != nil {
			return fmt.Errorf("failed to copy file: %w", err)
		}
	}

	return nil
}
