package shared

import (
	"fmt"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type AddMakefileStep struct{}

func (s *AddMakefileStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Add Makefile Step inside  " + parameters.ProjectName)

	destinationPath := fmt.Sprintf("%s/Makefile", parameters.ProjectName)
	sourceFile := "templates/makefile/Makefile.in"

	if err := utilities.CopyFile(sourceFile, destinationPath); err != nil {
		return fmt.Errorf("failed to copy Makefile: %w", err)
	}

	return nil
}
