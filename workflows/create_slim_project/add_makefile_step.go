package create

import (
	"fmt"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
)

type AddMakefileStep struct{}

func (s *AddMakefileStep) Execute(projectName string) error {
	fmt.Println("* Add Makefile Step inside  " + projectName)

	destinationPath := fmt.Sprintf("%s/Makefile", projectName)
	sourceFile := "templates/makefile/Makefile.in"

	if err := utilities.CopyFile(sourceFile, destinationPath); err != nil {
		return fmt.Errorf("failed to copy Makefile: %w", err)
	}

	return nil
}
