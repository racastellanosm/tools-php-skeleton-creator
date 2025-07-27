package create

import (
	"fmt"

	"github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/utilities"
)

type AddMakefileStep struct{}

func (s *AddMakefileStep) Execute(projectName string) error {
	fmt.Println("* Add Makefile Step inside  " + projectName)

	if err := utilities.CopyFile("templates/makefile/Makefile.in", projectName+"/Makefile"); err != nil {
		return fmt.Errorf("failed to copy Makefile: %w", err)
	}

	return nil
}
