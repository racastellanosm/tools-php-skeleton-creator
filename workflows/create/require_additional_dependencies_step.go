package create

import (
	"fmt"
	"os"

	"github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/utilities"
)

type RequireAdditionalDependenciesStep struct{}

func (s *RequireAdditionalDependenciesStep) Execute(projectName string) error {
	fmt.Println("* Require Additional Dependencies Step")

	if err := os.Chdir(projectName); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", projectName, err)
	}

	dependencies := []string{
		"symfony/messenger",
		"symfony/http-client",
		"symfony/uid",
	}

	for _, dep := range dependencies {
		if output, err := utilities.RunComposer(utilities.ComposerRunner{}, []string{"require", dep}); err != nil {
			return fmt.Errorf("failed to add dependency %s: %w\n%s", dep, err, output)
		}
	}

	return nil
}
