package shared

import (
	"fmt"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type AddMakefileStep struct{}

func (s *AddMakefileStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Add Makefile Step")

	destinationPath := fmt.Sprintf("%s/Makefile", parameters.ProjectName)
	sourceFile := "templates/makefile/Makefile"

	if err := utilities.CopyFile(sourceFile, destinationPath); err != nil {
		return fmt.Errorf("failed to copy Makefile: %w", err)
	}

	makefile := parameters.ProjectName + "/Makefile"
	oldString := "project_name_placeholder"
	newString := parameters.ProjectName

	if err := utilities.ReplaceAndCreate(makefile, oldString, newString); err != nil {
		return fmt.Errorf("failed to replace and create file: %w", err)
	}

	return nil
}
